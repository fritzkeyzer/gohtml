package gohtml

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"text/template"
)

//go:embed generate.go.template
var goTempl string

func (t *GoHTML) Generate(w io.Writer) error {
	buf := &bytes.Buffer{}
	err := template.Must(template.New(t.Name).Parse(goTempl)).Execute(buf, t)
	if err != nil {
		return fmt.Errorf("executing gen template: %w", err)
	}

	// apply std go formatting
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("go syntax error: %w", err)
	}

	_, err = w.Write(formatted)
	if err != nil {
		return fmt.Errorf("write gen file: %w", err)
	}

	return nil
}
