package order

import (
	"context"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"net/http"
	"strings"
)

const (
	GetOrderBookPath  = "/_mobile/order/{instrument}"
	GetOrderBooksPath = "/_mobile/market/orderbooklist/{orderbookIds}"
	PlaceOrderPath    = "/_api/order"
	EditOrderPath     = "/_api/order/{instrument}/{id}"
	DeleteOrderPath   = "/_api/order"
)

type OrderClient struct {
	*client.Client
}

func (a *OrderClient) GetOrderBook(ctx context.Context, params *models.GetOrderBookParams, options ...models.RequestOption) (*models.GetOrderBookResponse, error) {
	res := &models.GetOrderBookResponse{}
	params.Instrument = models.InstrumentType(strings.ToLower(string(params.Instrument)))
	err := a.Call(ctx, http.MethodGet, GetOrderBookPath, params, res, options...)
	return res, err
}

func (a *OrderClient) GetOrderBooks(ctx context.Context, params *models.GetOrderBooksParams, options ...models.RequestOption) (*models.GetOrderBooksResponse, error) {
	res := &models.GetOrderBooksResponse{}
	err := a.Call(ctx, http.MethodGet, GetOrderBooksPath, params, res, options...)
	return res, err
}

func (a *OrderClient) PlaceOrder(ctx context.Context, params *models.PlaceOrderParams, options ...models.RequestOption) (*models.PlaceOrderResponse, error) {
	res := &models.PlaceOrderResponse{}
	err := a.Call(ctx, http.MethodPost, PlaceOrderPath, params, res, options...)
	return res, err
}

func (a *OrderClient) EditOrder(ctx context.Context, params *models.EditOrderParams, options ...models.RequestOption) (*models.EditOrderResponse, error) {
	res := &models.EditOrderResponse{}
	err := a.Call(ctx, http.MethodPut, EditOrderPath, params, res, options...)
	return res, err
}

func (a *OrderClient) DeleteOrder(ctx context.Context, params *models.DeleteOrderParams, options ...models.RequestOption) (*models.DeleteOrderResponse, error) {
	res := &models.DeleteOrderResponse{}
	err := a.Call(ctx, http.MethodDelete, DeleteOrderPath, params, res, options...)
	return res, err
}
