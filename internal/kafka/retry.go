package kafka

import (
	"time"
)

type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func (p RetryPolicy) ShouldRetry(err error, attempt int) bool {
	// TODO(student): classify retryable errors and enforce the attempt cap.
	return false
}

func (p RetryPolicy) Backoff(attempt int) time.Duration {
	// TODO(student): exponential backoff with bounded jitter.
	return 0
}
