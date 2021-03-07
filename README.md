# Tabletop

A curated selection of tools, libraries and services that help tame your dataflow to
productively build reactive, elastically scalable and unified batch/stream applications
on a lakehouse.

Imagine [Rails](https://rubyonrails.org), [Elixir](https://elixir-lang.org)/[Erlang OTP](https://www.erlang.org) & [Spring](https://spring.io) had a baby.

[Getting Started](./docs/GETTING_STARTED.md) | [Commands](./docs/COMMANDS.md) | [Architecture](./docs/ARCHITECTURE.md)

## The Default Stack

- [golang](https://go.dev) for glue and custom producers/processors/consumers
- [S3/R2/MinIO](https://en.wikipedia.org/wiki/Amazon_S3) as an API for object storage
- [Grafana](https://github.com/grafana/grafana) for operational visualization that can also be used for BI tasks
- [Loki](https://github.com/grafana/loki) for logs
- [Prometheus](https://github.com/prometheus/prometheus) for metrics
- [Tempo](https://github.com/grafana/tempo) for tracing
- [Kafka](https://kafka.apache.org) as an API with [Redpanda](https://redpanda.com) for low latency message storage and processing
- [Flink](https://flink.apache.org) for unified batch/stream processing
- [Clickhouse](https://github.com/ClickHouse/ClickHouse) as an open cloud lakehouse with support for Postgres/MySQL wire protocols
- [gRPC](https://grpc.io) for service communication
- [Apache APISIX](https://github.com/apache/apisix) cloud-native API gateway
- [Next.JS](https://nextjs.org) for rich web frontends

## What Can I Do? TLDR;

Leverage Kafka as a first class log processing API to realize the [kappa](https://milinda.pathirage.org/kappa-architecture.com)
architecture with infinite cloud storage and tools to productively build reliable, performant
and elastically scalable services.

- ingest
- process
- deliver
- warehouse
- test
- monitor
- time travel
- debug

## Install

Use the `go install` command to download, build and install the `tabletop` binary into the path
set by your `GOBIN` environment variable.

```bash
> go install github.com/fremantle-industries/tabletop@latest
```

## Usage

Initialize a project and `go module` in the current working directory. The `new` command
will generate your project from the default stack of templates along with an application
by the same name.

```bash
> tabletop new github.com/myuser/myapp
creating directory /tmp/myapp
creating directory /tmp/myapp/internal/apps/myapp
creating template /tmp/myapp/README.md
...
```

## Endpoints

| Name                       | Endpoint                                              |
| ---------------------------| ------------------------------------------------------|
| Grafana                    | [`grafana.localhost`](http://grafana.localhost)       |
| Prometheus                 | [`prometheus.localhost`](http://prometheus.localhost) |
| Redpanda Kafka Console     | [`redpanda.localhost`](http://redpanda.localhost)     |
| MinIO Console              | [`minio.localhost`](http://minio.localhost)           |

## Contributing

### Development

Run the default `make` target which downloads dependencies to build and run
`docker compose` on the host machine.

```bash
> make
```

Build the `tabletop` binary

```bash
> make build.bin
```

### Test

Run the `tabletop` test suite

```bash
> make test
```

## Authors

- Alex Kwiatkowski - alex+git@fremantle.io

## License

`tabletop` is released under the [MIT license](./LICENSE.md)
