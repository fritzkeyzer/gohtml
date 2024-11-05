# GoHTML
Generate type-safe(ish) wrapper code for Go HTML templates.

[![Version](https://img.shields.io/badge/version-v0.1.0-blue.svg)](https://github.com/fritzkeyzer/gohtml/tags)

> Detailed documentation available in the `docs` directory

> Take a look at the `example` directory for a full example

## Key Features

### 🚀 Generated render functions
  - Render templates partials `template.HTML`
  - Render templates to `io.Writer` or `http.ResponseWriter`

### 🔒 Type Safety
  - Automatically generated Go types for each template.
  - Supports variables, loops, conditionals and sub-templates via static analysis.

### 💻 Developer Experience
  - Hot-reloading: templates are loaded from disk during development, so changes reflect immediately.
  - Combined with an IDE or external file-watcher, gohtml can re-run upon changes to template files.

## Installation
```sh
go install github.com/fritzkeyzer/gohtml/cmd/gohtml@v0.1.0
```

## Quick start

1. Create `gohtml.yaml`:
```yaml
version: "0.1.0"
directories:
  - path: "app/pages"
  - path: "app/components"
```
2. Run generator
```shell
gohtml
```

## Example
`hello.gohtml`:
```gotemplate
<h1>Hello {{.Name}}</h1>
<p>{{.Message}}</p>
```
Generated code:
```go
type HelloData struct {
    Name    any
    Message any
}

func Hello(data HelloData) template.HTML
func RenderHello(w io.Writer, data HelloData) error
```

## Development status
⚠️ **Version 0.x.x**: API may change before v1.0

### Roadmap
- [ ] Go type annotations
- [x] Multiple components per file
- [x] Component reuse with typed variables
- [x] Template caching
- [x] Configurable output location
- [x] YAML configuration
- [x] Root context selector support
- [x] HTTP rendering with error handling

### Known bugs
- If a .gohtml file **only** contains sub templates, a render function (for the file) is still created even if it will do nothing

## Contributing
Issues and PRs welcome!

# Changelog

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