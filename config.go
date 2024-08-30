package gohtml

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string `yaml:"version"`
	Dirs    []Dir  `yaml:"directories"`
}

type Dir struct {
	Path                   string `yaml:"path"`
	PackageName            string `yaml:"package_name"`
	OutputFilesSuffix      string `yaml:"output_files_suffix"`
	OutputTemplateFileName string `yaml:"output_template_file_name"`
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
		if cfg.Dirs[i].OutputFilesSuffix == "" {
			cfg.Dirs[i].OutputFilesSuffix = ".go"
		}
		if cfg.Dirs[i].OutputTemplateFileName == "" {
			cfg.Dirs[i].OutputTemplateFileName = "gohtml.gen.go"
		}
		if cfg.Dirs[i].PackageName == "" {
			cfg.Dirs[i].PackageName = filepath.Base(cfg.Dirs[i].Path)
		}
	}

	// apply relative paths from config dir
	cfgDir := filepath.Dir(path)
	for i := range cfg.Dirs {
		cfg.Dirs[i].Path = filepath.Join(cfgDir, cfg.Dirs[i].Path)
	}

	return &cfg, nil
}
