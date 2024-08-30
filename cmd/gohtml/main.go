package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fritzkeyzer/gohtml"
	"github.com/k0kubun/pp/v3"
)

var (
	cfgFlag     = flag.String("c", "gohtml.yaml", "config file")
	verboseFlag = flag.Bool("v", false, "verbose")
	debugFlag   = flag.Bool("debug", false, "debug")
)

func main() {
	flag.Parse()

	if *debugFlag {
		*verboseFlag = true
	}

	cfg, err := gohtml.ParseCfg(*cfgFlag)
	if err != nil {
		panic(err)
	}

	if *verboseFlag {
		pp.Println(cfg)
	}

	if len(cfg.Dirs) == 0 {
		fmt.Println("Config invalid")
		pp.Println(cfg)
		os.Exit(0)
	}

	for _, dir := range cfg.Dirs {
		run(dir)
	}
}

func run(c gohtml.Dir) {
	var files []string
	files, err := filepath.Glob(filepath.Join(c.Path, "*.gohtml"))
	if err != nil {
		panic(err)
	}

	if *debugFlag {
		fmt.Println("Found templates:", c.Path, pp.Sprint(files))
	}

	for _, filePath := range files {
		if *verboseFlag {
			fmt.Println("Parsing file:", filePath)
		}
		gohtml.DebugFlag = *debugFlag
		t, err := gohtml.ParseTemplate(filePath, c.PackageName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		if *verboseFlag {
			fmt.Println("Parsed file:", filePath)
			if *debugFlag {
				pp.Println(t)
				fmt.Println()
			}
		}

		genFilePath := filePath + c.OutputFilesSuffix
		file, err := os.Create(genFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if *verboseFlag {
			fmt.Println("Generating go file:", genFilePath)
		}
		err = t.Generate(file)
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Println("Parsed:", pp.Sprint(t))
			os.Exit(1)
		}
		if *verboseFlag {
			fmt.Println("Successfully generated:", genFilePath)
			//jsonLog(t)
			//fmt.Println()
		}
		file.Close()
	}

	if len(files) > 0 {
		gohtml.GohtmlFile(c.Path, c.OutputTemplateFileName)
	}
}
