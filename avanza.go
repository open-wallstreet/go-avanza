package goavanza

import (
	"github.com/open-wallstreet/go-avanza/avanza/account"
	"github.com/open-wallstreet/go-avanza/avanza/auth"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/market"
	"github.com/open-wallstreet/go-avanza/avanza/order"
	"github.com/open-wallstreet/go-avanza/avanza/websocket"
)

type AvanzaClient struct {
	*client.Client
	Auth      *auth.AuthClient
	Account   *account.AccountClient
	Market    *market.MarketClient
	Order     *order.OrderClient
	Websocket *websocket.WebsocketClient
}

func New(opts ...func(a *AvanzaClient)) *AvanzaClient {
	c := client.New()
	a := &AvanzaClient{
		Client:    c,
		Auth:      &auth.AuthClient{Client: c},
		Account:   &account.AccountClient{Client: c},
		Market:    &market.MarketClient{Client: c},
		Order:     &order.OrderClient{Client: c},
		Websocket: &websocket.WebsocketClient{Client: c},
	}
	for _, o := range opts {
		o(a)
	}
	return a
}

func (a *AvanzaClient) Close() {
	a.Auth.Close()
}

func WithDebug(debug bool) func(a *AvanzaClient) {
	return func(a *AvanzaClient) {
		a.HTTP.SetDebug(debug)
	}
}
