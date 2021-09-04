package goavanza

import (
	"encoding/json"
	"fmt"

	"github.com/monaco-io/request"
)

func (a *api) GetDealsAndOrders() (*DealsAndOrders, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	body, _, err := a.request("/_mobile/account/dealsandorders", request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var dealsAndOrders DealsAndOrders
	if err := json.Unmarshal([]byte(body), &dealsAndOrders); err != nil {
		return nil, err
	}
	return &dealsAndOrders, nil
}
