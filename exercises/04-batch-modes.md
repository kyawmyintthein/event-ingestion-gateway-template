# 04 — Atomic preflight and partial batches

**Level:** Intermediate

Implement server-configured `batch_mode=atomic|partial`. Atomic mode must validate, route, extract keys, look up schemas, and encode every event before any write. Partial mode publishes valid events and reports each input index and safe error code.

Atomic preflight is not transactional Kafka publishing: a failure during writes can leave some records delivered. Document this in API descriptions and do not claim all-or-nothing delivery.

**Acceptance:** with one invalid event, atomic mode writes zero records; partial mode writes valid records only and reports the invalid index. Test the mid-publish failure limitation separately with a fake publisher.
