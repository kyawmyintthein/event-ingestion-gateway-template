package consumer

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

// MessageHandler processes one Kafka record. Implementations should be idempotent
// because an offset can be delivered again after a crash or commit failure.
type MessageHandler interface {
	Handle(ctx context.Context, message kafkago.Message) error
}
