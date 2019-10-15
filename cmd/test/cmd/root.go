package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile          string
	userLicense      string
	masterAccountKey string
	endpoint         string
	accountType      string
	accountSecret    string

	rootCmd = &cobra.Command{
		Use:   "testher",
		Short: "Testing Tool For Herdius Chain",
		Long:  `CLI to measure the power which comes with Herdius Chain`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.herdius.yaml)")
	rootCmd.PersistentFlags().StringVar(&masterAccountKey, "masteraccount", "ecja1TZR8pWr/5OI5j7km4eg+BXh+RB2oN1d4oo6yZE=", "verbose output")
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "http://localhost", "verbose output")
	accountCmd.Flags().StringVarP(&accountType, "type", "t", "HER", "Account Type HER/BTC")
	accountCmd.Flags().StringVarP(&accountSecret, "secret", "s", "", "Account Type HER/BTC")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			//er(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
