package routing

import (
	"context"
	"errors"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

var ErrUnknownEventType = errors.New("unknown event type")

type Router interface {
	Resolve(ctx context.Context, event domain.Event) (domain.Route, error)
}

type ConfigRouter struct{ Routes []domain.Route }

func (r *ConfigRouter) Resolve(ctx context.Context, event domain.Event) (domain.Route, error) {
	// TODO(student): exact-match the configured event type and validate its route.
	return domain.Route{}, domain.ErrNotImplemented
}

func ExtractKey(payload []byte, field string) ([]byte, error) {
	// TODO(student): safely traverse the configured payload path and encode a stable Kafka key.
	return nil, domain.ErrNotImplemented
}
