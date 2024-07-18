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

# Examples
See a full example in the example directory.
Also take a look at the tests directory as it demonstrates the range input-output capabilities.

`hello.gohtml`
```gotemplate
<h1>Hello {{.Name}}</h1>
<p>{{.Message}}</p>
```

Generates: `hello.gohtml.go` that includes this:
```go
type HelloData struct{
	Name any
	Message any
}

// Hello renders the "Hello" template as an HTML fragment
func Hello(data HelloData) template.HTML

// RenderHello renders the "Hello" template to the specified writer
func RenderHello(w io.Writer, data HelloData) error
```

# V0
This tool and library are still in development.
Versions prior to v1 have no compatibility guarantees.

# Roadmap
- [ ] Support multiple template definitions within one file with. Including usage with args
- [ ] Handle top level (unnamed) variables
- [ ] Support templating JS: *.gojs
- [ ] Support for remaining text template spec
- [x] Cache parsed templates
- [x] Option to specify generated suffix
- [x] YAML config

# Contributing
Feel free to post issues - or if you're able to - fix it and submit a PR!

# Changelog

### v0.0.4, v0.0.5
- Fix generated filepath bug

### v0.0.3
- Simplified config
- Fix superfluous type definitions
- Add more tests, including parsing and generation
- Apply standard go formatting to generated code

### v0.0.2
- Added yaml config support. 
- By default, gohtml checks for a file alongside it: `gohtml.yaml` otherwise the config file can be specified with the `-c` flag.
- The `-d` and `-f` flags have been removed in favour of using a config file.
- Known issue: When generating types and a loop is involved - an unused type is generated. The failing test: `TestTemplate_Generate/tests/person.gohtml` captures this issue.

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
