package schema

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestEncoderExercisesAreExplicitlyUnimplemented(t *testing.T) {
	for name, encoder := range map[string]Encoder{"protobuf": &ProtobufEncoder{}, "avro": &AvroEncoder{}} {
		t.Run(name, func(t *testing.T) {
			_, err := encoder.Encode(context.Background(), domain.Event{}, domain.Route{})
			if !errors.Is(err, domain.ErrNotImplemented) {
				t.Fatalf("got %v, want ErrNotImplemented", err)
			}
		})
	}
}
