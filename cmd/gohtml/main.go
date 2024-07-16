package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/fritzkeyzer/gohtml/config"
	"github.com/fritzkeyzer/gohtml/parse"
)

var (
	cfgFlag     = flag.String("c", "gohtml.yaml", "config file")
	verboseFlag = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()

	cfg, err := config.Parse(*cfgFlag)
	if err != nil {
		panic(err)
	}

	for _, g := range cfg.Gohtml {
		run(g)
	}
}

func run(c config.GoHTML) {
	var files []string
	files, err := filepath.Glob(filepath.Join(c.Templates, "*.gohtml"))
	if err != nil {
		panic(err)
	}

	for _, filePath := range files {
		t := parse.MustParseTemplate(filePath)

		if *verboseFlag {
			logObj(t)
		}

		genFilePath := filePath + c.Gen.OutputFilesSuffix
		file, err := os.Create(genFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		t.Generate(file)
		file.Close()
	}

	if len(files) > 0 {
		parse.GohtmlFile(c.Templates)
	}
}

func logObj(v any) {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	e.Encode(v)
}
