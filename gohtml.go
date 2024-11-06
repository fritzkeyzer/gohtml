package gohtml

import (
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"slices"

	"github.com/fritzkeyzer/gohtml/logz"
)

type GoHTML struct {
	Name        string
	FilePath    string
	PackageName string
	TemplateDir string
	Templates   []TemplateFile
}

type TemplateFile struct {
	Name    string
	Structs []StructDef
	Fns     []FnDef
}

type StructDef struct {
	Name   string
	Fields []Field
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

		logz.Info("parse: " + tmpl.Name())

		parsed, err := ParseTemplateFile(tmpl)
		if err != nil {
			return nil, fmt.Errorf("parse template: %w", err)
		}

		g.Templates = append(g.Templates, parsed)

		//for _, subTemplate := range tmpl.Templates() {
		//	logz.Debug("Parse template", "subTemplate.Name()", subTemplate.Name())
		//	subName := subTemplate.Name()
		//	if subName == tmpl.Name() {
		//		subName = baseName
		//	}
		//	subName = strings.Title(subName)
		//
		//	// first check if the subTemplate has anything in it.
		//	// the first sub-template could be empty in cases where a template file exclusively contains sub-templates
		//	if len(subTemplate.Tree.Root.Nodes) == 0 {
		//		logz.Debug("Parse empty template", "templatePath", templatePath, "subTemplate.Name()", subTemplate.Name())
		//		continue
		//	}
		//
		//	// get a list of all variables used in the template, by traversing the AST of the template
		//	fields := extractTemplateFields(subName, subTemplate)
		//
		//	// build structs from fields
		//	var structs []StructDef
		//	for _, f := range fields {
		//		structs = addField(subName, structs, f)
		//	}
		//
		//	// rename base struct to ...Data
		//	if len(structs) > 0 {
		//		structs[0].Name += "Data"
		//	}
		//
		//	// register a function to generate
		//	fn := FnDef{
		//		Name:         subName,
		//		Args:         nil,
		//		TemplateName: subTemplate.Name(),
		//	}
		//	if len(structs) > 0 {
		//		fn.Args = append(fn.Args, structs[0].Name)
		//	}
		//	parsed.Fns = append(parsed.Fns, fn)
		//	parsed.Structs = append(parsed.Structs, structs...)
		//
		//	slices.SortFunc(parsed.Structs, func(a, b StructDef) int {
		//		if a.Name < b.Name {
		//			return -1
		//		}
		//		if a.Name > b.Name {
		//			return 1
		//		}
		//		return 0
		//	})
		//	slices.SortFunc(parsed.Fns, func(a, b FnDef) int {
		//		if a.Name < b.Name {
		//			return -1
		//		}
		//		if a.Name > b.Name {
		//			return 1
		//		}
		//		return 0
		//	})
		//}

		//g.Templates = append(g.Templates, t)
	}

	if err = g.validate(); err != nil {
		return nil, err
	}

	return g, nil
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
			//if s.Name == "" {
			//	errors = append(errors, fmt.Errorf("struct name is empty"))
			//	continue
			//}
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
