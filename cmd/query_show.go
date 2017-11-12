package cmd

import (
	"fmt"
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
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)

		id, err := strconv.Atoi(args[0])
		checkError(err)

		queryStrings := getDefaultQueryStrings(*apiKey)

		res, err := redash.GetQuery(*redashUrl, id, queryStrings)
		checkError(err)

		body, err := res.Body.ToString()
		checkError(err)

		flagJson, err := cmd.Flags().GetBool("json")
		checkError(err)
		if flagJson {
			fmt.Println(body)
			return
		}

		query, err := getQueryFromResponseBody(body)
		checkError(err)

		fmt.Println(strings.TrimSpace(*query))
	},
}

func init() {
	queryShowCmd.Flags().Bool("json", false, "Dump as JSON")
	queryCmd.AddCommand(queryShowCmd)
}
