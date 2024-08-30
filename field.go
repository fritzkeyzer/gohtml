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

	for i := range fields {
		if len(fields[i].Path) == 0 {
			fields[i].Path = []string{templateName}
		}
	}

	return fields
}

func parseNode(path []string, node parse.Node) []Field {
	var fields []Field
	switch n := node.(type) {
	case *parse.ActionNode:
		fields = []Field{parseActionNodeField(path, n.Pipe.String())}
	case *parse.IfNode:
		ifNode := parseActionNodeField(path, n.Pipe.String())
		fields = []Field{ifNode}

		for _, n := range n.List.Nodes {
			childFields := parseNode(path, n)
			fields = append(fields, childFields...)
		}

	case *parse.RangeNode:
		rangeField := parseRangeNodeField(path, n.Pipe.String())
		fields = []Field{rangeField}

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

	// root context variables
	isRoot := false
	if strings.HasPrefix(pipe, "$.") {
		isRoot = true
	}

	// loop variables iterator references
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

	if isRoot {
		path = nil
	}

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
