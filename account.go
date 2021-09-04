package goavanza

import (
	"encoding/json"
	"fmt"

	"github.com/monaco-io/request"
)

func (a *api) GetAccountOverview(accountId string) (*AccountOverview, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	body, _, err := a.request(fmt.Sprintf("/_mobile/account/%s/overview", accountId), request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var overview AccountOverview
	if err := json.Unmarshal([]byte(body), &overview); err != nil {
		return nil, err
	}
	return &overview, nil
}
