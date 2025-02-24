// Code generated by gohtml. DO NOT EDIT

package views

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"os"
)

import "net/http"

var LiveReload = os.Getenv("GOHTML_LIVERELOAD") != ""

//go:embed *.gohtml
var tmplFiles embed.FS
var templates = template.Must(template.ParseFS(tmplFiles, "*.gohtml"))

func tmpl() *template.Template {
	if LiveReload {
		return template.Must(template.ParseFS(os.DirFS("nested/views"), "*.gohtml"))
	}
	return templates
}

func mustHTML[T any](fn func(w io.Writer, data T) error, data T) template.HTML {
	w := new(bytes.Buffer)
	err := fn(w, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(w.String())
}

func mustHTMLNoArgs(fn func(w io.Writer) error) template.HTML {
	w := new(bytes.Buffer)
	err := fn(w)
	if err != nil {
		panic(err)
	}
	return template.HTML(w.String())
}

// BEGIN: Nested - - - - - - - -

type NestedData struct {
	Name any
}

// Nested renders the "Nested" template as an HTML fragment
func Nested(data NestedData) template.HTML {
	return mustHTML(RenderNested, data)
}

// RenderNested renders the "Nested" template to a writer
func RenderNested(w io.Writer, data NestedData) error {
	return tmpl().ExecuteTemplate(w, "nested.gohtml", data)
}

// RenderNestedHTTP renders "nested.gohtml" to an http.ResponseWriter
func RenderNestedHTTP(w http.ResponseWriter, data NestedData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl().ExecuteTemplate(w, "nested.gohtml", data)
}
