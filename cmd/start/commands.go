package start

import (
	"fmt"

	"github.com/fremantle-industries/tabletop/pkg/apps"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("to stop press Ctrl-C")

			node, err := apps.StartNode()
			if err != nil {
				return err
			}
			node.Wait()

			return nil
		},
	}
)
