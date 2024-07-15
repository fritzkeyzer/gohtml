package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/fritzkeyzer/gohtml/parse"
)

var (
	dirFlag     = flag.String("d", "", "directory to parse")
	fileFlag    = flag.String("f", "", "file to parse")
	verboseFlag = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()
	dir := *dirFlag
	file := *fileFlag
	verbose := *verboseFlag

	if dir == "" && file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var files []string
	var err error
	if file == "" {
		files, err = filepath.Glob(filepath.Join(dir, "*.gohtml"))
		if err != nil {
			panic(err)
		}
	} else {
		files = []string{file}
		dir = filepath.Dir(file)
	}

	for _, file := range files {
		t := parse.MustParseTemplate(file)

		if verbose {
			logObj(t)
		}

		file, err := os.Create(file + ".go")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		t.Generate(file)
		file.Close()
	}

	if len(files) > 0 {
		parse.GohtmlFile(dir)
	}
}

func logObj(v any) {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	e.Encode(v)
}
