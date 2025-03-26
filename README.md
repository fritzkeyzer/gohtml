# GoHTML
Generate type-safe (ish) binding code for Go HTML templates.

[![Version](https://img.shields.io/badge/version-v0.1.5-blue.svg)](https://github.com/fritzkeyzer/gohtml/tags)

> Take a look at the `example` directory for a full example or `tests` for a range of supported features

## Key Features

### 🚀 Generated render functions
  - Generate function to render templates partials as `template.HTML`
  - Generate functions to render templates to `io.Writer` or `http.ResponseWriter`

### 🔒❓ Type-safe (ish)
  - Generate data structs (props) for all templates and sub-templates
  - Supports variables, loops, conditionals and sub-templates via static analysis.

### 💻 Developer Experience
  - Hot-reloading: templates are loaded from disk during development, so changes reflect immediately
  - IDE support: Use well-defined, well-supported languages. HTML, CSS, JS and Go text templates
  - AI support: LLMs are incredibly familiar with HTML and Go text templates  
  - Compile time errors for invalid templates or usage of templates

## Installation
```sh
go install github.com/fritzkeyzer/gohtml/cmd/gohtml@latest
```

## Quick start

1. Create `gohtml.yaml`:
```yaml
version: "0"
directories:
  - path: "app/pages"
  - path: "app/components"
```
2. Run generator
```shell
gohtml generate
```

## Example
`person.gohtml`:
```gotemplate
<h1>Person</h1>
<p>Name: {{.Name}}</p>
<p>Age: {{.Age}}</p>
<p>Phone: {{.Contact.Phone}}</p>
<p>Email: {{.Contact.Email}}</p>
<div>{{range $link := .Socials}}
  <a href="{{$link.Href}}">{{$link.Name}}</a>{{end}}
</div>

```
Generated code:
```go
type PersonContact struct {
  Phone any
  Email any
}

type PersonData struct {
  Name    any
  Age     any
  Contact *PersonContact
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

// RenderPersonHTTP renders "person.gohtml" to an http.ResponseWriter
func RenderPersonHTTP(w http.ResponseWriter, data PersonData) error {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    return tmpl().ExecuteTemplate(w, "person.gohtml", data)
}
```

> 💡Look at the `tests` and `example` directories for more advanced examples

## Development status
⚠️ **Version 0.x.x**: API may change before v1.0

### Roadmap / ideas
- [ ] Additional installation options (other than go install ...)
- [ ] Go type annotations
- [x] Multiple components per file
- [x] Component reuse with typed variables
- [x] Template caching
- [x] Configurable output location
- [x] YAML configuration
- [x] Root context selector support
- [x] HTTP rendering with error handling

### Known bugs
- Check the [issues](https://github.com/fritzkeyzer/gohtml/issues) tab

## Contributing
Issues and PRs welcome! 

Please report errors and if it's possible, create a test case for your error and submit a PR.
I would greatly appreciate it.

# Changelog

### v0.1.5
- Fix empty template directory generating invalid go code

### v0.1.4
- Update CLI with `gohtml generate` command and additional help and version options.

### v0.1.3
- Fix range elements with only {{.}} children
- Fix missing else branches
- Use pointers for nested data to support passing nils
- Improve type detection when a variable appears multiple times eg, {{if .Data}} {{template "Component" .Data}} {{end}}

### v0.1.2
- No longer generate functions for empty templates 
(eg: files that only have sub-template definitions)

### v0.1.1
- Fix deeply nested template directories causing bad generation

### v0.1.0
- Fix generation for conditionals with operators (not, eq, etc)
- Define multiple template components per file (sub-templates can be reused within the same package)
    - Create: `{{define "component"}} ... {{end}}`
    - Reuse: `{{template "component"}}`
    - Use data: `{{template "person" .PersonData}}`
- Generate a single file per directory: `gohtml.gen.go`
- LiveReload with env var: `GOHTML_LIVERELOAD`
    - Can be set manually if needed eg: `views.LiveReload = (env == "local")`)
- CLI: improved debug logs
- Updated logic for naming generated loop structs
- Added a golden file test, that tests the entire `tests` directory
- Removed RenderHTTP error handler

### v0.0.6, v0.0.7
- Support `$` root context selector
- Fix variables nested within conditionals bug
- Add RenderHTTP function with configurable error handler

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