package tests

import (
	"testing"

	"github.com/good-oss-citizen/demo-taskrunner/pkg/config"
)

func TestValidateYAMLSyntax_Empty(t *testing.T) {
	err := config.ValidateYAMLSyntax([]byte{})
	if err == nil {
		t.Fatal("expected error for empty input")
	}
}

func TestValidateYAMLSyntax_Tabs(t *testing.T) {
	err := config.ValidateYAMLSyntax([]byte("tasks:\n\t- name: test"))
	if err == nil {
		t.Fatal("expected error for tab character")
	}
}

func TestValidateYAMLSyntax_Valid(t *testing.T) {
	err := config.ValidateYAMLSyntax([]byte("tasks:\n  - name: test"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
