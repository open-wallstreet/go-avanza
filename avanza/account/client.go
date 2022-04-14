package account

import (
	"context"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"net/http"
)

const (
	OverviewPath = "/_mobile/account/overview"
	AccountOverviewPath = "/_mobile/account/{accountId}/overview"
	GetPositionsPath = "/_mobile/account/positions"
	GetDealsAndOrdersPath = "/_mobile/account/dealsandorders"
)

type AccountClient struct {
	client.Client
}

// GetOverview retrieves overview of all accounts
func (a *AccountClient) GetOverview(ctx context.Context, options ...models.RequestOption) (*models.OverviewResponse, error) {
	res := &models.OverviewResponse{}
	params := &models.OverviewParams{}
	err := a.Call(ctx, http.MethodGet, OverviewPath, params, res, options...)
	return res, err
}

// GetAccountOverview get overview of a specified account.
func (a *AccountClient) GetAccountOverview(ctx context.Context, params *models.AccountOverviewParams, options ...models.RequestOption) (*models.AccountOverviewResponse, error) {
	res := &models.AccountOverviewResponse{}
	err := a.Call(ctx, http.MethodGet, AccountOverviewPath, params, res, options...)
	return res, err
}

// GetPositions gets all positions
func (a *AccountClient) GetPositions(ctx context.Context, options ...models.RequestOption) (*models.GetPositionsResponse, error) {
	res := &models.GetPositionsResponse{}
	params := &models.GetPositionsParams{}
	err := a.Call(ctx, http.MethodGet, GetPositionsPath, params, res, options...)
	return res, err
}

func (a *AccountClient) GetDealsAndOrders(ctx context.Context, options ...models.RequestOption) (*models.GetDealsAndOrdersResponse, error) {
	res := &models.GetDealsAndOrdersResponse{}
	params := &models.GetDealsAndOrdersParams{}
	err := a.Call(ctx, http.MethodGet, GetDealsAndOrdersPath, params, res, options...)
	return res, err
}
