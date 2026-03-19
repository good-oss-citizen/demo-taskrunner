package tests

import (
	"os"
	"testing"

	"github.com/good-oss-citizen/demo-taskrunner/pkg/config"
)

func TestParseValidConfig(t *testing.T) {
	content := []byte(`
tasks:
  - name: backup
    schedule: "0 2 * * *"
    command: "/usr/bin/backup.sh"
    timeout: "5m"
`)
	f, err := os.CreateTemp("", "config-*.yml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	if _, err := f.Write(content); err != nil {
		t.Fatal(err)
	}
	f.Close()

	cfg, err := config.Parse(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cfg.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(cfg.Tasks))
	}
	if cfg.Tasks[0].Name != "backup" {
		t.Errorf("expected task name 'backup', got '%s'", cfg.Tasks[0].Name)
	}
}

func TestParseFileNotFound(t *testing.T) {
	_, err := config.Parse("/nonexistent/path.yml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
