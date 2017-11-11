package cmd

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryArchiveCmd = &cobra.Command{
	Use:   "archive [id]",
	Short: "Archive a query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()

		values := url.Values{}
		values.Set("api_key", apiKey)

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		res, err := goreq.Request{
			Method:      "DELETE",
			Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
			QueryString: values,
		}.Do()
		if err != nil {
			panic(err)
		}

		println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryArchiveCmd)
}
