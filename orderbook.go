package goavanza

import (
	"encoding/json"
	"fmt"
	"github.com/monaco-io/request"
	"strings"
)

func (a *Client) GetOrderbook(id string, instrument Instrument) (*OrderbookResponse, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	body, _, err := a.request(fmt.Sprintf("/_mobile/order/%s?orderbookId=%s", strings.ToLower(string(instrument)), id), request.GET, nil, RequestOptions{})
	if err != nil {
		return nil, err
	}
	var orderbookResponse OrderbookResponse
	if err := json.Unmarshal([]byte(body), &orderbookResponse); err != nil {
		return nil, err
	}

	return &orderbookResponse, nil
}

