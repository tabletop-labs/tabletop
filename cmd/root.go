package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fremantle-industries/tabletop/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use: "tabletop",
		Long: `CLI interface for tabletop. Tame your microservice dataflow with a modern data toolkit.

Leverage Kafka as a first class API to realize the Kappa architecture with
infinite cloud storage and tools to productively build reliable, performant
and elastically scalable services.
    `,
		Version: pkg.Version,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tabletop.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tabletop" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tabletop")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
