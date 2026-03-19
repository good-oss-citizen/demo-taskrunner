package config

import "fmt"

// ValidateYAMLSyntax checks for common YAML syntax issues before parsing.
// Added after reports of confusing error messages from the YAML parser.
func ValidateYAMLSyntax(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("empty configuration file")
	}
	// Check for tab characters (common YAML mistake)
	for i, b := range data {
		if b == '\t' {
			line := 1
			for _, c := range data[:i] {
				if c == '\n' {
					line++
				}
			}
			return fmt.Errorf("tab character found at line %d: YAML requires spaces for indentation", line)
		}
	}
	return nil
}
