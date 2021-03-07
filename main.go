package main

import (
	"context"

	"github.com/fremantle-industries/tabletop/cmd"
)

func main() {
	cmd.AddCommands()
	cmd.Execute(context.Background())
}
