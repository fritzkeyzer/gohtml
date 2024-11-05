package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/fritzkeyzer/gohtml"
	"github.com/fritzkeyzer/gohtml/logz"
)

var (
	cfgFlag     = flag.String("c", "gohtml.yaml", "config file")
	verboseFlag = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()

	cfg, err := gohtml.ParseCfg(*cfgFlag)
	if err != nil {
		panic(err)
	}

	if *verboseFlag {
		logz.Level = slog.LevelDebug
	}

	logz.Debug("Parsed config", "cfg", cfg)

	if len(cfg.Dirs) == 0 {
		logz.Error(nil, "no dirs configured", "cfg", cfg)
		os.Exit(1)
	}

	for _, dir := range cfg.Dirs {

		err = run(dir)
		if err != nil {
			logz.Error(err, "failed on directory", "directory", dir)
			os.Exit(1)
		}
	}
}

func run(dir gohtml.Dir) (err error) {
	g, err := gohtml.ParseDir(dir.Path)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	f, err := os.Create(path.Join(dir.Path, dir.OutputFileName))
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer f.Close()
	err = g.Generate(f)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	logz.Info("Output file: " + path.Join(dir.Path, dir.OutputFileName))

	return nil
}
