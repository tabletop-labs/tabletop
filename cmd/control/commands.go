package control

import (
	"log"
	"sync"

	"github.com/fremantle-industries/tabletop/pkg/control"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run the control plane server",
	RunE: func(cmd *cobra.Command, args []string) error {
		wg := sync.WaitGroup{}

		uiPort := "9090"
		ui := control.NewUIServer(uiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(ui.ListenAndServe())
		}()

		apiPort := "9091"
		api := control.NewAPIServer(apiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(api.ListenAndServe())
		}()

		wg.Wait()
		return nil
	},
}
