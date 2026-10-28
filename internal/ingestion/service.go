package ingestion

import (
	"context"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/routing"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/schema"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/validation"
)

type BatchMode string
type DeliveryMode string

const (
	BatchAtomic   BatchMode    = "atomic"
	BatchPartial  BatchMode    = "partial"
	DeliverySync  DeliveryMode = "sync"
	DeliveryAsync DeliveryMode = "async"
)

type IngestOptions struct {
	BatchMode    BatchMode
	DeliveryMode DeliveryMode
}

type PublishResult struct {
	Index int
	Err   error
}

type Publisher interface {
	Publish(ctx context.Context, events []domain.EncodedEvent) []PublishResult
	Close() error
}

type Ingestor interface {
	Ingest(ctx context.Context, events []domain.Event, opts IngestOptions) (domain.BatchResult, error)
}

type Service struct {
	Validator validation.Validator
	Router    routing.Router
	Encoders  map[string]schema.Encoder
	Publisher Publisher
}

func (s *Service) Ingest(ctx context.Context, events []domain.Event, opts IngestOptions) (domain.BatchResult, error) {
	// TODO(student): implement atomic preflight and partial per-event orchestration.
	return domain.BatchResult{}, domain.ErrNotImplemented
}
