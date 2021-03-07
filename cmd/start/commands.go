package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> start all applications\n")
			return nil
		},
	}
)
