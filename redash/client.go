package redash

import (
	"fmt"
	"net/url"

	"github.com/bitly/go-simplejson"
	"github.com/franela/goreq"
)

func GetQuery(redashUrl string, id int, queryStrings url.Values) (*goreq.Response, error) {
	return goreq.Request{
		Method:      "GET",
		Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
		QueryString: queryStrings,
	}.Do()
}

func GetQueries(redashUrl string, queryStrings url.Values) (*goreq.Response, error) {
	return goreq.Request{
		Method:      "GET",
		Uri:         fmt.Sprintf("%s/api/queries", redashUrl),
		QueryString: queryStrings,
	}.Do()
}

func CreateQuery(redashUrl string, queryStrings url.Values, json *simplejson.Json) (*goreq.Response, error) {
	body, err := json.Encode()
	if err != nil {
		return nil, err
	}

	return goreq.Request{
		Method:      "POST",
		Uri:         fmt.Sprintf("%s/api/queries", redashUrl),
		QueryString: queryStrings,
		Body:        body,
	}.Do()
}

func ModifyQuery(redashUrl string, id int, queryStrings url.Values, json *simplejson.Json) (*goreq.Response, error) {
	body, err := json.Encode()
	if err != nil {
		return nil, err
	}

	return goreq.Request{
		Method:      "POST",
		Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
		QueryString: queryStrings,
		Body:        body,
	}.Do()
}

func ForkQuery(redashUrl string, id int, queryStrings url.Values) (*goreq.Response, error) {
	return goreq.Request{
		Method:      "POST",
		Uri:         fmt.Sprintf("%s/api/queries/%d/fork", redashUrl, id),
		QueryString: queryStrings,
	}.Do()
}

func ArchiveQuery(redashUrl string, id int, queryStrings url.Values) (*goreq.Response, error) {
	return goreq.Request{
		Method:      "DELETE",
		Uri:         fmt.Sprintf("%s/api/queries/%d", redashUrl, id),
		QueryString: queryStrings,
	}.Do()
}
