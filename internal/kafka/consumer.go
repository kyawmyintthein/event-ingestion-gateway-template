package kafka

import (
	"context"
	"time"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/consumer"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	kafkago "github.com/segmentio/kafka-go"
)

type ConsumerConfig struct {
	Brokers        []string
	Topic          string
	GroupID        string
	CommitInterval time.Duration
}

type Consumer struct {
	reader  *kafkago.Reader
	handler consumer.MessageHandler
}

func NewConsumer(cfg ConsumerConfig, handler consumer.MessageHandler) *Consumer {
	return &Consumer{reader: kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: cfg.Brokers, Topic: cfg.Topic, GroupID: cfg.GroupID, CommitInterval: cfg.CommitInterval,
	}), handler: handler}
}

func (c *Consumer) Run(ctx context.Context) error {
	// TODO(student, advanced): FetchMessage, call the handler, and commit only on success.
	return domain.ErrNotImplemented
}

func (c *Consumer) Close() error { return c.reader.Close() }
