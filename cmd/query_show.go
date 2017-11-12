package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var queryShowCmd = &cobra.Command{
	Use:   "show [id]",
	Short: "Show a query",
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

		res, err := redash.GetQuery(redashUrl, id, queryStrings)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		body, err := res.Body.ToString()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		flagJson, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if flagJson {
			fmt.Println(body)
			return
		}

		query := getQueryFromResponseBody(body)

		fmt.Println(strings.TrimSpace(query))
	},
}

func init() {
	queryShowCmd.Flags().Bool("json", false, "Dump as JSON")
	queryCmd.AddCommand(queryShowCmd)
}
