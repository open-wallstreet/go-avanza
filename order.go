package goavanza

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrderOptions struct {
	AccountId   string
	OrderbookId string
	OrderType   OrderType
	Price       float64
	ValidUntil  time.Time
	Volume      int
}

func (a *Client) PlaceOrder(options *OrderOptions) (*OrderActionResponse, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	if len(options.AccountId) <= 0 {
		return nil, fmt.Errorf("account id needs to be set")
	}
	if len(options.OrderbookId) <= 0 {
		return nil, fmt.Errorf("order book id needs to be set")
	}
	data := map[string]interface{}{
		"accountId":   options.AccountId,
		"orderbookId": options.OrderbookId,
		"orderType":   options.OrderType,
		"price":       options.Price,
		"volume":      options.Volume,
		"validUntil":  options.ValidUntil.Format(layoutISO),
	}
	body, _, err := a.request("/_api/order", "POST", data, RequestOptions{})
	if err != nil {
		a.logger.Error(body)
		return nil, err
	}
	var placeOrder OrderActionResponse
	if err := json.Unmarshal([]byte(body), &placeOrder); err != nil {
		return nil, err
	}
	return &placeOrder, nil
}

func (a *Client) EditOrder(instrumentType Instrument, orderId string, options *OrderOptions) (*OrderActionResponse, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	if len(options.AccountId) <= 0 {
		return nil, fmt.Errorf("account id needs to be set")
	}
	if len(options.OrderbookId) <= 0 {
		return nil, fmt.Errorf("order book id needs to be set")
	}
	data := map[string]interface{}{
		"accountId":   options.AccountId,
		"orderbookId": options.OrderbookId,
		"orderType":   options.OrderType,
		"price":       options.Price,
		"volume":      options.Volume,
		"validUntil":  options.ValidUntil.Format(layoutISO),
	}
	url := fmt.Sprintf("/_api/order/%s/%s", instrumentType, orderId)
	body, _, err := a.request(url, "PUT", data, RequestOptions{})
	if err != nil {
		a.logger.Error(body)
		return nil, err
	}
	var placeOrder OrderActionResponse
	if err := json.Unmarshal([]byte(body), &placeOrder); err != nil {
		return nil, err
	}
	return &placeOrder, nil
}

func (a *Client) DeleteOrder(accountId string, orderId string) (*OrderActionResponse, error) {
	if !a.IsAuthenticated {
		return nil, fmt.Errorf("not authenticated")
	}
	url := fmt.Sprintf("/_api/order?accountId=%s&orderId=%s", accountId, orderId)
	body, _, err := a.request(url, "DELETE", nil, RequestOptions{})
	if err != nil {
		a.logger.Error(body)
		return nil, err
	}
	var placeOrder OrderActionResponse
	if err := json.Unmarshal([]byte(body), &placeOrder); err != nil {
		return nil, err
	}
	return &placeOrder, nil
}
