# GOHTML
Generate type-safe wrapper code for html text templates.

# Install
```sh
go install github.com/fritzkeyzer/gohtml/cmd/gohtml@latest
```

# Usage
```sh
gohtml
# or specify a config file:
gohtml -c frontend/templates/gohtml.yaml
```

# Example
See a full example in the example directory

`hello.gohtml`
```gotemplate
<p>Hello {{.Name}}</p>
```

Generates: `hello.gohtml.go` that include this:
```go
type HelloData struct{
	Name any
}

// Hello renders the "Hello" template as an HTML fragment
func Hello(data HelloData) template.HTML

// RenderHello renders the "Hello" template to the specified writer
func RenderHello(w io.Writer, data HelloData) error
```


# V0
This tool and library are still in development.
Versions prior to v1 have no compatibility guarantees.

# TODO
- [ ] Cache parsed templates
- [ ] Handle top level (unnamed) variables
- [ ] Prune unused types / fix generation issue
- [ ] Option to specify generated suffix
- [ ] Support templating JS: *.gojs
- [ ] Support for remaining text template spec
- [x] YAML config

# Contributing
Feel free to post issues - or if you're able to - fix it and submit a PR!

# Known issues:
- When generating types and a loop is involved - an unused type is generated. The failing test: `TestTemplate_Generate/tests/person.gohtml` captures this issue.

# Changelog

### v0.0.2
- Added yaml config support. 
- By default, gohtml checks for a file alongside it: `gohtml.yaml` otherwise the config file can be specified with the `-c` flag.
- The `-d` and `-f` flags have been removed in favour of using a config file. 

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
