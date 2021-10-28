package connectors

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured connectors",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> connectors list")
		return nil
	},
}
