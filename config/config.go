package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string   `yaml:"version"`
	Gohtml  []GoHTML `yaml:"gohtml"`
}

type GoHTML struct {
	Templates string `yaml:"templates"`
	Gen       Gen    `yaml:"gen"`
}

type Gen struct {
	Package                string `yaml:"package"`
	OutputFilesSuffix      string `yaml:"output_files_suffix"`
	OutputTemplateFileName string `yaml:"output_template_file_name"`
}

func Parse(path string) (*Config, error) {
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
	for _, gohtml := range cfg.Gohtml {
		if gohtml.Gen.OutputFilesSuffix == "" {
			gohtml.Gen.OutputFilesSuffix = ".go"
		}
	}

	return &cfg, nil
}
