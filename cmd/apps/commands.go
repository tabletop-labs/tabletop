package apps

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps list")
			return nil
		},
	}

	StartCmd = &cobra.Command{
		Use:   "start name",
		Short: "Start an application by name",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps start name")
			return nil
		},
	}

	StopCmd = &cobra.Command{
		Use:   "stop name",
		Short: "Stop an application by name",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps stop name")
			return nil
		},
	}
)
