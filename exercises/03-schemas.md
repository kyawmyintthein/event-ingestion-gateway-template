# 03 — Protobuf and Avro encoding

**Level:** Intermediate

Map generic payloads into typed schema-specific records. The Protobuf schema carries event ID, event type, timestamp, and user fields; the Avro schema carries equivalent envelope metadata and typed order fields. Reject mismatched payloads instead of encoding arbitrary JSON.

Commit generated Protobuf Go code and regenerate it with `make proto`; never hand-edit generated files. Load and validate the Avro schema from the repository. Publish raw protobuf bytes or raw Avro binary datum. Add `content-type`, `schema-id`, and `schema-version` Kafka headers. This is not Confluent framing and there is no schema registry.

**Acceptance:** consume records from isolated Kafka and decode them with the matching schema; verify metadata and typed fields.
