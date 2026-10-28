package kafka

import (
	"context"
	"errors"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

var ErrQueueFull = errors.New("publish queue full")

type QueueItem struct {
	Event  domain.EncodedEvent
	Result chan error
}

type Worker struct{ queue chan QueueItem }

func NewWorker(queueSize int) *Worker { return &Worker{queue: make(chan QueueItem, queueSize)} }

func (w *Worker) Enqueue(ctx context.Context, event domain.EncodedEvent) error {
	// TODO(student): enqueue without blocking past context and report queue saturation.
	return domain.ErrNotImplemented
}

func (w *Worker) Run(ctx context.Context) error {
	// TODO(student): publish queued events, report every outcome, and stop on cancellation.
	return domain.ErrNotImplemented
}

func (w *Worker) Shutdown(ctx context.Context) error {
	// TODO(student): stop accepting work and drain the bounded queue within the deadline.
	return domain.ErrNotImplemented
}
