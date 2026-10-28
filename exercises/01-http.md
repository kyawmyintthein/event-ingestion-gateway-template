# 01 — HTTP input and validation

**Level:** Beginner

Implement the `POST /v1/events` request path. Accept only a top-level JSON array, enforce the configured body-size and event-count limits, reject empty batches, malformed JSON, and trailing content, and validate each envelope. `payload` must be an object. Require nonempty event ID/type and an RFC3339 timestamp. UUID validation is an optional extension.

Return safe per-index errors and meaningful `400`/`422` responses. Do not return internal errors. The handler should remain transport-focused and delegate validation/ingestion through interfaces.

**Acceptance:** valid batch, malformed/trailing JSON, empty batch, missing fields, non-object payload, too many events, and oversized request cases are covered. Student TODOs should make the relevant tests fail until implemented.
