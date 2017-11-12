package cmd

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
)

func getUrlFlag() string {
	flag, err := queryCmd.PersistentFlags().GetString("url")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(flag)
}

func getApiKeyFlag() string {
	flag, err := queryCmd.PersistentFlags().GetString("api-key")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(flag)
}

func getQueryFromResponseBody(body string) string {
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

	return query
}
