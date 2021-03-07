package generate

import (
	"github.com/fremantle-industries/tabletop/pkg/appcontext"
	"github.com/fremantle-industries/tabletop/pkg/generators"
	"github.com/spf13/cobra"
)

var (
	producerContextDir  string
	producerType        string
	processorContextDir string
	processorType       string
	consumerContextDir  string
	consumerType        string

	ProducerCmd = &cobra.Command{
		Use:   "producer app name",
		Short: "Generate a new producer for an application",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generators.Execute(&appcontext.ProducerGenerator{
				ApplicationContext: generators.ApplicationContext{
					App:     args[0],
					Context: producerContextDir,
				},
				Name: args[1],
				Type: producerType,
			})
		},
	}

	ProcessorCmd = &cobra.Command{
		Use:   "processor app name",
		Short: "Generate a new processor for an application",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generators.Execute(&appcontext.ProcessorGenerator{
				ApplicationContext: generators.ApplicationContext{
					App:     args[0],
					Context: processorContextDir,
				},
				Name: args[1],
				Type: processorType,
			})
		},
	}

	ConsumerCmd = &cobra.Command{
		Use:   "consumer name",
		Short: "Generate a new consumer for an application",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generators.Execute(&appcontext.ConsumerGenerator{
				ApplicationContext: generators.ApplicationContext{
					App:     args[0],
					Context: consumerContextDir,
				},
				Name: args[1],
				Type: consumerType,
			})
		},
	}
)

func init() {
	ProducerCmd.Flags().StringVarP(&producerContextDir, "context", "c", "ingest", "directory to add the producer")
	ProducerCmd.Flags().StringVarP(&producerType, "type", "t", "hello_tick", "template for the producer")

	ProcessorCmd.Flags().StringVarP(&processorContextDir, "context", "c", "transform", "directory to add the processor")
	ProcessorCmd.Flags().StringVarP(&processorType, "type", "t", "hello_tick", "template for the processor")

	ConsumerCmd.Flags().StringVarP(&consumerContextDir, "context", "c", "deliver", "directory to add the consumer")
	ConsumerCmd.Flags().StringVarP(&consumerType, "type", "t", "hello_tick", "template for the consumer")
}
