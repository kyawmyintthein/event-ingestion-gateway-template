# 02 — Routing and synchronous Kafka

**Level:** Beginner to intermediate

Load exact event-type mappings from `config/routes.yaml`. Implement route lookup, unknown-type errors, topic validation, safe extraction of configured key paths such as `payload.user_id`, and synchronous publishing with `kafka-go`. Do not create topics or fall back to a default destination.

Use the key to keep events for one entity on the same partition when possible. Kafka ordering is per partition, not global. Preserve default `kafka-go` producer settings unless the exercise explicitly requires an override.

**Acceptance:** integration tests use isolated Testcontainers Kafka, create two topics in test setup, and assert topic, key, message count, and same-key ordering.
