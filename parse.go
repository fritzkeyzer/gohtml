package gohtml

import (
	"fmt"
	"html/template"
	"path/filepath"
	"sort"
	"strings"

	"github.com/k0kubun/pp/v3"
)

var DebugFlag bool

type GoHTML struct {
	Name        string
	FilePath    string
	PackageName string
	Templates   []Template
}

type Template struct {
	TemplateString   string
	TemplateFilePath string
	Name             string
	EmbedFilePath    string
	Structs          []StructDef
}

type StructDef struct {
	Name   string
	Fields []Field
}

func ParseTemplate(templatePath, packageName string) (*GoHTML, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("parse template file: %v", err)
	}

	templateName := strings.Title(strings.TrimSuffix(filepath.Base(templatePath), ".gohtml"))

	data := &GoHTML{
		Name:        templateName,
		FilePath:    templatePath,
		PackageName: packageName,
	}

	for _, subTemplate := range tmpl.Templates() {
		subName := subTemplate.Name()
		if subName == tmpl.Name() {
			subName = templateName
		}

		// TODO check if the subTemplate has anything in it.
		// the first sub-template could be empty in cases where a template file exclusively contains sub-templates

		// get a list of all variables used in the template, by traversing the AST of the template
		fields := extractTemplateFields(subName, subTemplate)
		if DebugFlag {
			fmt.Println("Parse fields from template:", templateName, pp.Sprint(fields))
		}

		var structs []StructDef
		for _, f := range fields {
			structs = addField(templateName, structs, f)
		}

		// sort structs
		sort.Slice(structs, func(i, j int) bool {
			if structs[i].Name == subName+"Data" {
				return true
			}
			if structs[j].Name == subName+"Data" {
				return false
			}
			return structs[i].Name < structs[j].Name
		})

		// rename base struct to ...Data
		if len(structs) > 0 {
			structs[0].Name += "Data"
		}

		data.Templates = append(data.Templates, Template{
			TemplateString:   nodeToString(subTemplate.Tree.Root),
			TemplateFilePath: templatePath,
			Name:             strings.Title(strings.TrimSuffix(subTemplate.Name(), ".gohtml")),
			EmbedFilePath:    tmpl.Name(),
			Structs:          structs,
		})
	}

	return data, nil
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
