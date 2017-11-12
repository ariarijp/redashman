package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new query with text from STDIN",
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)

		query, err := ioutil.ReadAll(os.Stdin)
		checkError(err)

		queryStrings := getDefaultQueryStrings(*apiKey)

		json := simplejson.New()
		json.Set("query", string(query))
		json.Set("data_source_id", 1)
		json.Set("name", "New Query")

		res, err := redash.CreateQuery(*redashUrl, queryStrings, json)
		checkError(err)

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryCreateCmd)
}
