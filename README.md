# Event Ingestion Gateway

A beginner-to-advanced Go learning template for accepting JSON event batches over HTTP and publishing schema-encoded records to Kafka. The source intentionally contains `TODO(student)` exercises and `ErrNotImplemented` placeholders; it is scaffolding, not a finished gateway. Do not expect ingestion requests to succeed until the exercises are completed.

## Prerequisites

- Go 1.27.3 (or the minimum toolchain declared in `go.mod`)
- Docker with Compose v2
- `protoc` and `protoc-gen-go` for regenerating Protobuf types
- `golangci-lint` for `make lint`

## Start local Kafka

```sh
make kafka-up
make topics
cp .env.example .env
make run
```

Kafka listens at `localhost:9092` for the host-run app and `kafka:19092` for clients inside the Compose network. The single-node broker and replication factor one are for learning only. Topic creation is an explicit developer operation (`make topics`); production topics are managed externally. The gateway must never create topics dynamically.

The server exposes `GET /healthz` for liveness and `GET /readyz` for readiness. `POST /v1/events` remains HTTP 501 while its exercise is unimplemented. See [docs/api.md](docs/api.md).

## Learn in levels

Work through [exercises/01-http.md](exercises/01-http.md) to [exercises/06-reliability.md](exercises/06-reliability.md). The test suite includes checks for unfinished behavior. Unit tests should run without Kafka; integration tests use a fresh Kafka container and require Docker:

```sh
make test
make test-integration
```

Integration tests do not use Compose or a developer's local broker. They provision isolated test topics and clean up their containers. `go test ./...` is the fast compile/unit-test path; some exercise-specific tests are expected to fail until implemented.

## Architecture and delivery semantics

Read [docs/architecture.md](docs/architecture.md), [docs/failure-semantics.md](docs/failure-semantics.md), and the progressive exercises before implementing behavior. In particular, “atomic” is preflight-before-publish, not a Kafka transaction. Async `202 Accepted` means only that an event entered a bounded in-memory queue; a process crash can lose queued events. Retries can create duplicates.

The optional consumer handler is an advanced extension and is kept separate from the required HTTP-to-Kafka path. It should acknowledge/commit only after successful handling and must account for redelivery.

## Useful commands

| Command | Purpose |
| --- | --- |
| `make kafka-up` / `make kafka-down` | Start/stop local KRaft Kafka |
| `make topics` | Create local example topics explicitly |
| `make run` | Run the gateway on the host |
| `make test` | Run unit tests and compile packages |
| `make test-integration` | Run isolated Testcontainers tests (Docker required) |
| `make proto` | Generate Protobuf Go code |
| `make lint` | Run golangci-lint |
