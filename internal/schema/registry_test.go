package schema

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestLookupExerciseIsExplicitlyUnimplemented(t *testing.T) {
	_, err := (&FileRegistry{Root: "schemas"}).Lookup(context.Background(), "user_event_v1", "1")
	if !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("got %v, want ErrNotImplemented", err)
	}
}
