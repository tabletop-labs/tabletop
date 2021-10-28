package streams

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured streams",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> streams list")
		return nil
	},
}
