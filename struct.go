package gohtml

import "strings"

type StructDef struct {
	Name   string
	Fields []Field
}

func addField(templateName string, structs []StructDef, field Field) []StructDef {
	structName := strings.Join(field.Path, "")

	// find and append struct (if it exists)
	for i := range structs {
		if structs[i].Name == structName {
			fieldExists := false
			for j, f := range structs[i].Fields {
				if f.Name == field.Name {
					fieldExists = true

					// update to use the more specific type
					if f.Type == "any" && field.Type != "any" {
						structs[i].Fields[j].Type = field.Type
					}

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
			Type: "*" + structName,
		}

		// find and add field here
		structs = addField(templateName, structs, parentField)
	}

	return structs
}
