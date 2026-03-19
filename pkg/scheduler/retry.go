package scheduler

import (
	"math"
	"math/rand"
	"time"
)

// BackoffStrategy defines how retry delays are calculated.
type BackoffStrategy int

const (
	// FixedBackoff uses a constant delay between retries.
	FixedBackoff BackoffStrategy = iota
	// ExponentialBackoff doubles the delay with each retry.
	ExponentialBackoff
	// ExponentialWithJitter adds random jitter to exponential backoff.
	ExponentialWithJitter
)

// RetryConfig holds retry behavior configuration.
type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Strategy    BackoffStrategy
}

// DefaultRetryConfig returns sensible retry defaults.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts: 3,
		BaseDelay:   time.Second,
		MaxDelay:    30 * time.Second,
		Strategy:    ExponentialWithJitter,
	}
}

// Delay calculates the delay for a given attempt number.
func (c RetryConfig) Delay(attempt int) time.Duration {
	switch c.Strategy {
	case FixedBackoff:
		return c.BaseDelay
	case ExponentialBackoff:
		delay := time.Duration(math.Pow(2, float64(attempt))) * c.BaseDelay
		if delay > c.MaxDelay {
			return c.MaxDelay
		}
		return delay
	case ExponentialWithJitter:
		delay := time.Duration(math.Pow(2, float64(attempt))) * c.BaseDelay
		if delay > c.MaxDelay {
			delay = c.MaxDelay
		}
		jitter := time.Duration(rand.Int63n(int64(delay / 2)))
		return delay/2 + jitter
	default:
		return c.BaseDelay
	}
}
