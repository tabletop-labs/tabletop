# Tabletop

A curated selection of tools, libraries and services that help tame your dataflow to
productively build reactive, elastically scalable and unified batch/stream applications
on a lakehouse.

Imagine [Rails](https://rubyonrails.org), [Elixir](https://elixir-lang.org)/[Erlang OTP](https://www.erlang.org) & [Spring](https://spring.io) had a baby.

## The Default Stack

- [golang](https://go.dev) for glue and custom producers/processors/consumers
- [S3/R2/MinIO](https://en.wikipedia.org/wiki/Amazon_S3) as an API for object storage
- [Grafana](https://github.com/grafana/grafana) for operational visualization that can also be used for BI tasks
- [Loki](https://github.com/grafana/loki) for logs
- [Prometheus](https://github.com/prometheus/prometheus) for metrics
- [Kafka](https://kafka.apache.org) as an API with [Redpanda](https://redpanda.com) for low latency message storage and processing

## What Can I Do? TLDR;

Leverage Kafka as a first class log processing API to realize the [Kappa](https://milinda.pathirage.org/kappa-architecture.com)
architecture with infinite cloud storage and tools to productively build reliable, performant
and elastically scalable services.

- ingest
- process
- warehouse
- deliver
- test
- monitor
- time travel
- debug

## Sections

[Getting Started](./docs/GETTING_STARTED.md) | [Commands](./docs/COMMANDS.md)

## Install

Use the `go install` command to download, build and install the `tabletop` binary into the path
set by your `GOBIN` environment variable.

```bash
> go install github.com/fremantle-industries/tabletop
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

Build a single binary for all applications

```bash
> make build.bin
```

Start the applications in your binary

```bash
> make start.bin
```

## Endpoints

| Name                       | Endpoint                                              |
| ---------------------------| ------------------------------------------------------|
| Tabletop Control Plane UI  | [`tabletop.localhost`](http://tabletop.localhost)     |
| Tabletop Control Plane API | [`api.localhost`](http://api.localhost)               |
| Grafana                    | [`grafana.localhost`](http://grafana.localhost)       |
| Prometheus                 | [`prometheus.localhost`](http://prometheus.localhost) |
| Redpanda Kafka Console     | [`redpanda.localhost`](http://redpanda.localhost)     |
| MinIO Console              | [`minio.localhost`](http://minio.localhost)           |

## Contributing

### Development

Create `tabletop` configuration file

```bash
> make configure
```

Run the default `make` target which downloads dependencies to build and run
`docker compose` on the host machine.

```bash
> make
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
