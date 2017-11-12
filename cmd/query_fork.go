package cmd

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var queryForkCmd = &cobra.Command{
	Use:   "fork [id]",
	Short: "Fork a query from an existing one",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, err := getUrlFlag()
		checkError(err)
		apiKey, err := getApiKeyFlag()
		checkError(err)

		id, err := strconv.Atoi(args[0])
		checkError(err)

		queryStrings := url.Values{}
		queryStrings.Set("api_key", *apiKey)

		res, err := redash.ForkQuery(*redashUrl, id, queryStrings)
		checkError(err)

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryForkCmd)
}
