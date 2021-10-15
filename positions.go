package goavanza

import (
	"encoding/json"
	"fmt"

	"github.com/monaco-io/request"
)

func (a *Client) GetPositions() (*Positions, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	body, _, err := a.request("/_mobile/account/positions", request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var positions Positions
	if err := json.Unmarshal([]byte(body), &positions); err != nil {
		return nil, err
	}
	a.logger.Info(positions.InstrumentPositions[0].Positions[0].Name)

	return &positions, nil
}
