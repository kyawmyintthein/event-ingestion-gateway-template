package validation

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestValidateExerciseIsExplicitlyUnimplemented(t *testing.T) {
	err := (EnvelopeValidator{}).Validate(context.Background(), domain.Event{})
	if !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("got %v, want ErrNotImplemented", err)
	}
}
