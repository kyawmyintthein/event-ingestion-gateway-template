#!/usr/bin/env bash
set -euo pipefail

# Developer-only provisioning for local Compose. Production topics are managed
# by the platform team; the gateway never creates topics.
compose_file="$(dirname "$0")/../deployments/docker-compose.yml"
bootstrap_server="${KAFKA_INTERNAL_BROKERS:-kafka:19092}"

for topic in platform.users.v1 platform.orders.v1; do
  docker compose -f "$compose_file" exec -T kafka \
    /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server "$bootstrap_server" \
    --create --if-not-exists --topic "$topic" \
    --partitions 3 --replication-factor 1
done
