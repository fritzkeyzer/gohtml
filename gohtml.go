package gohtml

import (
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"text/template/parse"

	"github.com/fritzkeyzer/gohtml/logz"
)

type GoHTML struct {
	Name        string
	FilePath    string
	PackageName string
	TemplateDir string
	Templates   []Template
}

type Template struct {
	Name    string
	Structs []StructDef
	Fns     []FnDef
}

type Field struct {
	Path []string
	Name string
	Type string
}

type FnDef struct {
	Name         string
	Args         []string
	TemplateName string
}

func ParseDir(dir string) (*GoHTML, error) {
	templateFiles, err := filepath.Glob(path.Join(dir, "*.gohtml"))
	if err != nil {
		return nil, fmt.Errorf("glob files in directory: %w", err)
	}
	if len(templateFiles) == 0 {
		return nil, fmt.Errorf("no template files found in %s", dir)
	}

	logz.Debug("found templates", "files", templateFiles)

	// check if the templates are valid
	t, err := template.New(dir).ParseFiles(templateFiles...)
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}
	logz.Debug("templates valid")

	// generate types and handler functions for each file
	g := &GoHTML{
		PackageName: path.Base(dir),
		Templates:   nil,
		TemplateDir: dir,
	}

	orderedTemplates := t.Templates()
	slices.SortStableFunc(orderedTemplates, func(a, b *template.Template) int {
		if a.Name() < b.Name() {
			return -1
		}
		if a.Name() > b.Name() {
			return 1
		}
		return 0
	})

	for _, tmpl := range orderedTemplates {
		// skip the package template
		if path.Base(tmpl.Name()) == path.Base(dir) {
			continue
		}

		// TODO skip empty templates, eg: file only containing subtemplates with empty root template
		hasContent, err := hasContent(tmpl)
		if err != nil {
			return nil, fmt.Errorf("check template content: %w", err)
		}
		if !hasContent {
			logz.Debug("skip empty template", "name", tmpl.Name())
			continue
		}

		logz.Debug("parse: " + tmpl.Name())

		parsed, err := ConvertTemplate(tmpl)
		if err != nil {
			return nil, fmt.Errorf("parse template: %w", err)
		}

		g.Templates = append(g.Templates, parsed)
	}

	if err = g.validate(); err != nil {
		return nil, err
	}

	return g, nil
}

func hasContent(t *template.Template) (bool, error) {
	recursiveContent := ""
	for _, n := range t.Tree.Root.Nodes {
		recursiveContent += renderTextNode(n)
	}

	// Remove all whitespace characters
	trimmed := strings.TrimSpace(recursiveContent)
	return len(trimmed) > 0, nil
}

func renderTextNode(n parse.Node) string {
	content := ""
	switch n := n.(type) {
	case *parse.TextNode:
		content += strings.TrimSpace(n.String())
	case *parse.RangeNode:
		if n.List != nil && n.List.Nodes != nil {
			for _, n := range n.List.Nodes {
				content += renderTextNode(n)
			}
		}
		if n.ElseList != nil && n.ElseList.Nodes != nil {
			for _, n := range n.ElseList.Nodes {
				content += renderTextNode(n)
			}
		}
	}
	return content
}

func (g *GoHTML) validate() error {
	var errors []error
	namespace := make(map[string]bool)

	for _, t := range g.Templates {
		for _, s := range t.Structs {
			if namespace[s.Name] {
				errors = append(errors, fmt.Errorf("duplicate name: %s", s.Name))
			}
			namespace[s.Name] = true
		}
		for _, f := range t.Fns {
			if namespace[f.Name] {
				errors = append(errors, fmt.Errorf("duplicate name: %s", f.Name))
			}
			namespace[f.Name] = true
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %+v", errors)
	}

	return nil
}
