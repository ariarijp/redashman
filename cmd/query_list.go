package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/franela/goreq"
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

		values := url.Values{}
		values.Set("api_key", apiKey)
		values.Set("page_size", pageSize)

		res, err := goreq.Request{
			Method:      "GET",
			Uri:         fmt.Sprintf("%s/api/queries", redashUrl),
			QueryString: values,
		}.Do()
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

		js, err := simplejson.NewJson([]byte(body))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, query := range js.Get("results").MustArray() {
			q := query.(map[string]interface{})
			queryUrl := fmt.Sprintf("%s/queries/%s", redashUrl, q["id"])
			fmt.Printf("%s\t%s\t%s\n",
				q["id"].(json.Number).String(),
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
