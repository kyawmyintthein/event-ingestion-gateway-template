//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/schema"
)

func TestAvroSchemaEncodingHasMetadata(t *testing.T) {
	_, topic := startKafkaWithTopic(t, "schema")
	event := domain.Event{EventID: "evt-42", EventType: "order.placed", Timestamp: time.Date(2026, 10, 28, 9, 0, 0, 0, time.UTC), Payload: json.RawMessage(`{"order_id":"o-9","user_id":"u-42","total_amount":25.5,"currency":"AUD"}`)}
	encoded, err := (&schema.AvroEncoder{Registry: &schema.FileRegistry{Root: "../../schemas"}}).Encode(context.Background(), event, domain.Route{Topic: topic, Encoding: "avro", Schema: "order_event_v1"})
	if err != nil {
		t.Fatal(err)
	}
	if encoded.Topic != topic || len(encoded.Value) == 0 {
		t.Fatalf("unexpected encoded event: %#v", encoded)
	}
	if len(encoded.Headers) != 3 {
		t.Fatalf("headers = %d, want content-type/schema-id/schema-version", len(encoded.Headers))
	}
}
