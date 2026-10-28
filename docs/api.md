# HTTP API

## `POST /v1/events`

The request body is a top-level JSON array of event envelopes:

```json
[
  {
    "event_id": "7bb72b30-a81e-4e7d-9eaf-b9a19ef8c512",
    "event_type": "user.created",
    "timestamp": "2026-10-28T09:00:00Z",
    "payload": {"user_id": "u-123", "source": "signup"}
  }
]
```

`event_id` and `event_type` must be nonempty, `timestamp` must be RFC3339, and `payload` must be a JSON object. UUID validation is an exercise. Empty arrays, malformed JSON, trailing content, too many events, or oversized bodies are rejected. Limits are configured server-side.

The route and schema configuration determine the destination and encoding. Request fields do not select arbitrary topics or encodings.

### Response behavior to implement

- Sync success: `200`, with `accepted`, `published`, `failed`, and per-index `results`.
- Async queue admission: `202`, with `accepted`, `queued`, and per-index `results`; do not claim the records were published.
- Invalid atomic batch: `400` or `422`, with validation errors and no writes.
- Partial mode: `200` if all publish, `207` for mixed outcomes, `400`/`422` if all are invalid. `207` is nonstandard and clients may need explicit support.
- Infrastructure errors: `503` or `504` as appropriate.

Each result reports the original input index, event ID when available, status, and a safe error code. Never return raw broker errors.

Batch and delivery modes are server-side configuration. Atomic mode means preflight every event before writing, not transactional Kafka delivery. See [failure-semantics.md](failure-semantics.md).

## Health endpoints

- `GET /healthz`: process liveness.
- `GET /readyz`: readiness. Kafka reachability is a later exercise; the initial endpoint may report application readiness without probing Kafka.
- `GET /metrics`: optional advanced exercise.

## Optional consumer

The optional advanced Kafka consumer is not part of this HTTP API. It has no gRPC endpoint or server.
