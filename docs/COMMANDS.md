# Tabletop/Commands

[Getting Started](./GETTING_STARTED.md) | [Commands](./COMMANDS.md)

```bash
CLI interface for tabletop, the modern data toolkit for infinite real-time streams.

Realize the Kappa architecture with infinite cloud storage and tools to ingest, process, deliver,
test and monitor your microservices.

Usage:
  tabletop [command]

Available Commands:
  apps        Commands that manage applications
  completion  Generate the autocompletion script for the specified shell
  connectors  Commands that manage connectors
  consumers   Commands that manage consumers
  control     Commands that manage the control plane UI & API
  generate    Commands that generate components
  help        Help about any command
  new         Create a new project
  processors  Commands that manage processors
  producers   Commands that manage producers
  start       Start all applications
  stop        Stop all applications
  streams     Commands that manage streams
  topics      Commands that manage topics required to operate tabletop

Flags:
      --config string   config file (default is $HOME/.tabletop.yaml)
  -h, --help            help for tabletop
  -v, --version         version for tabletop
      --viper           use Viper for configuration (default true)

Use "tabletop [command] --help" for more information about a command.
```

```bash
> tabletop new myapp --stack default
```

```bash
> tabletop generate compose
```

```bash
> tabletop generate k8s
```

```bash
> tabletop generate control --theme winter
```

```bash
> tabletop generate connector coinbase --plugin websocket
```

```bash
> tabletop generate connector binance --plugin httppoll
```

```bash
> tabletop generate processor normalizeprice --plugin stream
```

```bash
> tabletop generate processor 7dayavg --plugin sql
```

```bash
> tabletop generate valet price --plugin duckdb/graphql
```

```bash
> tabletop generate valet 7dayavg --plugin octosql/graphql
```

```bash
> tabletop generate command pullinstruments
```

```bash
> tabletop pipelines
```

```bash
> tabletop connectors
```

```bash
> tabletop processors
```

```bash
> tabletop valets
```

```bash
> tabletop commands
```

```bash
> tabletop topics
```

```bash
> tabletop plugins
```
