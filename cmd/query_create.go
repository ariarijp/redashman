package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Songmu/prompter"
	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryCreateCmd = &cobra.Command{
	Use:   "create [file]",
	Short: "Create a new query with text from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)
		force, _ := cmd.Flags().GetBool("force")

		inputFilePath := args[0]
		checkError(err)
		query, err := ioutil.ReadFile(inputFilePath)
		checkError(err)

		if !force && !prompter.YN("Are you sure you want to create a new query?", false) {
			os.Exit(1)
		}

		queryStrings := getDefaultQueryStrings(*apiKey)

		json := simplejson.New()
		json.Set("query", string(query))
		json.Set("data_source_id", 1)
		json.Set("name", "New Query")

		res, err := redash.CreateQuery(*redashUrl, queryStrings, json)
		checkError(err)
		checkStatusCode(res)

		fmt.Println(res.Status)
	},
}

func init() {
	queryCreateCmd.Flags().BoolP("force", "f", false, "Run without asking")
	queryCmd.AddCommand(queryCreateCmd)
}
