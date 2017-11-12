package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryModifyCmd = &cobra.Command{
	Use:   "modify [id]",
	Short: "Modify a query with text from STDIN",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl := getUrlFlag()
		apiKey := getApiKeyFlag()
		backupDir, err := cmd.Flags().GetString("backup-dir")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

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

		queryStrings := url.Values{}
		queryStrings.Set("api_key", apiKey)

		if backupDir != "" {
			makeBackupFile(redashUrl, id, queryStrings, backupDir)
		}

		json := simplejson.New()
		json.Set("query", string(query))

		res, err := redash.ModifyQuery(redashUrl, id, queryStrings, json)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(res.Status)
	},
}

func makeBackupFile(redashUrl string, id int, queryStrings url.Values, backupDir string) {
	res, err := redash.GetQuery(redashUrl, id, queryStrings)
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
