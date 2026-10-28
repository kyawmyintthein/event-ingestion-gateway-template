//go:build integration

package integration

import (
	"context"
	"testing"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/routing"
)

func TestConfiguredRouteAndEntityKey(t *testing.T) {
	_, topic := startKafkaWithTopic(t, "route")
	router := &routing.ConfigRouter{Routes: []domain.Route{{EventType: "user.created", Topic: topic, Encoding: "protobuf", Schema: "user_event_v1", KeyField: "payload.user_id"}}}
	route, err := router.Resolve(context.Background(), domain.Event{EventType: "user.created"})
	if err != nil {
		t.Fatal(err)
	}
	if route.Topic != topic {
		t.Fatalf("route topic = %q, want %q", route.Topic, topic)
	}
	key, err := routing.ExtractKey([]byte(`{"user_id":"u-42"}`), route.KeyField)
	if err != nil {
		t.Fatal(err)
	}
	if string(key) != "u-42" {
		t.Fatalf("key = %q, want u-42", key)
	}
}
