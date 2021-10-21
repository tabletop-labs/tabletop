package main

import (
	"context"

	"github.com/fremantle-industries/tabletop/cmd"
)

func main() {
	cmd.AddCoreCommands(cmd.RootCmd)
	cmd.AddDataCommands(cmd.RootCmd)
	cmd.Execute(context.Background())
}
