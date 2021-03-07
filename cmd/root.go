package cmd

import (
	"context"

	"github.com/fremantle-industries/tabletop/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tabletop",
		Long: `CLI interface for tabletop. Tame your microservice dataflow with a modern data toolkit.

Leverage Kafka as a first class API to realize the Kappa architecture with
infinite cloud storage and tools to productively build reliable, performant
and elastically scalable services.
    `,
		Version: pkg.Version,
	}
)

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
