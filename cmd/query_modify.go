package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryModifyCmd = &cobra.Command{
	Use:   "modify [id]",
	Short: "Modify a query",
	Args:  cobra.ExactArgs(1),
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

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		js := simplejson.New()
		js.Set("query", string(query))

		body, err := js.Encode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		res, err := goreq.Request{
			Method:      "POST",
			Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
			QueryString: values,
			Body:        body,
		}.Do()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryModifyCmd)
}
