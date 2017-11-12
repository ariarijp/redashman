package cmd

import (
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
)

func getUrlFlag() (*string, error) {
	flag, err := queryCmd.PersistentFlags().GetString("url")
	if err != nil {
		return nil, err
	}

	url := string(flag)
	return &url, nil
}

func getApiKeyFlag() (*string, error) {
	flag, err := queryCmd.PersistentFlags().GetString("api-key")
	if err != nil {
		return nil, err
	}

	apiKey := string(flag)
	return &apiKey, nil
}

func getQueryFromResponseBody(body string) (*string, error) {
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return nil, err
	}

	query, err := js.Get("query").String()
	checkError(err)

	return &query, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
