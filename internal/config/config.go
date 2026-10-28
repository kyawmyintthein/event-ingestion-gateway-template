package config

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPAddress            string
	Brokers                []string
	RoutesFile             string
	MaxBodyBytes           int64
	MaxEvents              int
	BatchMode              string
	DeliveryMode           string
	WriteTimeoutSeconds    int
	AsyncQueueSize         int
	AsyncWorkers           int
	ShutdownTimeoutSeconds int
}

type routeFile struct {
	Routes []domain.Route `yaml:"routes"`
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddress:            env("HTTP_ADDRESS", ":8080"),
		Brokers:                splitBrokers(env("KAFKA_BROKERS", "localhost:9092")),
		RoutesFile:             env("ROUTES_FILE", "config/routes.yaml"),
		MaxBodyBytes:           envInt64("MAX_BODY_BYTES", 1048576),
		MaxEvents:              envInt("MAX_EVENTS", 100),
		BatchMode:              env("BATCH_MODE", "atomic"),
		DeliveryMode:           env("DELIVERY_MODE", "sync"),
		WriteTimeoutSeconds:    envInt("KAFKA_WRITE_TIMEOUT_SECONDS", 10),
		AsyncQueueSize:         envInt("ASYNC_QUEUE_SIZE", 1000),
		AsyncWorkers:           envInt("ASYNC_WORKERS", 4),
		ShutdownTimeoutSeconds: envInt("SHUTDOWN_TIMEOUT_SECONDS", 10),
	}
	if len(cfg.Brokers) == 0 {
		return Config{}, errors.New("at least one Kafka broker is required")
	}
	if cfg.MaxBodyBytes <= 0 || cfg.MaxEvents <= 0 {
		return Config{}, errors.New("request limits must be positive")
	}
	return cfg, nil
}

func LoadRoutes(path string) ([]domain.Route, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var file routeFile
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, err
	}
	return file.Routes, nil
}

func splitBrokers(raw string) []string {
	var brokers []string
	for _, broker := range strings.Split(raw, ",") {
		if broker = strings.TrimSpace(broker); broker != "" {
			brokers = append(brokers, broker)
		}
	}
	return brokers
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func envInt(key string, fallback int) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil || v < 1 {
		return fallback
	}
	return v
}
func envInt64(key string, fallback int64) int64 {
	v, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil || v < 1 {
		return fallback
	}
	return v
}
