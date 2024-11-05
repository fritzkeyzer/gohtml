package gohtml

import (
	"html/template"
	"slices"
	"strings"

	"github.com/fritzkeyzer/gohtml/logz"
)

func ParseTemplateFile(tmpl *template.Template) (TemplateFile, error) {
	name := strings.TrimSuffix(tmpl.Name(), ".gohtml")
	name = strings.Title(name)

	parsed := TemplateFile{
		Name: name,
	}

	logz.Debug("Parse template", "subTemplate.Name()", tmpl.Name())
	subName := tmpl.Name()
	if subName == tmpl.Name() {
		subName = name
	}
	subName = strings.Title(subName)

	// get a list of all variables used in the template, by traversing the AST of the template
	fields := extractTemplateFields(subName, tmpl)

	// build structs from fields
	var structs []StructDef
	for _, f := range fields {
		structs = addField(subName, structs, f)
	}

	// rename base struct to ...Data
	for i := range structs {
		if structs[i].Name == subName {
			structs[i].Name = subName + "Data"
		}
	}

	// register a function to generate
	fn := FnDef{
		Name:         subName,
		Args:         nil,
		TemplateName: tmpl.Name(),
	}
	if len(structs) > 0 {
		fn.Args = append(fn.Args, structs[0].Name)
	}
	parsed.Fns = append(parsed.Fns, fn)
	parsed.Structs = append(parsed.Structs, structs...)

	slices.SortFunc(parsed.Structs, func(a, b StructDef) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})
	slices.SortFunc(parsed.Fns, func(a, b FnDef) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})

	return parsed, nil
}

func addField(templateName string, structs []StructDef, field Field) []StructDef {
	structName := strings.Join(field.Path, "")

	// find and append struct (if it exists)
	for i := range structs {
		if structs[i].Name == structName {
			fieldExists := false
			for _, f := range structs[i].Fields {
				if f.Name == field.Name {
					fieldExists = true
					break
				}
			}

			if !fieldExists {
				structs[i].Fields = append(structs[i].Fields, field)
			}

			return structs
		}
	}

	// create new struct def
	structs = append(structs, StructDef{
		Name:   structName,
		Fields: []Field{field},
	})

	// create missing links in data model
	if len(field.Path) > 1 && !strings.HasPrefix(field.Type, "[]") {
		parentField := Field{
			Path: field.Path[:len(field.Path)-1],
			Name: field.Path[len(field.Path)-1],
			Type: structName,
		}

		// find and add field here
		structs = addField(templateName, structs, parentField)
	}

	return structs
}
