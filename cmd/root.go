package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "redashman",
	Short: "redashman is a query management tool for Redash",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	queryCmd.PersistentFlags().String("api-key", "", "A help for foo")
	queryCmd.PersistentFlags().String("url", "", "URL")
	queryCmd.MarkPersistentFlagRequired("api-key")
	queryCmd.MarkPersistentFlagRequired("url")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		checkError(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".redashman")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
