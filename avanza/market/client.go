package market

import (
	"context"
	"net/http"
	"strings"

	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
)

const (
	GetInstrumentPath = "/_api/market-guide/{instrument}/{id}"
	SearchPath        = "/_mobile/market/search/{instrument}"
	GetMarketDataPath = "/_cqbe/trading/marketdata/{orderbookID}"
	GetChartDataPath  = "/_api/price-chart/stock/{orderbookID}?timePeriod=five_years&resolution=day"
)

type MarketClient struct {
	*client.Client
}

// GetInstrument gets metadata information about a specific instrument
func (a *MarketClient) GetInstrument(ctx context.Context, params *models.GetInstrumentParams, options ...models.RequestOption) (*models.GetInstrumentResponse, error) {
	res := &models.GetInstrumentResponse{}
	params.Instrument = models.InstrumentType(strings.ToLower(string(params.Instrument)))
	err := a.Call(ctx, http.MethodGet, GetInstrumentPath, params, res, options...)
	return res, err
}

// Search for metadata information about tickers, fund or other instruments
func (a *MarketClient) Search(ctx context.Context, params *models.SearchParams, options ...models.RequestOption) (*models.SearchResponse, error) {
	res := &models.SearchResponse{}
	err := a.Call(ctx, http.MethodGet, SearchPath, params, res, options...)
	return res, err
}

// GetMarketData gets the latest market for specific order-book
func (a *MarketClient) GetMarketData(ctx context.Context, params *models.GetMarketDataParams, options ...models.RequestOption) (*models.GetMarketDataResponse, error) {
	res := &models.GetMarketDataResponse{}
	err := a.Call(ctx, http.MethodGet, GetMarketDataPath, params, res, options...)
	return res, err
}

func (a *MarketClient) GetChartData(ctx context.Context, params *models.GetMarketDataParams, options ...models.RequestOption) (*models.GetChartDataResponse, error) {
	res := &models.GetChartDataResponse{}
	err := a.Call(ctx, http.MethodGet, GetChartDataPath, params, res, options...)
	return res, err
}
