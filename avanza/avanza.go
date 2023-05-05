// Package avanza is a Go Unoffical API Client for Avanza Bank AB
//
// Please note that I am not affiliated with Avanza Bank AB in any way. The underlying API can be taken down or changed without warning at any point in time.
//
// To install the package simply run
//
//	go get github.com/open-wallstreet/go-avanza
//
// You can create a new client simply like this
//
//	func main() {
//	    client := avanza.New()
//	    defer client.Close()
//	}
//
// Or if you need to debug http responses
//
//	func main() {
//	    client := avanza.New(avanza.WithDebug(true))
//	    defer client.Close()
//	}
package avanza

import (
	"github.com/open-wallstreet/go-avanza/avanza/account"
	"github.com/open-wallstreet/go-avanza/avanza/auth"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/market"
	"github.com/open-wallstreet/go-avanza/avanza/order"
	"github.com/open-wallstreet/go-avanza/avanza/websocket"
)

type Client struct {
	*client.Client
	Auth      *auth.AuthClient
	Account   *account.AccountClient
	Market    *market.MarketClient
	Order     *order.OrderClient
	Websocket *websocket.Client
}

func New(opts ...func(a *Client)) *Client {
	c := client.New()
	a := &Client{
		Client:    c,
		Auth:      auth.NewAuthClient(c),
		Account:   &account.AccountClient{Client: c},
		Market:    &market.MarketClient{Client: c},
		Order:     &order.OrderClient{Client: c},
		Websocket: &websocket.Client{Client: c},
	}
	for _, o := range opts {
		o(a)
	}
	return a
}

func (a *Client) Close() {
	a.Auth.Close()
}

func WithDebug(debug bool) func(a *Client) {
	return func(a *Client) {
		a.HTTP.SetDebug(debug)
	}
}
