# Config
A config file is required for gohtml to target the correct directories.

By default, gohtml will check for a config file named: `gohtml.yaml` in the same directory as the execution of the command.
You can specify a config file, if needed, with the -c flag. 

### Structure
Example with one target directory.
```yaml
version: "0.1.2"
directories:
- path: "path/to/templates/dir"
  output_file: "gohtml.gen.go"               # default value
```

### Example
This example targets 4 directories. All gohtml templates in these directories will be used to generate code.
```yaml
version: "0.1.2"
directories:
- path: "app/page"
- path: "app/partial"
- path: "dashboard/page"
- path: "dashboard/partial"
```
