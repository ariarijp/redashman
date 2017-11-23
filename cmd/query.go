package cmd

import (
	"time"

	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use: "query",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		connTimeout, _ := RootCmd.PersistentFlags().GetUint("timeout")
		goreq.SetConnectTimeout(time.Duration(connTimeout) * time.Millisecond)
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)
}
