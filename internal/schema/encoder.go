package schema

import (
	"context"

	avro "github.com/awaken/avro/v2"
	"google.golang.org/protobuf/proto"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

type Encoder interface {
	Encode(ctx context.Context, event domain.Event, route domain.Route) (domain.EncodedEvent, error)
}

type ProtobufEncoder struct {
	Registry  Registry
	Prototype proto.Message
}

func (e *ProtobufEncoder) Encode(ctx context.Context, event domain.Event, route domain.Route) (domain.EncodedEvent, error) {
	// TODO(student): map into the generated message, marshal, and add schema headers.
	return domain.EncodedEvent{}, domain.ErrNotImplemented
}

type AvroEncoder struct {
	Registry Registry
	Schema   avro.Schema
}

func (e *AvroEncoder) Encode(ctx context.Context, event domain.Event, route domain.Route) (domain.EncodedEvent, error) {
	// TODO(student): load schema, map and validate datum, serialize, and add schema headers.
	return domain.EncodedEvent{}, domain.ErrNotImplemented
}
