package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

func LoadConfig(path string) (*DBConfig, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg DBConfig
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return &cfg, nil
}
