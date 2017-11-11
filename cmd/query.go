package cmd

import (
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use: "query",
}

func init() {
	RootCmd.AddCommand(queryCmd)
}
