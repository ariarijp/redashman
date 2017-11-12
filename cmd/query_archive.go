package cmd

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var queryArchiveCmd = &cobra.Command{
	Use:   "archive [id]",
	Short: "Archive a query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		queryStrings := url.Values{}
		queryStrings.Set("api_key", apiKey)

		res, err := redash.ArchiveQuery(redashUrl, id, queryStrings)
		if err != nil {
			panic(err)
		}

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryArchiveCmd)
}
