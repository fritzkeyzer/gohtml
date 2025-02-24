package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/fritzkeyzer/gohtml"
	"github.com/fritzkeyzer/gohtml/logz"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "gohtml",
		Version: "v0.1.5",
		Usage:   "Generate type-safe go bindings for *.gohtml text templates",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "verbose",
				Action: func(ctx context.Context, command *cli.Command, b bool) error {
					if b {
						logz.Level = slog.LevelDebug
					}
					return nil
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "generate",
				Description: "Run the generator",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Value:   "gohtml.yaml",
						Aliases: []string{"c"},
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cfgPath := cmd.String("config")
					cfg, err := gohtml.ParseCfg(cfgPath)
					if err != nil {
						return fmt.Errorf("parse config: %v", err)
					}

					logz.Debug("Parsed config", "cfg", cfg)

					if len(cfg.Dirs) == 0 {
						logz.Error(nil, "invalid config", "cfg", cfg)
						return fmt.Errorf("no dirs configured")
					}

					var errors []string
					for _, dir := range cfg.Dirs {
						logz.Debug("Running on dir", "dir", dir)
						err = run(dir)
						if err != nil {
							logz.Error(err, "failed on directory", "directory", dir)
							errors = append(errors, err.Error())
						}
					}

					if len(errors) > 0 {
						return fmt.Errorf("generate errors: [%v]", strings.Join(errors, "; "))
					}

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		logz.Error(err, "gohtml exited")
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

	logz.Info("gohtml generated: " + path.Join(dir.Path, dir.OutputFileName))

	return nil
}
