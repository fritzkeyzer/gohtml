package main

import (
	"os"

	"github.com/fritzkeyzer/gohtml/example/views"
)

func main() {
	// imagine that this function is a http request handler
	w := os.Stdout

	views.LiveReload = true // defaults to os.Getenv("GOHTML_LIVERELOAD") != ""

	// render a HTML partial
	person := views.Person(views.PersonData{
		Name: "John Doe",
		Age:  42,
		Contact: views.PersonContact{
			Phone: "0123456789",
			Email: "john@doe.com",
		},
		Socials: []views.PersonSocial{
			{Href: "https://twitter.com/johndoe", Name: "Twitter"},
			{Href: "https://facebook.com/johndoe", Name: "Facebook"},
		},
	})

	// render another template out to writer
	err := views.RenderPage(w, views.PageData{
		Title:       "Example page",
		Description: "Demonstrate basic usage of gohtml",
		SignedIn:    false,
		Username:    nil,
		Body:        person, // note that we can nest partials within each other
	})
	if err != nil {
		panic(err)
	}

	/*Output:
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>Example page</title>
	    <meta name="description" content="Demonstrate basic usage of gohtml">
	</head>
	<body>
	<div>

	</div>
	<h1>Person</h1>
	<p>Name: John Doe</p>
	<p>Age: 42</p>
	<p>Phone: 0123456789</p>
	<p>Email: john@doe.com</p>
	<div>
	        <a href="https://twitter.com/johndoe">Twitter</a>
	        <a href="https://facebook.com/johndoe">Facebook</a>
	</div>

	</body>
	</html>
	*/
}
