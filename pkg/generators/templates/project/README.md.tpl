# {{.ProjectName}}

Welcome to your [tabletop](https://github.com/fremantle-industries/tabletop) instance for `{{.ProjectName}}`.

## Getting Started

Build your binary for all applications

```bash
> make build.bin
```

Start the applications in your binary

```bash
> make start.bin
```

## Test

Run the `{{.ProjectName}}` test suite

```bash
> make test
```

## Endpoints

| Name                       | Endpoint                                              |
| ---------------------------| ------------------------------------------------------|
| Grafana                    | [`grafana.localhost`](http://grafana.localhost)       |
| Prometheus                 | [`prometheus.localhost`](http://prometheus.localhost) |
| Redpanda Kafka Console     | [`redpanda.localhost`](http://redpanda.localhost)     |
| MinIO Console              | [`minio.localhost`](http://minio.localhost)           |
