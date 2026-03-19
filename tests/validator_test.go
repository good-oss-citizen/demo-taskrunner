package tests

import (
	"testing"
	"github.com/good-oss-citizen/demo-taskrunner/pkg/config"
)

func TestEmptyName(t *testing.T) {
	task := config.Task{Schedule: "* * * * *", Command: "echo hi"}
	if err := config.ValidateTaskConfig(task); err == nil {
		t.Fatal("should fail")
	}
}

func TestEmptySchedule(t *testing.T) {
	task := config.Task{Name: "test", Command: "echo hi"}
	if err := config.ValidateTaskConfig(task); err == nil {
		t.Fatal("should fail")
	}
}

func TestValid(t *testing.T) {
	task := config.Task{Name: "test", Schedule: "* * * * *", Command: "echo hi"}
	if err := config.ValidateTaskConfig(task); err != nil {
		t.Fatal(err)
	}
}
