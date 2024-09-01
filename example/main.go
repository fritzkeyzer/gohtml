package main

import (
	"bytes"
	"fmt"

	"github.com/fritzkeyzer/gohtml/example/views"
)

func main() {
	// imagine that this function is a http request handler

	// you could make some database query here

	// render the person.gohtml template, with a typesafe function
	bodyHTML := views.Person(views.PersonData{
		Name: "Bob",
		Age:  123,
		Contact: views.PersonContact{
			Phone: "012 234",
			Email: "bob@example.com",
		},
		Socials: []views.PersonSocialsLink{
			{
				Name: "facebook",
				Href: "facebook.com/bob",
			}, {
				Name: "linkedin",
				Href: "linkedin.com/bob",
			},
		},
	})

	// render another template to Stdout (this would typically be an http.ResponseWriter)
	// notice how we inject an html fragment from one component into another
	buf := new(bytes.Buffer)
	err := views.RenderPage(buf, views.PageData{
		Title:       "Hello world",
		Description: "Example page",
		SignedIn:    true,
		Username:    "Bob",
		Body:        bodyHTML,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

/* Output:
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello world</title>
    <meta name="description" content="Example page">
</head>
<body>
<div>
    <p>Hello Bob</p>
</div>
<h1>Person</h1>
<p>Name: Bob</p>
<p>Age: 123</p>
<p>Phone: 012 234</p>
<p>Email: bob@example.com</p>
<div>
        <a href="facebook.com/bob">facebook</a>
        <a href="linkedin.com/bob">linkedin</a>
</div>

</body>
</html>

*/
