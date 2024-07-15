# GOHTML
Generate type-safe wrapper code for html text templates.
Eg:

Input: `partials/hello.gohtml`
```gotemplate
<p>Hello {{.Name}}</p>
<p>Age: {{.Age}}</p>
```

Output: `partials/hello.gohtml.go`
```go
type HelloData struct {
	Name any
	Age  any
}

// Hello renders the "Hello" template as an HTML fragment
func Hello(data HelloData) template.HTML{
	...
}

// RenderHello renders the "Hello" template to the specified writer
func RenderHello(w io.Writer, data HelloData) error{
	...
}
```

Which can be used like this:
```go
func Handler(w http.ResponseWriter, r *http.Request){
  // render a fragment - based on example template above
  body := partials.Hello(partials.HelloData{
    Name: "Hello",
    Age:  123,
  })
  
  // you could compose your templates like this 
  // and write the output directly to a writer
  err := partials.RenderPage(w, partials.PageData{
    Title:  "Hello world",
    Navbar: partials.Navbar(partials.NavbarData{SignedIn: true}),
    Body:   body,
    Footer: partials.Footer(),
  })
}
```

# Install CLI
```sh
go install github.com/fritzkeyzer/gohtml/cmd/gohtml@latest
```

# Usage
```sh
gohtml -d 'path/to/templates'
```
# V0
This tool and library are still in development.
Versions prior to v1 have no compatibility guarantees.

# TODO
- [ ] Test more edge cases
- [ ] Prune unused types / fix generation issue
- [ ] YAML config
- [ ] Option to specify generated suffix
- [ ] Support templating JS: *.gojs
- [ ] Support for remaining text template spec

# Contributing
Feel free to post issues - or if you're able to - fix it and submit a PR!

# Changelog

### V0.0.1
Initial version supports:
- variables and nested variables, eg: `{{ .Name }}` or `{{ .User.Location.City }}`
- conditionals eg: `{{ if .IsSignedIn }} ... {{ else }} ... {{ end }}`
- loops eg:
    ```gotemplate
    {{range $link := .Socials}}
      <a href="{{ $link.Href }}">{{ $link.Name }}</li>
    {{end}}
    ```
Known issues:
- When generating types and a loop is involved - an unused type is generated. The failing test: `TestTemplate_Generate/tests/person.gohtml` captures this issue.
