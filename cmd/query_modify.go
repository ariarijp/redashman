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
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)
		backupDir, err := cmd.Flags().GetString("backup-dir")
		checkError(err)

		query, err := ioutil.ReadAll(os.Stdin)
		checkError(err)

		id, err := strconv.Atoi(args[0])
		checkError(err)

		queryStrings := getDefaultQueryStrings(*apiKey)

		if backupDir != "" {
			err = makeBackupFile(*redashUrl, id, queryStrings, backupDir)
			checkError(err)
		}

		json := simplejson.New()
		json.Set("query", string(query))

		res, err := redash.ModifyQuery(*redashUrl, id, queryStrings, json)
		checkError(err)
		checkStatusCode(res)

		fmt.Println(res.Status)
	},
}

func makeBackupFile(redashUrl string, id int, queryStrings url.Values, backupDir string) error {
	res, err := redash.GetQuery(redashUrl, id, queryStrings)
	if err != nil {
		return err
	}

	body, err := res.Body.ToString()
	if err != nil {
		return err
	}

	query, err := getQueryFromResponseBody(body)
	if err != nil {
		return err
	}

	now := time.Now()
	backupFileName := fmt.Sprintf("query_%d_%s.sql", id, now.Format("20060102150405"))
	backupFilePath := path.Join(backupDir, backupFileName)

	err = ioutil.WriteFile(backupFilePath, []byte(*query), 0644)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	queryModifyCmd.Flags().String("backup-dir", "", "Backup file path")
	queryCmd.AddCommand(queryModifyCmd)
}
