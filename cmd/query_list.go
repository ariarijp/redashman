package cmd

import (
	encjson "encoding/json"
	"fmt"
	"net/url"

	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryListCmd = &cobra.Command{
	Use:   "list [page_size]",
	Short: "List queries",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, err := getUrlFlag()
		checkError(err)
		apiKey, err := getApiKeyFlag()
		checkError(err)
		pageSize := args[0]

		queryStrings := url.Values{}
		queryStrings.Set("api_key", *apiKey)
		queryStrings.Set("page_size", pageSize)

		res, err := redash.GetQueries(*redashUrl, queryStrings)
		checkError(err)

		body, err := res.Body.ToString()
		checkError(err)

		flagJson, err := cmd.Flags().GetBool("json")
		checkError(err)
		if flagJson {
			fmt.Println(body)
			return
		}

		json, err := simplejson.NewJson([]byte(body))
		checkError(err)

		for _, query := range json.Get("results").MustArray() {
			q := query.(map[string]interface{})
			queryUrl := fmt.Sprintf("%s/queries/%s", *redashUrl, q["id"])
			fmt.Printf("%s\t%s\t%s\n",
				q["id"].(encjson.Number).String(),
				q["name"].(string),
				queryUrl,
			)
		}
	},
}

func init() {
	queryListCmd.Flags().Bool("json", false, "Dump as JSON")
	queryCmd.AddCommand(queryListCmd)
}
