# 06 — Reliability and operations

**Level:** Advanced

Add request IDs, structured logs, counters, and latency measurements. Exercise Kafka unavailable, missing topic, write timeout, schema errors, and duplicate event IDs using injectable dependencies. Explain at-least-once behavior, downstream deduplication, partition ordering, and compatible schema evolution.

Optionally implement a Kafka consumer handler as a separate advanced extension. Inject the handler, process a message, and commit its offset only after success. Decide how to handle poison messages; do not silently discard them. A successful side effect followed by a failed offset commit can cause redelivery.

**Acceptance:** tests are bounded and deterministic, use fake publishers/clocks where appropriate, and integration cases run against their own Testcontainers broker. Document any future registry integration without adding one to this starter.
