package cmd

import (
	"context"

	"{{.Module}}/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "{{.ProjectName}}",
		Version: pkg.Version,
	}
)

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
