.PHONY: kafka-up kafka-down topics run test test-integration proto lint

kafka-up:
	docker compose -f deployments/docker-compose.yml up -d

kafka-down:
	docker compose -f deployments/docker-compose.yml down

topics:
	./scripts/create-topics.sh

run:
	set -a; [ ! -f .env ] || . ./.env; set +a; go run ./cmd/gateway

test:
	go test ./...

test-integration:
	go test -tags=integration ./tests/integration/...

proto:
	protoc --proto_path=schemas/proto --go_out=. --go_opt=module=github.com/kyawmyintthein/event-ingestion-gateway schemas/proto/user_event_v1.proto

lint:
	golangci-lint run ./...
