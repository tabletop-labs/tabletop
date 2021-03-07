package apps

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps list")

			return nil
		},
	}

	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps start")

			return nil
		},
	}

	StopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps stop")

			return nil
		},
	}
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		ListCmd,
		StartCmd,
		StopCmd,
	)
}
