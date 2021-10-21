package cmd

import (
	"context"

	"github.com/fremantle-industries/tabletop/pkg"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "tabletop",
		Long: `CLI interface for tabletop. A curated selection of tools, libraries and services that
help tame your dataflow to productively build ambitious, data driven & reactive
applications on a streaming lakehouse.

Imagine Rails, Elixir/Erlang OTP & Spring had a baby.`,
		Version: pkg.Version,
	}
)

func Execute(ctx context.Context) {
	cobra.CheckErr(RootCmd.ExecuteContext(ctx))
}
