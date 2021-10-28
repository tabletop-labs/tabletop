package cmd

import (
	"fmt"

	"github.com/fremantle-industries/tabletop/cmd/apps"
	"github.com/fremantle-industries/tabletop/cmd/connectors"
	"github.com/fremantle-industries/tabletop/cmd/consumers"
	"github.com/fremantle-industries/tabletop/cmd/control"
	"github.com/fremantle-industries/tabletop/cmd/generate"
	"github.com/fremantle-industries/tabletop/cmd/processors"
	"github.com/fremantle-industries/tabletop/cmd/producers"
	"github.com/fremantle-industries/tabletop/cmd/streams"
	"github.com/fremantle-industries/tabletop/cmd/topics"
	"github.com/fremantle-industries/tabletop/pkg/generators"
	"github.com/fremantle-industries/tabletop/pkg/project"
	"github.com/spf13/cobra"
)

var (
	newProjectDir string

	appsCmd = &cobra.Command{
		Use:   "apps",
		Short: "Commands that manage applications on a node",
	}

	connectorsCmd = &cobra.Command{
		Use:   "connectors",
		Short: "Commands that manage connectors",
	}

	consumersCmd = &cobra.Command{
		Use:   "consumers",
		Short: "Commands that manage consumers",
	}

	controlCmd = &cobra.Command{
		Use:   "control",
		Short: "Commands that manage the control plane UI & API",
		Long:  "Tabletop control plane UI & API for system administration",
	}

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

	processorsCmd = &cobra.Command{
		Use:   "processors",
		Short: "Commands that manage processors",
	}

	producersCmd = &cobra.Command{
		Use:   "producers",
		Short: "Commands that manage producers",
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> start\n")
			return nil
		},
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> stop\n")
			return nil
		},
	}

	streamsCmd = &cobra.Command{
		Use:   "streams",
		Short: "Commands that manage streams",
	}

	topicsCmd = &cobra.Command{
		Use:   "topics",
		Short: "Commands that manage topics required to operate tabletop",
	}
)

func AddCommands() {
	appsCmd.AddCommand(apps.ListCmd)
	appsCmd.AddCommand(apps.StartCmd)
	appsCmd.AddCommand(apps.StopCmd)
	rootCmd.AddCommand(appsCmd)

	connectorsCmd.AddCommand(connectors.ListCmd)
	rootCmd.AddCommand(connectorsCmd)

	consumersCmd.AddCommand(consumers.ListCmd)
	rootCmd.AddCommand(consumersCmd)

	controlCmd.AddCommand(control.StartCmd)
	rootCmd.AddCommand(controlCmd)

	generateCmd.AddCommand(generate.ConsumerCmd)
	generateCmd.AddCommand(generate.ProcessorCmd)
	generateCmd.AddCommand(generate.ProducerCmd)
	rootCmd.AddCommand(generateCmd)

	newCmd.Flags().StringVarP(&newProjectDir, "dir", "d", ".", "create the project in this directory")
	rootCmd.AddCommand(newCmd)

	processorsCmd.AddCommand(processors.ListCmd)
	rootCmd.AddCommand(processorsCmd)

	producersCmd.AddCommand(producers.ListCmd)
	rootCmd.AddCommand(producersCmd)

	streamsCmd.AddCommand(streams.ListCmd)
	rootCmd.AddCommand(streamsCmd)

	topicsCmd.AddCommand(topics.ListCmd)
	rootCmd.AddCommand(topicsCmd)
}
