// Code generated by gohtml. DO NOT EDIT

//go:build !ignore_autogenerated

package tests

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
)

//go:embed person.gohtml
var PersonTemplate string

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
	Name any
	Href any
}

// Person renders the "Person" template as an HTML fragment
func Person(data PersonData) template.HTML {
	buf := new(bytes.Buffer)
	err := RenderPerson(buf, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}

// RenderPerson renders the "Person" template to the specified writer
func RenderPerson(w io.Writer, data PersonData) error {
	var tmpl *template.Template
	var err error
	if LiveReload {
		tmpl, err = template.ParseFiles("tests/person.gohtml")
	} else {
		tmpl, err = template.New("Person").Parse(PersonTemplate)
	}
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}
