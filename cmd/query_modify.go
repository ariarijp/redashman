package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

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
		backupDir, err := cmd.Flags().GetString("backup-dir")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

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

		if backupDir != "" {
			makeBackupFile(redashUrl, id, values, backupDir)
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

func makeBackupFile(redashUrl string, id int, values url.Values, backupDir string) {
	res, err := goreq.Request{
		Method:      "GET",
		Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
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

	query := getQueryFromResponseBody(body)

	now := time.Now()
	backupFileName := fmt.Sprintf("query_%d_%s.sql", id, now.Format("20060102150405"))
	backupFilePath := path.Join(backupDir, backupFileName)

	err = ioutil.WriteFile(backupFilePath, []byte(query), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	queryModifyCmd.Flags().String("backup-dir", "", "Backup file path")
	queryCmd.AddCommand(queryModifyCmd)
}
