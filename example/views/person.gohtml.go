// Code generated by gohtml. DO NOT EDIT

//go:build !ignore_autogenerated

package views

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

// <<< START TEMPLATE: Person

var rawPersonTemplate =
// language=gotemplate
`<h1>Person</h1>
<p>Name: {{.Name}}</p>
<p>Age: {{.Age}}</p>
<p>Phone: {{.Contact.Phone}}</p>
<p>Email: {{.Contact.Email}}</p>
<div>{{range $link := .Socials}}
        <a href="{{$link.Href}}">{{$link.Name}}</a>{{end}}
</div>
`

var PersonTemplate = template.Must(template.New("Person").Parse(rawPersonTemplate))

type PersonData struct {
	Name    any
	Age     any
	Contact PersonContact
	Socials []PersonSocialsLink
}

type PersonContact struct {
	Phone any
	Email any
}

type PersonSocialsLink struct {
	Href any
	Name any
}

// Person renders the person.gohtml template as an HTML fragment
func Person(data PersonData) template.HTML {
	buf := new(bytes.Buffer)
	err := RenderPerson(buf, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}

// RenderPerson renders the person.gohtml template to the specified writer.
// For writing to an http.ResponseWriter - use RenderPersonHTTP instead.
func RenderPerson(w io.Writer, data PersonData) error {
	tmpl := PersonTemplate
	if LiveReload {
		var err error
		tmpl, err = template.ParseFiles("views/person.gohtml")
		if err != nil {
			return err
		}
	}

	return tmpl.Execute(w, data)
}

// RenderPersonHTTP renders the person.gohtml template to the http.ResponseWriter.
// Errors are handled with the package global views.ErrorFn function (which can be customized) and returned.
// You can choose to handle errors with the views.ErrorFn handler, the returned error, or both.
func RenderPersonHTTP(w http.ResponseWriter, data PersonData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err := RenderPerson(buf, data)
	if err != nil {
		ErrorFn(w, err)
		return err
	}

	_, _ = w.Write(buf.Bytes())

	return nil
}

// >>> END TEMPLATE: Person
