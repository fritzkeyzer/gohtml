// Code generated by gohtml. DO NOT EDIT

//go:build !ignore_autogenerated

package tests

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

// <<< START TEMPLATE: Loops

var rawLoopsTemplate =
// language=gotemplate
`{{range .Widgets}}
    {{$.Currency}} {{.Price}} - {{.Name}}
{{end}}

{{range $link := .Socials}}
    {{$link.Name}} {{$link.Href}}
{{end}}
`

var LoopsTemplate = template.Must(template.New("Loops").Parse(rawLoopsTemplate))

type LoopsData struct {
	Widgets  []LoopsWidget
	Currency any
	Socials  []LoopsSocialsLink
}

type LoopsSocialsLink struct {
	Name any
	Href any
}

type LoopsWidget struct {
	Price any
	Name  any
}

// Loops renders the loops.gohtml template as an HTML fragment
func Loops(data LoopsData) template.HTML {
	buf := new(bytes.Buffer)
	err := RenderLoops(buf, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}

// RenderLoops renders the loops.gohtml template to the specified writer.
// For writing to an http.ResponseWriter - use RenderLoopsHTTP instead.
func RenderLoops(w io.Writer, data LoopsData) error {
	tmpl := LoopsTemplate
	if LiveReload {
		var err error
		tmpl, err = template.ParseFiles("tests/loops.gohtml")
		if err != nil {
			return err
		}
	}

	return tmpl.Execute(w, data)
}

// RenderLoopsHTTP renders the loops.gohtml template to the http.ResponseWriter.
// Errors are handled with the package global tests.ErrorFn function (which can be customized) and returned.
// You can choose to handle errors with the tests.ErrorFn handler, the returned error, or both.
func RenderLoopsHTTP(w http.ResponseWriter, data LoopsData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err := RenderLoops(buf, data)
	if err != nil {
		ErrorFn(w, err)
		return err
	}

	_, _ = w.Write(buf.Bytes())

	return nil
}

// >>> END TEMPLATE: Loops
