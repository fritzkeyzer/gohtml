# GOHTML
Generate type-safe wrapper code for html text templates.

> More documentation can be found in the `docs` directory

# Features
- Functions to render templates
  - To string (technically template.HTML)
  - To io.Writer
  - To http.ResponseWriter
- Generated types for each template based on static analysis. Supports: 
  - variables
  - loops
  - conditionals
- Hot-reloading for local development
  - Templates are loaded from disk allowing for immediate changes to reflect
  - Types can be re-generated whenever changes are detected, by using a file watcher to call `gohtml`
- Error handling
  - Customisable error handler with default

# Install
```sh
go install github.com/fritzkeyzer/gohtml/cmd/gohtml@v0.1.0
```

# Usage
1. Create a `gohtml.yaml` config file, eg:
   ```yaml
    version: "0"
    directories:
    - path: "app/pages"
    - path: "app/components"
    ```
2. Run `gothml` optionally use the `-c` flag to specify a file (defaults to `gohtml.yaml` in the same directory)

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

// RenderHello renders the "Hello" template to a writer
func RenderHello(w io.Writer, data HelloData) error
```

# V0
This tool and library are still in development.
Versions prior to v1 have no compatibility guarantees.

# Roadmap
- [ ] Go type annotations!
- [x] Define multiple template components per file
- [x] Use components from other components with typed variables
- [x] Cache parsed templates
- [x] Option to specify generated file
- [x] YAML config
- [x] Support `$` root context selector
- [x] RenderHTTP, with error handling

# Contributing
Feel free to post issues - or if you're able to - fix it and submit a PR!

# Changelog

### v0.1.0 (⚠️ minor breaking changes)
- Fix generation for conditionals with operators (not, eq, etc)
- Define multiple template components per file (sub-templates can be reused within the same package)
  - Create: `{{define "component"}} ... {{end}}`
  - Reuse: `{{template "component"}}`
  - Use data: `{{template "person" .PersonData}}`
- Generate to a single file: `gohtml.gen.go`
- LiveReload with env var: `GOHTML_LIVERELOAD` 
  - Can be set manually if needed eg: `views.LiveReload = (env == "local")`)
- CLI: improved debug logs
- Updated logic for naming generated loop structs
- Added a golden file test, that tests the entire `tests` directory

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
