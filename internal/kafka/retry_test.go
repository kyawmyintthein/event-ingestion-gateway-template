package kafka

import (
	"errors"
	"testing"
	"time"
)

func TestRetryPolicyExercisesAreExplicitlyUnimplemented(t *testing.T) {
	policy := RetryPolicy{MaxAttempts: 3, BaseDelay: 10 * time.Millisecond, MaxDelay: time.Second}
	if policy.ShouldRetry(errors.New("temporary"), 1) {
		t.Fatal("ShouldRetry placeholder unexpectedly retried")
	}
	if got := policy.Backoff(1); got != 0 {
		t.Fatalf("Backoff placeholder = %s, want zero", got)
	}
}
