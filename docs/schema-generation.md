# Schema generation and wire encoding

The checked-in `.proto` and `.avsc` files are schema sources. Protobuf Go code is generated into `gen/proto/`; do not edit generated files by hand. Install `protoc` and `protoc-gen-go`, then run:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.12
make proto
```

The starter's Kafka value is the raw Protobuf message bytes or raw Avro binary datum. The gateway adds `content-type`, `schema-id`, and `schema-version` headers. This convention is for this learning project; it is not Confluent Schema Registry framing. Schema identity is distinct from the topic name.

The project uses the maintained `github.com/awaken/avro/v2` fork rather than archived `hamba/avro`; applications that decode attacker-controlled Avro payloads must use its bounded map allocation setting. Keep schemas compatible with existing consumers when evolving them: add optional/defaulted fields where possible, avoid changing field meaning or wire type, and test old and new reader/writer combinations.
