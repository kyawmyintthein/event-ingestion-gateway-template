package routing

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestResolveExerciseIsExplicitlyUnimplemented(t *testing.T) {
	_, err := (&ConfigRouter{}).Resolve(context.Background(), domain.Event{EventType: "user.created"})
	if !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("got %v, want ErrNotImplemented", err)
	}
}

func TestExtractKeyExerciseIsExplicitlyUnimplemented(t *testing.T) {
	_, err := ExtractKey([]byte(`{"user_id":"u-1"}`), "payload.user_id")
	if !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("got %v, want ErrNotImplemented", err)
	}
}
