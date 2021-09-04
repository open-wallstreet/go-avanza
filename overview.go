package goavanza

import (
	"encoding/json"
	"fmt"

	"github.com/monaco-io/request"
)

func (a *api) GetOverview() (*Overview, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	body, _, err := a.request("/_mobile/account/overview", request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var overview Overview
	if err := json.Unmarshal([]byte(body), &overview); err != nil {
		return nil, err
	}

	return &overview, nil
}
