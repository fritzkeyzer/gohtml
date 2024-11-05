package gohtml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string `yaml:"version"`
	Dirs    []Dir  `yaml:"directories"`
}

type Dir struct {
	Path           string `yaml:"path"`
	OutputFileName string `yaml:"output_file"`
}

func ParseCfg(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	cfg := Config{}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	// apply defaults
	for i := range cfg.Dirs {
		if cfg.Dirs[i].OutputFileName == "" {
			cfg.Dirs[i].OutputFileName = "gohtml.gen.go"
		}
	}

	return &cfg, nil
}
