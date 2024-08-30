package gohtml

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed generate_pkg.go.template
var goPkgTempl string

func GohtmlFile(dir, name string) {
	f, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	err = template.Must(template.New("package").Parse(goPkgTempl)).Execute(buf, map[string]any{
		"PackageName": filepath.Base(dir),
	})
	if err != nil {
		panic(fmt.Errorf("executing template: %v", err))
	}

	// apply std go formatting
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Errorf("go syntax error: %w", err))
	}

	_, err = f.Write(formatted)
	if err != nil {
		panic(fmt.Errorf("write gen file: %w", err))
	}
}
