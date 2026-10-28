//go:build integration

package integration

import (
	"fmt"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/testcontainers/testcontainers-go"
	tckafka "github.com/testcontainers/testcontainers-go/modules/kafka"
)

func startKafkaWithTopic(t *testing.T, prefix string) ([]string, string) {
	t.Helper()
	container, err := tckafka.Run(t.Context(), "confluentinc/confluent-local:7.5.0", tckafka.WithClusterID("gateway-test-cluster"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Errorf("terminate Kafka: %v", err)
		}
	})
	brokers, err := container.Brokers(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	topic := fmt.Sprintf("%s_%s", prefix, t.Name())
	admin, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		t.Fatal(err)
	}
	if err := admin.CreateTopics(kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}); err != nil {
		t.Fatal(err)
	}
	if err := admin.Close(); err != nil {
		t.Fatal(err)
	}
	return brokers, topic
}
