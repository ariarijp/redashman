package cmd

import (
	"fmt"

	"github.com/ariarijp/redashman/redash"
	"github.com/spf13/cobra"
)

var adminTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Show tasks",
	Run: func(cmd *cobra.Command, args []string) {
		redashUrl, apiKey, err := getRequiredFlags()
		checkError(err)

		queryStrings := getDefaultQueryStrings(*apiKey)

		res, err := redash.GetTasks(*redashUrl, queryStrings)
		checkError(err)

		body, err := res.Body.ToString()
		checkError(err)

		fmt.Println(body)
	},
}

func init() {
	adminCmd.AddCommand(adminTasksCmd)
}
