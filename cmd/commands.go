package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/fremantle-industries/tabletop/cmd/generate"
	"github.com/fremantle-industries/tabletop/cmd/lakes"
	"github.com/fremantle-industries/tabletop/pkg/generators"
	"github.com/fremantle-industries/tabletop/pkg/project"
)

var (
	newProjectDir string

	BrokersCmd = &cobra.Command{
		Use:   "brokers",
		Short: "List cluster brokers",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> brokers\n")

			return nil
		},
	}

	GenerateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Use templates to generate components",
	}

	LakesCmd = &cobra.Command{
		Use:   "lakes",
		Short: "Commands to manage data lakes available to the cluster",
	}

	NewCmd = &cobra.Command{
		Use:   "new name",
		Short: "Create a new tabletop project",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generators.Execute(&project.ProjectGenerator{
				Module: args[0],
				Dir:    newProjectDir,
			})
		},
	}
)

func AddCommands() {
	rootCmd.AddCommand(BrokersCmd)

	generate.AddCommands(GenerateCmd)
	rootCmd.AddCommand(GenerateCmd)

	lakes.AddCommands(LakesCmd)
	rootCmd.AddCommand(LakesCmd)

	NewCmd.Flags().StringVarP(&newProjectDir, "dir", "d", ".", "create the project in this directory")
	rootCmd.AddCommand(NewCmd)
}
