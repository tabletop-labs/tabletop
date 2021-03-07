package lakes

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all data lakes configured in the cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> lakes list\n")
			return nil
		},
	}

	RemoveCmd = &cobra.Command{
		Use:   "remove lakename",
		Short: "Remove a data lake from the cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> lakes remove\n")
			return nil
		},
	}

	SetCmd = &cobra.Command{
		Use:   "set lakename url accesskey secretkey",
		Short: "Set the configuration for a data lake",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> lakes set\n")
			return nil
		},
	}
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		ListCmd,
		RemoveCmd,
		SetCmd,
	)
}
