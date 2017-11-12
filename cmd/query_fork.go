package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var queryForkCmd = &cobra.Command{
	Use:   "fork [id]",
	Short: "Fork a query from an existing one",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		queryStrings := url.Values{}
		queryStrings.Set("api_key", apiKey)

		res, err := redash.ForkQuery(redashUrl, id, queryStrings)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryForkCmd)
}
