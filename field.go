package gohtml

import (
	"fmt"
	"html/template"
	"strings"
	"text/template/parse"
)

type Field struct {
	Path []string
	Name string
	Type string
}

func extractTemplateFields(templateName string, t *template.Template) []Field {
	var fields []Field
	for _, node := range t.Tree.Root.Nodes {
		fields = append(fields, parseNode([]string{templateName}, node)...)
	}
	
	return fields
}

func parseNode(path []string, node parse.Node) []Field {
	var fields []Field
	switch n := node.(type) {
	case *parse.ActionNode:
		fields = []Field{parseActionNodeField(path, n.Pipe.String())}
	case *parse.IfNode:
		fields = []Field{parseActionNodeField(path, n.Pipe.String())}
	case *parse.RangeNode:
		rangeField := parseRangeNodeField(path, n.Pipe.String())
		fields = append(fields, rangeField)

		rangeFieldType := strings.TrimPrefix(rangeField.Type, "[]")

		for _, n := range n.List.Nodes {
			childFields := parseNode([]string{rangeFieldType}, n)
			fields = append(fields, childFields...)
		}
	}

	return fields
}

func parseActionNodeField(path []string, pipe string) Field {
	if DebugFlag {
		fmt.Printf("parsing action node {{%s}} with path: %v\n", pipe, path)
	}

	// in the case of named range variables, eg: {{$link.Name}}
	if strings.HasPrefix(pipe, "$") {
		pipe = strings.TrimLeft(pipe, "$ ")
		split := strings.Split(pipe, ".")
		if len(split) > 1 {
			pipe = strings.Join(split[1:], ".")
		}
	}

	pipe = strings.TrimLeft(pipe, ".")
	parts := strings.Split(pipe, ".")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	path = append(path, parts[:len(parts)-1]...)

	// the last part - is a field of type any
	return Field{
		Path: path,
		Name: parts[len(parts)-1],
		Type: "any",
	}
}

func parseRangeNodeField(path []string, pipe string) Field {
	if DebugFlag {
		fmt.Printf("parsing range node {{%s}} with path: %v\n", pipe, path)
	}

	parts := strings.Split(pipe, ".")
	fieldName := parts[len(parts)-1]
	fieldType := "[]" + strings.Join(path, "")

	if strings.Contains(pipe, ":=") {
		defSide, _, _ := strings.Cut(pipe, " := ")
		defSide = strings.TrimLeft(defSide, "$. ")
		fieldType += fieldName + strings.Title(defSide)
	} else if strings.HasSuffix(fieldName, "s") {
		fieldType += strings.TrimSuffix(fieldName, "s")
	} else {
		fieldType += fieldName + "Item"
	}

	return Field{
		Path: path,
		Name: fieldName,
		Type: fieldType,
	}
}
