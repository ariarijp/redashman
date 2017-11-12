package cmd

import (
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use: "admin",
}

func init() {
	RootCmd.AddCommand(adminCmd)
}
