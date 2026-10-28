package validation

import (
	"context"
	"errors"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

var ErrInvalidEvent = errors.New("invalid event")

type Validator interface {
	Validate(ctx context.Context, event domain.Event) error
}

type EnvelopeValidator struct{}

func (EnvelopeValidator) Validate(ctx context.Context, event domain.Event) error {
	// TODO(student): validate required fields, timestamp semantics, and object payload.
	return domain.ErrNotImplemented
}
