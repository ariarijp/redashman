package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new query",
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()

		values := url.Values{}
		values.Set("api_key", apiKey)

		query, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		js := simplejson.New()
		js.Set("query", string(query))
		js.Set("data_source_id", 1)
		js.Set("name", "New Query")

		body, err := js.Encode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		res, err := goreq.Request{
			Method:      "POST",
			Uri:         fmt.Sprintf("%s/api/queries", redashUrl),
			QueryString: values,
			Body:        body,
		}.Do()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryCreateCmd)
}
