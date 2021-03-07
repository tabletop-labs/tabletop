package cmd

import (
	"context"

	"github.com/fremantle-industries/tabletop/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tabletop",
		Long: `CLI interface for tabletop. A curated selection of tools, libraries and services that
help tame your dataflow to productively build reactive, elastically scalable and unified
batch/stream applications on a lakehouse.

Imagine Rails, Elixir/Erlang OTP & Spring had a baby.`,
		Version: pkg.Version,
	}
)

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
