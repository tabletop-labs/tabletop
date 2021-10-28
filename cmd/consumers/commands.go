package consumers

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured consumers",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> consumers list")
		return nil
	},
}
