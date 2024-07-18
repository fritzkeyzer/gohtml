package gohtml

import (
	"fmt"
	"os"
	"path/filepath"
)

func GohtmlFile(dir string) {
	f, err := os.Create(filepath.Join(dir, "gohtml.gen.go"))
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, `package %[1]s

var LiveReload bool
`,
		filepath.Base(dir),
	)
}
