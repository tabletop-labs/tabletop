package cmd

import (
	"github.com/fremantle-industries/tabletop/cmd/start"
)

func AddCommands() {
	rootCmd.AddCommand(start.StartCmd)
}
