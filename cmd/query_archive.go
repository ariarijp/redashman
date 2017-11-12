package cmd

import (
	"fmt"
	"strconv"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var queryArchiveCmd = &cobra.Command{
	Use:   "archive [id]",
	Short: "Archive a query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)

		id, err := strconv.Atoi(args[0])
		checkError(err)

		queryStrings := getDefaultQueryStrings(*apiKey)

		res, err := redash.ArchiveQuery(*redashUrl, id, queryStrings)
		checkError(err)

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryArchiveCmd)
}
