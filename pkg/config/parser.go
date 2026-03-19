package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Task represents a scheduled task configuration.
type Task struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
	Command  string `yaml:"command"`
	Timeout  string `yaml:"timeout,omitempty"`
	Retries  int    `yaml:"retries,omitempty"`
}

// Config holds the full configuration.
type Config struct {
	Tasks []Task `yaml:"tasks"`
}

// Parse reads and parses a YAML config file.
//
// BUG: Does not validate that required fields (name, schedule, command)
// are present. A Task with empty name or command will be silently accepted.
// See issue #1.
func Parse(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
