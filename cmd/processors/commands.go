package processors

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured processors",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> processors list")
		return nil
	},
}
