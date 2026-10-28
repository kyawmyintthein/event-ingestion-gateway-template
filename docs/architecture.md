# Architecture

The required path is HTTP ingestion to Kafka publication:

```text
HTTP JSON array
    ↓
Gin handler → request/body limits and HTTP response mapping
    ↓
Ingestion service → validation → exact route lookup → key extraction → encoding
    ↓
Publisher interface → synchronous writer or bounded async worker
    ↓
Externally provisioned Kafka topic
```

Packages are separated by responsibility: `api/http` owns transport concerns; `domain` holds event and result types; `validation`, `routing`, and `schema` implement the preparation steps; `ingestion` coordinates them; `kafka` owns publication and the optional consumer extension; `config` loads environment and route configuration. Small interfaces allow deterministic tests with fakes instead of a broker.

## Routing and schemas

`config/routes.yaml` maps an exact event type to a topic, encoding, schema identity, and payload key field. Event type, topic, and schema identity are distinct values. There is no default topic or JSON fallback. The gateway does not provision topics. Per-entity keys are intended to keep records for the same entity on the same partition, subject to Kafka partitioning and partition-count constraints.

The starter wire convention is raw Protobuf bytes or raw Avro binary datum. Kafka headers carry `content-type`, `schema-id`, and `schema-version`. This is a local teaching convention, not Confluent Schema Registry framing. Schemas are checked into `schemas/`; generated Protobuf code belongs under `gen/proto/` and is regenerated using `make proto`.

## Batch and delivery modes

Batch and delivery modes are server configuration, not arbitrary client overrides. Atomic mode validates, routes, and encodes all records before attempting writes. This avoids writes for a batch that fails preflight, but Kafka publication can fail partway through, so it is not all-or-nothing delivery. Partial mode prepares and publishes valid events while returning per-index outcomes.

Sync mode waits for Kafka writer results. Async mode confirms admission to a bounded local queue only. It does not promise durable storage or successful Kafka publication. Queue-full behavior and shutdown drain limits are documented in [failure-semantics.md](failure-semantics.md).

## Optional consumer extension

The consumer handler is an advanced, optional extension; it is not required to run the HTTP gateway. A consumer loop fetches a Kafka message, invokes an injected handler, and commits only after successful handling. A handler error can cause redelivery, so consumers should be idempotent or otherwise account for duplicates. Concurrency and commit ordering must be designed together when per-key ordering matters.
