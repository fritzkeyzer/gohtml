package gohtml

import (
	"strings"
	"text/template/parse"
)

// nodeToString recursively converts a template.Node to its string representation
func nodeToString(node parse.Node) string {
	switch n := node.(type) {
	case *parse.TextNode:
		return string(n.Text)
	case *parse.ActionNode:
		return "{{" + n.Pipe.String() + "}}"
	case *parse.CommandNode:
		return n.String()
	case *parse.IdentifierNode:
		return n.String()
	case *parse.PipeNode:
		return n.String()
	case *parse.IfNode:
		return "{{if " + n.Pipe.String() + "}}" + nodesToString(n.List) + "{{end}}"
	case *parse.RangeNode:
		return "{{range " + n.Pipe.String() + "}}" + nodesToString(n.List) + "{{end}}"
	case *parse.WithNode:
		return "{{with " + n.Pipe.String() + "}}" + nodesToString(n.List) + "{{end}}"
	case *parse.TemplateNode:
		return "{{template \"" + n.Name + "\" " + n.Pipe.String() + "}}"
	case *parse.ListNode:
		return nodesToString(n)
	case *parse.BranchNode:
		return nodeToString(n.Pipe)
	default:
		return ""
	}
}

// nodesToString converts a list of nodes to their string representation
func nodesToString(list *parse.ListNode) string {
	var sb strings.Builder
	for _, node := range list.Nodes {
		sb.WriteString(nodeToString(node))
	}
	return sb.String()
}
