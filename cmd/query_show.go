package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryShowCmd = &cobra.Command{
	Use:   "show [id]",
	Short: "Show a query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()

		values := url.Values{}
		values.Set("api_key", apiKey)

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		res, err := goreq.Request{
			Method:      "GET",
			Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
			QueryString: values,
		}.Do()

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

		query, err := js.Get("query").String()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(strings.TrimSpace(query))
	},
}

func init() {
	queryShowCmd.Flags().Bool("json", false, "Dump as JSON")
	queryCmd.AddCommand(queryShowCmd)
}
