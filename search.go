package goavanza

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/monaco-io/request"
)

func (a *api) Search(query string, instrumentType Instrument) (*Search, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	path := "/_mobile/market/search/%s"
	if instrumentType != ANY {
		path = fmt.Sprintf(path, instrumentType)
	} else {
		path = strings.Replace(path, "/%s", "", 1)
	}
	body, _, err := a.request(path, request.GET, nil, RequestOptions{
		query: map[string]string{
			"limit": "100",
			"query": query,
		},
	})
	if err != nil {
		return nil, err
	}
	var search Search
	if err := json.Unmarshal([]byte(body), &search); err != nil {
		return nil, err
	}
	return &search, nil
}
