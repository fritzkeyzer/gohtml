package parse

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"text/template/parse"
)

type Template struct {
	Name        string
	FilePath    string
	PackageName string
	Functions   []FuncDef
}

type FuncDef struct {
	Name          string
	EmbedFilePath string
	Types         []StructDef
}

type StructDef struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name string
	Type string
	Path []string
}

func MustParseTemplate(templatePath string) *Template {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	t, err := ParseTemplate(templatePath, tmpl)
	if err != nil {
		panic(err)
	}
	return t
}

func ParseTemplate(templatePath string, tmpl *template.Template) (*Template, error) {
	name := strings.TrimSuffix(filepath.Base(templatePath), ".gohtml")
	templateName := strings.ToUpper(string(name[0])) + name[1:]

	functions, err := parseFunctions(templateName, tmpl)
	if err != nil {
		return nil, fmt.Errorf("parse functions: %w", err)
	}

	packageName := filepath.Base(filepath.Dir(templatePath))
	data := &Template{
		Name:        templateName,
		FilePath:    templatePath,
		PackageName: packageName,
		Functions:   functions,
	}

	return data, nil
}

func parseFunctions(templateName string, t *template.Template) ([]FuncDef, error) {
	fns := []FuncDef{}

	types, err := parseTypes(templateName, t)
	if err != nil {
		return nil, fmt.Errorf("parse types: %w", err)
	}

	fns = append(fns, FuncDef{
		Name:          strings.Title(strings.TrimSuffix(t.Name(), ".gohtml")),
		EmbedFilePath: t.Name(),
		Types:         types,
	})

	return fns, nil
}

func parseTypes(templateName string, t *template.Template) ([]StructDef, error) {
	var structs []StructDef

	// get a list of all variables used in the template, by going though the AST of t

	var fields []StructField

	nodes := t.Tree.Root.Nodes
	for _, node := range nodes {
		fields = parseNode([]string{templateName}, node, fields)
	}

	for _, f := range fields {
		structs = addField(structs, f)
	}

	for i := range structs {
		if structs[i].Name == templateName {
			structs[i].Name += "Data"
		}
	}

	return structs, nil
}

func parseNode(path []string, node parse.Node, fields []StructField) []StructField {
	switch n := node.(type) {
	case *parse.ActionNode:
		fields = append(fields, parseActionNodeField(path, n.Pipe.String()))
	case *parse.IfNode:
		fields = append(fields, parseActionNodeField(path, n.Pipe.String()))
	case *parse.RangeNode:
		//fmt.Println(n.Pipe.String())

		rangeField := parseRangeNodeField(path, n.Pipe.String())
		fields = append(fields, rangeField)

		for _, n := range n.List.Nodes {
			fields = parseNode(append([]string{path[0], rangeField.Name}, path[1:]...), n, fields)
		}
	}

	return fields
}

func addField(structs []StructDef, field StructField) []StructDef {
	//fmt.Printf("add field %+v\n", field)

	structName := strings.Join(field.Path, "")

	// find and append struct
	for i := range structs {
		if structs[i].Name == structName {
			// append and return

			found := false
			for _, f := range structs[i].Fields {
				if f.Name == field.Name {
					found = true
					break
				}
			}

			if !found {
				structs[i].Fields = append(structs[i].Fields, field)
			}

			return structs
		}
	}

	// create new struct def
	structs = append(structs, StructDef{
		Name:   structName,
		Fields: []StructField{field},
	})

	// add field to reference the struct
	if len(field.Path) > 1 && !strings.HasPrefix(field.Type, "[]") {
		// 2nd last item in path
		//refStructName := field.Path[len(field.Path)-2]

		// find and add field here
		structs = addField(structs, StructField{
			Path: field.Path[:len(field.Path)-1],
			Name: field.Path[len(field.Path)-1],
			Type: structName,
		})
	}

	return structs
}

func parseActionNodeField(path []string, pipe string) StructField {
	//fmt.Println("parsing action node:", pipe)

	pipe = strings.TrimLeft(pipe, ".$")
	parts := strings.Split(pipe, ".")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	path = append(path, parts[:len(parts)-1]...)

	// the last part - is a field of type any
	return StructField{
		Path: path,
		Name: parts[len(parts)-1],
		Type: "any",
	}
}

func parseRangeNodeField(path []string, pipe string) StructField {
	// example this syntax:
	// Item := .Options

	// path: templateName
	// name: Options
	// type: []OptionsItem

	//fmt.Println("parsing range node:", pipe)

	split := strings.Split(pipe, " := ")
	if len(split) != 2 {
		panic(fmt.Sprintf("invalid range node: %s", pipe))
	}

	defSide := strings.TrimPrefix(split[0], "$")
	varSide := strings.TrimPrefix(split[1], ".")

	itemName := strings.Title(defSide)
	parts := strings.Split(varSide, ".")

	//path = append(path, parts[1:]...)

	// the last part - is a field of type any
	return StructField{
		Path: path,
		Name: parts[len(parts)-1],
		Type: "[]" + strings.Join(path, "") + parts[len(parts)-1] + itemName,
	}
}
