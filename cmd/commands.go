package cmd

import (
	"github.com/fremantle-industries/tabletop/cmd/generate"
	"github.com/fremantle-industries/tabletop/pkg/generators"
	"github.com/fremantle-industries/tabletop/pkg/project"
	"github.com/spf13/cobra"
)

var (
	newProjectDir string

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Use templates to generate components",
	}

	newCmd = &cobra.Command{
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
	generateCmd.AddCommand(generate.ConsumerCmd)
	generateCmd.AddCommand(generate.ProcessorCmd)
	generateCmd.AddCommand(generate.ProducerCmd)
	rootCmd.AddCommand(generateCmd)

	newCmd.Flags().StringVarP(&newProjectDir, "dir", "d", ".", "create the project in this directory")
	rootCmd.AddCommand(newCmd)
}
