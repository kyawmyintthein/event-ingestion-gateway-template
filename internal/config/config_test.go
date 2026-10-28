package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadReadsEnvironmentThroughViper(t *testing.T) {
	t.Setenv("HTTP_ADDRESS", "127.0.0.1:9090")
	t.Setenv("KAFKA_BROKERS", "broker-a:9092, broker-b:9092")
	t.Setenv("MAX_EVENTS", "17")

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.HTTPAddress != "127.0.0.1:9090" {
		t.Fatalf("HTTPAddress = %q", cfg.HTTPAddress)
	}
	if len(cfg.Brokers) != 2 || cfg.Brokers[1] != "broker-b:9092" {
		t.Fatalf("Brokers = %#v", cfg.Brokers)
	}
	if cfg.MaxEvents != 17 {
		t.Fatalf("MaxEvents = %d, want 17", cfg.MaxEvents)
	}
}

func TestLoadRoutesUsesViperYAML(t *testing.T) {
	path := filepath.Join(t.TempDir(), "routes.yaml")
	data := []byte("routes:\n  - event_type: user.created\n    topic: users.v1\n    encoding: protobuf\n    schema: user_event_v1\n    key_field: payload.user_id\n")
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}

	routes, err := LoadRoutes(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(routes) != 1 || routes[0].EventType != "user.created" || routes[0].KeyField != "payload.user_id" {
		t.Fatalf("routes = %#v", routes)
	}
}
