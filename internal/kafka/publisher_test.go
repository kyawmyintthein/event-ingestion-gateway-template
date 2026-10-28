package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestPublishSyncExerciseReturnsExplicitFailures(t *testing.T) {
	p := NewSyncPublisher(PublisherConfig{Brokers: []string{"localhost:9092"}})
	defer p.Close()
	results := p.PublishSync(context.Background(), []domain.EncodedEvent{{}})
	if len(results) != 1 || !errors.Is(results[0].Err, domain.ErrNotImplemented) {
		t.Fatalf("unexpected results: %#v", results)
	}
}
