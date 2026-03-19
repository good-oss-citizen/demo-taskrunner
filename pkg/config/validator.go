package config

import "fmt"

// This function validates the task configuration by checking all required fields
// are present and properly formatted, ensuring robustness and reliability of the
// configuration parsing pipeline.
func ValidateTaskConfig(task Task) error {
	if task.Name == "" {
		return fmt.Errorf("task name is required")
	}
	if task.Schedule == "" {
		return fmt.Errorf("task schedule is required")
	}
	if task.Command == "" {
		return fmt.Errorf("task command is required")
	}
	return nil
}
