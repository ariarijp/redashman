package cmd

import (
	"fmt"
	"os"
)

func getUrlFlag() string {
	flag, err := queryCmd.PersistentFlags().GetString("url")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(flag)
}

func getApiKeyFlag() string {
	flag, err := queryCmd.PersistentFlags().GetString("api-key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(flag)
}
