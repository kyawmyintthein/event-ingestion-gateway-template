package domain

import (
	"encoding/json"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

type Event struct {
	EventID   string          `json:"event_id"`
	EventType string          `json:"event_type"`
	Timestamp time.Time       `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
}

type Route struct {
	EventType string `yaml:"event_type" mapstructure:"event_type" json:"event_type"`
	Topic     string `yaml:"topic" mapstructure:"topic" json:"topic"`
	Encoding  string `yaml:"encoding" mapstructure:"encoding" json:"encoding"`
	Schema    string `yaml:"schema" mapstructure:"schema" json:"schema"`
	KeyField  string `yaml:"key_field" mapstructure:"key_field" json:"key_field"`
}

type EncodedEvent struct {
	Topic   string
	Key     []byte
	Value   []byte
	Headers []kafkago.Header
}
