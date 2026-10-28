package kafka

import (
	"context"
	"time"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/ingestion"
	kafkago "github.com/segmentio/kafka-go"
)

type PublisherConfig struct {
	Brokers      []string
	WriteTimeout time.Duration
}

type SyncPublisher struct{ writer *kafkago.Writer }

func NewSyncPublisher(cfg PublisherConfig) *SyncPublisher {
	return &SyncPublisher{writer: &kafkago.Writer{Addr: kafkago.TCP(cfg.Brokers...), WriteTimeout: cfg.WriteTimeout}}
}

func (p *SyncPublisher) Publish(ctx context.Context, events []domain.EncodedEvent) []ingestion.PublishResult {
	return p.PublishSync(ctx, events)
}

func (p *SyncPublisher) PublishSync(ctx context.Context, events []domain.EncodedEvent) []ingestion.PublishResult {
	// TODO(student): write messages and preserve per-message outcomes; a batch is not atomic.
	results := make([]ingestion.PublishResult, len(events))
	for i := range results {
		results[i] = ingestion.PublishResult{Index: i, Err: domain.ErrNotImplemented}
	}
	return results
}

func (p *SyncPublisher) Close() error { return p.writer.Close() }
