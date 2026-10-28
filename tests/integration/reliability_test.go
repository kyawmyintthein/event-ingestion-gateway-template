//go:build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	gatewaykafka "github.com/kyawmyintthein/event-ingestion-gateway/internal/kafka"
)

func TestSyncPublisherReportsDeliveryOutcome(t *testing.T) {
	brokers, topic := startKafkaWithTopic(t, "reliability")
	publisher := gatewaykafka.NewSyncPublisher(gatewaykafka.PublisherConfig{Brokers: brokers, WriteTimeout: 5 * time.Second})
	defer publisher.Close()
	results := publisher.PublishSync(context.Background(), []domain.EncodedEvent{{Topic: topic, Key: []byte("k-1"), Value: []byte("record")}})
	if len(results) != 1 {
		t.Fatalf("results count = %d, want 1", len(results))
	}
	if results[0].Err != nil {
		t.Fatalf("publish failed: %v", results[0].Err)
	}
}
