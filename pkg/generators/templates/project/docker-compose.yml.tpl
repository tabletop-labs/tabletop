version: '3.8'

services:
  minio:
    image: minio/minio:latest
    volumes:
      - minio_data:/data
    environment:
      - MINIO_ROOT_USER=${MINIO_ROOT_USER:-admin}
      - MINIO_ROOT_PASSWORD=${MINIO_ROOT_PASSWORD:-password}
    command: server --console-address ":9001" /data

  loki:
    build:
      context: ./docker/services/loki

  prometheus:
    build:
      context: ./docker/services/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.local.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/usr/share/prometheus/console_libraries'
      - '--web.console.templates=/usr/share/prometheus/consoles'

  # Root node is enabled by setting an empty seed [REDPANDA_SEEDS]
  # - Only 1 root node is allowed in the cluster
  redpanda-broker-1:
    build:
      context: ./docker/services/redpanda_broker
    hostname: redpanda-broker-1
    volumes:
      - "redpanda_broker_1_data:/var/lib/redpanda/data"
    environment:
      - NODE_ID=1
      - ADVERTISE_KAFKA_ADDR_PLAINTEXT_HOST=redpanda-broker-1
      - ADVERTISE_RPC_ADDR_HOST=redpanda-broker-1

  redpanda-broker-2:
    build:
      context: ./docker/services/redpanda_broker
    hostname: redpanda-broker-2
    volumes:
      - "redpanda_broker_2_data:/var/lib/redpanda/data"
    environment:
      - NODE_ID=2
      - REDPANDA_SEEDS=redpanda-broker-1:33145
      - ADVERTISE_KAFKA_ADDR_PLAINTEXT_HOST=redpanda-broker-2
      - ADVERTISE_RPC_ADDR_HOST=redpanda-broker-2

  redpanda-broker-3:
    build:
      context: ./docker/services/redpanda_broker
    hostname: redpanda-broker-3
    volumes:
      - "redpanda_broker_3_data:/var/lib/redpanda/data"
    environment:
      - NODE_ID=3
      - REDPANDA_SEEDS=redpanda-broker-1:33145
      - ADVERTISE_KAFKA_ADDR_PLAINTEXT_HOST=redpanda-broker-3
      - ADVERTISE_RPC_ADDR_HOST=redpanda-broker-3

  redpanda-console:
    build:
      context: ./docker/services/redpanda_console
    depends_on:
      - redpanda-broker-1
      - redpanda-broker-2
      - redpanda-broker-3
    environment:
      - KAFKA_BROKERS=redpanda-broker-1:29092,redpanda-broker-2:29092,redpanda-broker-3:29092
      - KAFKA_SCHEMA_REGISTRY_URLS=http://redpanda-broker-1:8081,http://redpanda-broker-2:8081,http://redpanda-broker-3:8081
      - REDPANDA_ADMIN_API_URLS=http://redpanda-broker-1:9644,http://redpanda-broker-2:9644,http://redpanda-broker-3:9644

  clickhouse:
    build:
      context: ./docker/services/clickhouse
    ports:
      - 8123:8123
      - 9000:9000

  grafana:
    build:
      context: ./docker/services/grafana

  reverse-proxy:
    build:
      context: ./docker/services/reverse_proxy
    ports:
      - 80:8080
    environment:
      - REDPANDA_CONSOLE_INTERNAL_HOST=redpanda-console

volumes:
  minio_data:
  prometheus_data:
  redpanda_broker_1_data:
  redpanda_broker_2_data:
  redpanda_broker_3_data:
  grafana_data:
