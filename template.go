package gohtml

import (
	"html/template"
	"slices"
	"strings"

	"github.com/fritzkeyzer/gohtml/logz"
	"github.com/iancoleman/strcase"
)

func ConvertTemplate(tmpl *template.Template) (Template, error) {
	name := strings.TrimSuffix(tmpl.Name(), ".gohtml")
	name = strcase.ToCamel(name)
	name = strings.Title(name)

	parsed := Template{
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
