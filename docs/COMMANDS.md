# Tabletop/Commands

The tabletop CLI is the primary interface to manage your cluster

```bash
> tabletop
CLI interface for tabletop. A curated selection of tools, libraries and services that
help tame your dataflow to productively build reactive, elastically scalable and unified
batch/stream applications on a lakehouse.

Imagine Rails, Elixir/Erlang OTP & Spring had a baby.

Usage:
  tabletop [command]

Available Commands:
  brokers     List cluster brokers
  completion  Generate the autocompletion script for the specified shell
  generate    Use templates to generate components
  help        Help about any command
  lakes       Commands to manage data lakes available to the cluster
  new         Create a new tabletop project

Flags:
  -h, --help      help for tabletop
  -v, --version   version for tabletop

Use "tabletop [command] --help" for more information about a command.
```

## brokers

```bash
> tabletop brokers
TODO> brokers
```

## lakes

```bash
> tabletop lakes
Commands to manage data lakes available to the cluster

Usage:
  tabletop lakes [command]

Available Commands:
  list        List all data lakes configured in the cluster
  remove      Remove a data lake from the cluster
  set         Set the configuration for a data lake

Flags:
  -h, --help   help for lakes

Use "tabletop lakes [command] --help" for more information about a command.
```
