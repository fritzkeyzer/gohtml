// Code generated by gohtml. DO NOT EDIT

package views

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"net/http"
	"os"
)

var LiveReload = os.Getenv("GOHTML_LIVERELOAD") != ""

//go:embed *.gohtml
var tmplFiles embed.FS
var templates = template.Must(template.ParseFS(tmplFiles, "*.gohtml"))

func tmpl() *template.Template {
	if LiveReload {
		return template.Must(template.ParseFS(os.DirFS("views"), "*.gohtml"))
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

// BEGIN: Page - - - - - - - -

type PageData struct {
	Title       any
	Description any
	SignedIn    any
	Username    any
	Body        any
}

// Page renders the "Page" template as an HTML fragment
func Page(data PageData) template.HTML {
	return mustHTML(RenderPage, data)
}

// RenderPage renders the "Page" template to a writer
func RenderPage(w io.Writer, data PageData) error {
	return tmpl().ExecuteTemplate(w, "page.gohtml", data)
}

// RenderPageHTTP renders page.gohtml to an http.ResponseWriter
func RenderPageHTTP(w http.ResponseWriter, data PageData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl().ExecuteTemplate(w, "page.gohtml", data)
}

// BEGIN: Person - - - - - - - -

type PersonContact struct {
	Phone any
	Email any
}

type PersonData struct {
	Name    any
	Age     any
	Contact PersonContact
	Socials []PersonSocial
}

type PersonSocial struct {
	Href any
	Name any
}

// Person renders the "Person" template as an HTML fragment
func Person(data PersonData) template.HTML {
	return mustHTML(RenderPerson, data)
}

// RenderPerson renders the "Person" template to a writer
func RenderPerson(w io.Writer, data PersonData) error {
	return tmpl().ExecuteTemplate(w, "person.gohtml", data)
}

// RenderPersonHTTP renders person.gohtml to an http.ResponseWriter
func RenderPersonHTTP(w http.ResponseWriter, data PersonData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl().ExecuteTemplate(w, "person.gohtml", data)
}
