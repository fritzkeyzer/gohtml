# GoHTML Example Project

A minimal demonstration of GoHTML template generation and usage.

## Quick Start

1. Create a `gohtml.yaml` config file in your project root
2. Place your `*.gohtml` templates in the `views` directory
3. Run `gohtml` to generate the code
4. Import and use the generated types in your Go code

## Project Structure
├── gohtml.yaml     # Configuration file
├── views/          # Template directory
│   └── *.gohtml    # Template files
└── main.go         # Usage example


## Configuration

The `gohtml.yaml` file specifies which directories to process:

```yaml
templates:
  - views/          # Process all *.gohtml files in views/
```

To use a different config file:
```shell
gohtml -c path/to/config.yaml
```

## Generated Code
GoHTML generates a `gohtml.gen.go` file containing:

- Type-safe template structs
- Rendering functions


## Usage Example
See main.go for a complete working example of the generated code in action.