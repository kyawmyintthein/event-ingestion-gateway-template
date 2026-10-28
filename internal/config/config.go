package config

import (
	"errors"
	"strings"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/spf13/viper"
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

func Load() (Config, error) {
	v, err := newViper()
	if err != nil {
		return Config{}, err
	}
	cfg := Config{
		HTTPAddress:            v.GetString("http.address"),
		Brokers:                splitBrokers(v.GetString("kafka.brokers")),
		RoutesFile:             v.GetString("routes.file"),
		MaxBodyBytes:           v.GetInt64("request.max_body_bytes"),
		MaxEvents:              v.GetInt("request.max_events"),
		BatchMode:              v.GetString("ingestion.batch_mode"),
		DeliveryMode:           v.GetString("ingestion.delivery_mode"),
		WriteTimeoutSeconds:    v.GetInt("kafka.write_timeout_seconds"),
		AsyncQueueSize:         v.GetInt("async.queue_size"),
		AsyncWorkers:           v.GetInt("async.workers"),
		ShutdownTimeoutSeconds: v.GetInt("shutdown.timeout_seconds"),
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
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var routes []domain.Route
	if err := v.UnmarshalKey("routes", &routes); err != nil {
		return nil, err
	}
	return routes, nil
}

func newViper() (*viper.Viper, error) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	for key, envName := range map[string]string{
		"http.address":                "HTTP_ADDRESS",
		"kafka.brokers":               "KAFKA_BROKERS",
		"kafka.write_timeout_seconds": "KAFKA_WRITE_TIMEOUT_SECONDS",
		"routes.file":                 "ROUTES_FILE",
		"request.max_body_bytes":      "MAX_BODY_BYTES",
		"request.max_events":          "MAX_EVENTS",
		"ingestion.batch_mode":        "BATCH_MODE",
		"ingestion.delivery_mode":     "DELIVERY_MODE",
		"async.queue_size":            "ASYNC_QUEUE_SIZE",
		"async.workers":               "ASYNC_WORKERS",
		"shutdown.timeout_seconds":    "SHUTDOWN_TIMEOUT_SECONDS",
	} {
		if err := v.BindEnv(key, envName); err != nil {
			return nil, err
		}
	}
	v.SetDefault("http.address", ":8080")
	v.SetDefault("kafka.brokers", "localhost:9092")
	v.SetDefault("kafka.write_timeout_seconds", 10)
	v.SetDefault("routes.file", "config/routes.yaml")
	v.SetDefault("request.max_body_bytes", 1048576)
	v.SetDefault("request.max_events", 100)
	v.SetDefault("ingestion.batch_mode", "atomic")
	v.SetDefault("ingestion.delivery_mode", "sync")
	v.SetDefault("async.queue_size", 1000)
	v.SetDefault("async.workers", 4)
	v.SetDefault("shutdown.timeout_seconds", 10)
	return v, nil
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
