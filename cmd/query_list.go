package cmd

import (
	encjson "encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryListCmd = &cobra.Command{
	Use:   "list [page_size]",
	Short: "List queries",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()
		pageSize := args[0]

		queryStrings := url.Values{}
		queryStrings.Set("api_key", apiKey)
		queryStrings.Set("page_size", pageSize)

		res, err := redash.GetQueries(redashUrl, queryStrings)
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

		json, err := simplejson.NewJson([]byte(body))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, query := range json.Get("results").MustArray() {
			q := query.(map[string]interface{})
			queryUrl := fmt.Sprintf("%s/queries/%s", redashUrl, q["id"])
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
