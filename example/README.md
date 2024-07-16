# Example

## Templates
GoHTML `*.gohtml` templates are defined in the `template` and `layout` directories.

## Config
By running the `gohtml` cli in this directory, the `gohtml.yaml` file is used by default.
You can specify a config file using the `-c` flag.

## Generated files
The `*.gohtml.go` files in the two template directories have been generated according to the YAML config.

## Example usage
Check the `main.go` file to see how the generated code can be used. 