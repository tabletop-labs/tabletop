package producers

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured producers",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> producers list")
		return nil
	},
}
