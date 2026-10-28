package ingestion

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestIngestExerciseIsExplicitlyUnimplemented(t *testing.T) {
	_, err := (&Service{}).Ingest(context.Background(), nil, IngestOptions{})
	if !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("got %v, want ErrNotImplemented", err)
	}
}
