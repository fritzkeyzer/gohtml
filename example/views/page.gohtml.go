// Code generated by gohtml. DO NOT EDIT

//go:build !ignore_autogenerated

package views

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

// <<< START TEMPLATE: Page

var rawPageTemplate =
// language=gotemplate
`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <meta name="description" content="{{.Description}}">
</head>
<body>
<div>
    {{if .SignedIn}}<p>Hello {{.Username}}</p>{{end}}
</div>
{{.Body}}
</body>
</html>`

var PageTemplate = template.Must(template.New("Page").Parse(rawPageTemplate))

type PageData struct {
	Title       any
	Description any
	SignedIn    any
	Username    any
	Body        any
}

// Page renders the page.gohtml template as an HTML fragment
func Page(data PageData) template.HTML {
	buf := new(bytes.Buffer)
	err := RenderPage(buf, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}

// RenderPage renders the page.gohtml template to the specified writer.
// For writing to an http.ResponseWriter - use RenderPageHTTP instead.
func RenderPage(w io.Writer, data PageData) error {
	tmpl := PageTemplate
	if LiveReload {
		var err error
		tmpl, err = template.ParseFiles("views/page.gohtml")
		if err != nil {
			return err
		}
	}

	return tmpl.Execute(w, data)
}

// RenderPageHTTP renders the page.gohtml template to the http.ResponseWriter.
// Errors are handled with the package global views.ErrorFn function (which can be customized) and returned.
// You can choose to handle errors with the views.ErrorFn handler, the returned error, or both.
func RenderPageHTTP(w http.ResponseWriter, data PageData) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err := RenderPage(buf, data)
	if err != nil {
		ErrorFn(w, err)
		return err
	}

	_, _ = w.Write(buf.Bytes())
	w.WriteHeader(http.StatusOK)

	return nil
}

// >>> END TEMPLATE: Page
