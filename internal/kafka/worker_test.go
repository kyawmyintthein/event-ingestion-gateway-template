package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestWorkerExercisesAreExplicitlyUnimplemented(t *testing.T) {
	w := NewWorker(1)
	if err := w.Enqueue(context.Background(), domain.EncodedEvent{}); !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("Enqueue: %v", err)
	}
	if err := w.Run(context.Background()); !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("Run: %v", err)
	}
	if err := w.Shutdown(context.Background()); !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("Shutdown: %v", err)
	}
}

func TestConsumerRunExerciseIsExplicitlyUnimplemented(t *testing.T) {
	c := NewConsumer(ConsumerConfig{Brokers: []string{"localhost:9092"}, Topic: "events", GroupID: "tests"}, nil)
	defer c.Close()
	if err := c.Run(context.Background()); !errors.Is(err, domain.ErrNotImplemented) {
		t.Fatalf("Run: %v", err)
	}
}
