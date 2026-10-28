package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	httpapi "github.com/kyawmyintthein/event-ingestion-gateway/internal/api/http"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/config"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/ingestion"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/kafka"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/routing"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/schema"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/validation"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}
	routes, err := config.LoadRoutes(cfg.RoutesFile)
	if err != nil {
		slog.Error("load routes", "error", err)
		os.Exit(1)
	}

	broker := kafka.NewSyncPublisher(kafka.PublisherConfig{Brokers: cfg.Brokers, WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second})
	defer broker.Close()
	ingestor := &ingestion.Service{
		Validator: validation.EnvelopeValidator{},
		Router:    &routing.ConfigRouter{Routes: routes},
		Encoders: map[string]schema.Encoder{
			"protobuf": &schema.ProtobufEncoder{Registry: &schema.FileRegistry{Root: "schemas"}},
			"avro":     &schema.AvroEncoder{Registry: &schema.FileRegistry{Root: "schemas"}},
		},
		Publisher: broker,
	}
	handler := &httpapi.Handler{
		Ingestor:     ingestor,
		MaxBodyBytes: cfg.MaxBodyBytes,
		MaxEvents:    cfg.MaxEvents,
		BatchMode:    ingestion.BatchMode(cfg.BatchMode),
		DeliveryMode: ingestion.DeliveryMode(cfg.DeliveryMode),
	}
	router := gin.New()
	router.Use(gin.Recovery(), httpapi.RequestID())
	handler.Register(router)

	srv := &http.Server{Addr: cfg.HTTPAddress, Handler: router, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		slog.Info("gateway listening", "address", cfg.HTTPAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server", "error", err)
			os.Exit(1)
		}
	}()

	stop, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	<-stop.Done()
	ctx, shutdown := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeoutSeconds)*time.Second)
	defer shutdown()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown", "error", err)
	}
}
