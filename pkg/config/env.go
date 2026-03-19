package config

import (
	"fmt"
	"os"
)

// EnvVar represents an environment variable requirement.
type EnvVar struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
	Default  string `yaml:"default,omitempty"`
}

// ValidateEnvVars checks that all required environment variables are set.
// This is the pattern to follow for config validation.
func ValidateEnvVars(vars []EnvVar) error {
	for _, v := range vars {
		val := os.Getenv(v.Name)
		if val == "" && v.Required && v.Default == "" {
			return fmt.Errorf("required environment variable %s is not set", v.Name)
		}
	}
	return nil
}
