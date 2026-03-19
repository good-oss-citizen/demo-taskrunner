package tests

import (
	"testing"
	"time"

	"github.com/good-oss-citizen/demo-taskrunner/pkg/scheduler"
)

func TestDefaultRetryConfig(t *testing.T) {
	cfg := scheduler.DefaultRetryConfig()
	if cfg.MaxAttempts != 3 {
		t.Errorf("expected 3 max attempts, got %d", cfg.MaxAttempts)
	}
	if cfg.Strategy != scheduler.ExponentialWithJitter {
		t.Errorf("expected ExponentialWithJitter strategy")
	}
}

func TestFixedBackoffDelay(t *testing.T) {
	cfg := scheduler.RetryConfig{
		BaseDelay: time.Second,
		Strategy:  scheduler.FixedBackoff,
	}
	for i := 0; i < 5; i++ {
		if cfg.Delay(i) != time.Second {
			t.Errorf("attempt %d: expected 1s fixed delay, got %v", i, cfg.Delay(i))
		}
	}
}

func TestExponentialBackoffDelay(t *testing.T) {
	cfg := scheduler.RetryConfig{
		BaseDelay: time.Second,
		MaxDelay:  30 * time.Second,
		Strategy:  scheduler.ExponentialBackoff,
	}
	delay0 := cfg.Delay(0)
	delay1 := cfg.Delay(1)
	delay2 := cfg.Delay(2)
	if delay0 >= delay1 || delay1 >= delay2 {
		t.Errorf("expected increasing delays: %v, %v, %v", delay0, delay1, delay2)
	}
}

func TestExponentialBackoffMaxDelay(t *testing.T) {
	cfg := scheduler.RetryConfig{
		BaseDelay: time.Second,
		MaxDelay:  5 * time.Second,
		Strategy:  scheduler.ExponentialBackoff,
	}
	delay := cfg.Delay(10)
	if delay > 5*time.Second {
		t.Errorf("delay %v exceeds max %v", delay, 5*time.Second)
	}
}
