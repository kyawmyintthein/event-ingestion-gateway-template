# Failure and delivery semantics

## Atomic batches are not Kafka transactions

Atomic mode must validate, resolve routes, extract keys, look up schemas, and encode every event before the first write. A preflight failure means zero writes. After publishing begins, a broker/network failure can occur after some records were delivered. This mode therefore provides all-or-none preflight, not all-or-nothing delivery. Do not describe it as a Kafka transaction.

## Async acceptance is volatile

Async mode returns `202 Accepted` after placing events in a bounded in-memory queue. Acceptance is not publication or durable persistence. Events still queued when the process crashes can be lost. A full queue must return the documented overload response (the template chooses `503 Service Unavailable`) and must not silently drop the event. Graceful shutdown should stop admission, drain up to a configured deadline, and report remaining failures.

## Retries and duplicates

Retries can publish a record even when the producer did not receive confirmation that an earlier attempt succeeded. Delivery is therefore at least once under retry scenarios, and duplicate events are possible. `event_id` can support downstream deduplication, but the gateway does not guarantee deduplication. Retry only classified transient errors, cap attempts, use bounded exponential backoff with jitter, and honor context cancellation.

## Ordering

Kafka preserves order within a partition. A stable entity key helps route related events to the same partition, but ordering is not guaranteed across partitions, across different keys, or when application-level workers reorder writes. Changing a topic's partition count can also change key-to-partition assignment.

## Consumer redelivery

The optional consumer extension should commit an offset only after successful handling. If handling succeeds but committing fails, a message may be handled again. Handlers should be idempotent or use a deduplication strategy appropriate to the downstream side effect. Poison-message handling and dead-letter topics are future extensions, not silent discard paths.

## Broker and schema errors

Map internal failures to safe API codes; never return raw broker details. Missing topics and unavailable brokers must produce observable failures, not implicit topic creation or a JSON encoding fallback. Schema identity and version travel in headers under this local convention; it is not Schema Registry wire framing. Schema evolution should preserve compatibility for existing readers and be reviewed before changing a published schema.
