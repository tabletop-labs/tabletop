package main

import (
  "context"

  "{{.Module}}/cmd"
)

func main() {
  cmd.AddCommands()
  cmd.Execute(context.Background())
}
