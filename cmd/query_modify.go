package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Songmu/prompter"
	"github.com/ariarijp/redashman/redash"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cobra"
)

var queryModifyCmd = &cobra.Command{
	Use:   "modify [id] [file]",
	Short: "Modify a query with text from file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)
		backupDir, err := cmd.Flags().GetString("backup-dir")
		checkError(err)

		id, err := strconv.Atoi(args[0])
		checkError(err)

		inputFilePath := args[1]
		checkError(err)
		query, err := ioutil.ReadFile(inputFilePath)
		checkError(err)

		if !prompter.YN("Are you sure you want to modify this query?", false) {
			os.Exit(1)
		}

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
