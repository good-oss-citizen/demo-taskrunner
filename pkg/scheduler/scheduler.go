package scheduler

import (
	"context"
	"os/exec"
	"time"

	"github.com/good-oss-citizen/demo-taskrunner/pkg/config"
)

// Scheduler runs tasks on their configured schedules.
type Scheduler struct {
	tasks []config.Task
}

// New creates a scheduler from the given tasks.
func New(tasks []config.Task) *Scheduler {
	return &Scheduler{tasks: tasks}
}

// RunOnce executes a task immediately.
//
// BUG: Timeout handling doesn't respect context cancellation.
// When a task exceeds its timeout, we call cancel() but don't wait for
// the goroutine to actually finish. This can cause resource leaks.
// See issue #2.
func (s *Scheduler) RunOnce(ctx context.Context, task config.Task) error {
	timeout, _ := time.ParseDuration(task.Timeout)
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", task.Command)
	return cmd.Run()
}
