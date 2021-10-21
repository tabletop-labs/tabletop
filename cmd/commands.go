package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdapps "github.com/fremantle-industries/tabletop/cmd/apps"
	"github.com/fremantle-industries/tabletop/cmd/generate"
	"github.com/fremantle-industries/tabletop/cmd/lakes"
	"github.com/fremantle-industries/tabletop/pkg/apps"
	"github.com/fremantle-industries/tabletop/pkg/generators"
	"github.com/fremantle-industries/tabletop/pkg/project"
)

var (
	newProjectDir string
	nodeName      string
	nodeShortName string

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

	AppsCmd = &cobra.Command{
		Use:   "apps",
		Short: "Commands to manage OTP applications",
	}

	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("to stop press Ctrl-C")

			node, err := apps.StartNode()
			if err != nil {
				return err
			}
			node.Wait()

			return nil
		},
	}
)

func AddCoreCommands(parentCmd *cobra.Command) {
	generate.AddCommands(GenerateCmd)
	parentCmd.AddCommand(GenerateCmd)

	NewCmd.Flags().StringVarP(&newProjectDir, "dir", "d", ".", "create the project in this directory")
	parentCmd.AddCommand(NewCmd)
}

func AddDataCommands(parentCmd *cobra.Command) {
	parentCmd.AddCommand(BrokersCmd)

	lakes.AddCommands(LakesCmd)
	parentCmd.AddCommand(LakesCmd)
}

func AddAppCommands(parentCmd *cobra.Command) {
	cmdapps.AddCommands(AppsCmd)
	parentCmd.AddCommand(AppsCmd)

	StartCmd.Flags().StringVar(&nodeName, "name", "", "node name")
	StartCmd.Flags().StringVar(&nodeShortName, "sname", "", "node short name")
	parentCmd.AddCommand(StartCmd)
}
