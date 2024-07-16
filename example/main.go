package main

import (
	"os"

	"github.com/fritzkeyzer/gohtml/example/layout"
	"github.com/fritzkeyzer/gohtml/example/template"
)

func main() {
	body := template.Hello(template.HelloData{
		Name: "Bob",
		Age:  "123",
	})

	_ = layout.RenderBase(os.Stdout, layout.BaseData{
		Title: "Hello world",
		Body:  body,
	})
}

/* Output:
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello world</title>
</head>
<body>
<p>Hello Bob</p>
<p>Age: 123</p>
</body>
</html>

*/
