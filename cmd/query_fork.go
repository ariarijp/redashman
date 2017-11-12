package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/franela/goreq"
	"github.com/spf13/cobra"
)

var queryForkCmd = &cobra.Command{
	Use:   "fork [id]",
	Short: "Fork a query from an existing one",
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
			Method:      "POST",
			Uri:         fmt.Sprintf("%s/api/queries/%d/fork", redashUrl, id),
			QueryString: values,
		}.Do()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(res.Status)
	},
}

func init() {
	queryCmd.AddCommand(queryForkCmd)
}
