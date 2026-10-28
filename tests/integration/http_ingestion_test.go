//go:build integration

package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	httpapi "github.com/kyawmyintthein/event-ingestion-gateway/internal/api/http"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/ingestion"
	gatewaykafka "github.com/kyawmyintthein/event-ingestion-gateway/internal/kafka"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/routing"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/schema"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/validation"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/testcontainers/testcontainers-go"
	tckafka "github.com/testcontainers/testcontainers-go/modules/kafka"
)

func TestPostEventsPublishesToKafka(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	container, err := tckafka.Run(ctx, "confluentinc/confluent-local:7.5.0", tckafka.WithClusterID("gateway-test-cluster"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Errorf("terminate Kafka: %v", err)
		}
	})
	brokers, err := container.Brokers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	topic := fmt.Sprintf("platform_users_%d", time.Now().UnixNano())
	admin, err := kafkago.Dial("tcp", brokers[0])
	if err != nil {
		t.Fatal(err)
	}
	if err := admin.CreateTopics(kafkago.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}); err != nil {
		t.Fatal(err)
	}
	if err := admin.Close(); err != nil {
		t.Fatal(err)
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()
	publisher := gatewaykafka.NewSyncPublisher(gatewaykafka.PublisherConfig{Brokers: brokers, WriteTimeout: 10 * time.Second})
	defer publisher.Close()
	route := domain.Route{EventType: "user.created", Topic: topic, Encoding: "protobuf", Schema: "user_event_v1", KeyField: "payload.user_id"}
	ingestor := &ingestion.Service{
		Validator: validation.EnvelopeValidator{},
		Router:    &routing.ConfigRouter{Routes: []domain.Route{route}},
		Encoders:  map[string]schema.Encoder{"protobuf": &schema.ProtobufEncoder{Registry: &schema.FileRegistry{Root: "../../schemas"}}},
		Publisher: publisher,
	}
	handler := &httpapi.Handler{Ingestor: ingestor, MaxBodyBytes: 1 << 20, MaxEvents: 10, BatchMode: ingestion.BatchAtomic, DeliveryMode: ingestion.DeliverySync}
	handler.Register(router)
	server := httptest.NewServer(router)
	defer server.Close()
	body := `[{"event_id":"7bb72b30-a81e-4e7d-9eaf-b9a19ef8c512","event_type":"user.created","timestamp":"2026-10-28T09:00:00Z","payload":{"user_id":"u-123","source":"signup"}}]`
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, server.URL+"/v1/events", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("POST /v1/events status = %d, want 200 (exercise may be unimplemented)", response.StatusCode)
	}

	reader := kafkago.NewReader(kafkago.ReaderConfig{Brokers: brokers, Topic: topic, Partition: 0, MinBytes: 1, MaxBytes: 1 << 20, MaxWait: 5 * time.Second})
	defer reader.Close()
	readCtx, readCancel := context.WithTimeout(ctx, 10*time.Second)
	defer readCancel()
	message, err := reader.FetchMessage(readCtx)
	if err != nil {
		t.Fatal(err)
	}
	if string(message.Key) != "u-123" {
		t.Fatalf("key = %q, want u-123", message.Key)
	}
	if len(message.Value) == 0 {
		t.Fatal("published message has an empty value")
	}
}
