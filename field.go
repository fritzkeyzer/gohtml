package gohtml

import (
	"html/template"
	"strings"
	"text/template/parse"

	"github.com/fritzkeyzer/gohtml/logz"
)

func extractTemplateFields(templateName string, t *template.Template) []Field {
	if t.Tree == nil || t.Tree.Root == nil || t.Tree.Root.Nodes == nil {
		return nil
	}
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
	case *parse.TextNode: // Plain text.
	case *parse.ActionNode: // A non-control action such as a field evaluation.
		actionField := parseActionNodeField(path, n.Pipe.String())
		fields = []Field{actionField}
	case *parse.BoolNode: // A boolean constant.
	case *parse.ChainNode: // A sequence of field accesses.
	case *parse.CommandNode: // An element of a pipeline.
	case *parse.DotNode: // The cursor, dot.
	case *parse.FieldNode: // A field or method name.
	case *parse.IdentifierNode: // An identifier; always a function name.
	case *parse.IfNode: // An if action.
		ifNode := parseIfNodeField(path, n.Pipe.String())
		fields = []Field{ifNode}
		for _, n := range n.List.Nodes {
			childFields := parseNode(path, n)
			fields = append(fields, childFields...)
		}
	case *parse.ListNode: // A list of Nodes.
	case *parse.NilNode: // An untyped nil constant.
	case *parse.NumberNode: // A numerical constant.
	case *parse.PipeNode: // A pipeline of commands.
	case *parse.RangeNode: // A range action.
		rangeField := parseRangeNodeField(path, n.Pipe.String())
		fields = []Field{rangeField}

		rangeFieldType := strings.TrimPrefix(rangeField.Type, "[]")

		for _, n := range n.List.Nodes {
			childFields := parseNode([]string{rangeFieldType}, n)
			fields = append(fields, childFields...)
		}
	case *parse.StringNode: // A string constant.
	case *parse.TemplateNode: // A template invocation action.
		templateField := parseTemplateNodeField(path, n.String())
		if templateField.Name != "" {
			fields = []Field{templateField}
		}
	case *parse.VariableNode: // A $ variable.
	case *parse.WithNode: // A with action.
	case *parse.CommentNode: // A comment.
	case *parse.BreakNode: // A break action.
	case *parse.ContinueNode: // A continue action.
	}

	if len(fields) > 0 {
		logz.Debug("parsed node", "path", path, "fields", fields, "node", node.String())
	}
	
	return fields
}

func parseActionNodeField(path []string, pipe string) Field {
	//logz.Debug("parsing action pipe", "pipe", pipe, "path", path)

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

// parseIfNodeField from
func parseIfNodeField(path []string, pipe string) Field {
	//logz.Debug("parsing if pipe", "pipe", pipe, "path", path)

	// trim operator
	pipeParts := strings.Split(pipe, " ")
	pipe = pipeParts[len(pipeParts)-1]

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
	logz.Debug("parsing range pipe", "pipe", pipe, "path", path)

	parts := strings.Split(pipe, ".")
	fieldName := parts[len(parts)-1]
	fieldType := "[]" + strings.Join(path, "")

	if strings.Contains(pipe, ":=") {
		defSide, _, _ := strings.Cut(pipe, " := ")

		keyStr, valStr, _ := strings.Cut(defSide, ", ")
		key := strings.TrimLeft(keyStr, "$. ")

		if valStr != "" {
			fieldType += fieldName + strings.Title(key)
		} else if strings.HasSuffix(fieldName, "s") {
			fieldType += strings.TrimSuffix(fieldName, "s")
		} else {
			fieldType += fieldName + "Item"
		}
	} else {
		if strings.HasSuffix(fieldName, "s") {
			fieldType += strings.TrimSuffix(fieldName, "s")
		} else {
			fieldType += fieldName + "Item"
		}
	}

	return Field{
		Path: path,
		Name: fieldName,
		Type: fieldType,
	}
}

func parseTemplateNodeField(path []string, node string) Field {
	//logz.Debug("parsing template node", "node", node, "path", path)
	node = strings.TrimFunc(node, func(r rune) bool {
		return r == '{' || r == '}'
	})
	parts := strings.Split(node, " ")

	referencedTemplateName := strings.Trim(parts[1], "\"")
	pipe := strings.Join(parts[2:], " ")

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
	parts = strings.Split(pipe, ".")

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
		Type: strings.Title(referencedTemplateName) + "Data",
	}
}
