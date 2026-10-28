# 05 — Bounded asynchronous publishing

**Level:** Advanced

Implement a bounded queue, configurable worker count, admission response, retries, and graceful drain with a timeout. `202 Accepted` means only that the local process accepted the item into memory; it does not mean Kafka published it. A process crash can lose queued items. Queue saturation returns `503` in this template and never silently drops work.

Use capped retry attempts, exponential backoff with jitter, error classification, and context cancellation. Retries may duplicate messages. Workers and shutdown must report publish failures and account for in-flight work.

**Acceptance:** deterministic tests cover queue full, accepted versus published status, retries with injected clock/publisher, cancellation, and shutdown drain without timing-sensitive sleeps.
