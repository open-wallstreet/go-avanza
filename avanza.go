package goavanza

import (
	"github.com/open-wallstreet/go-avanza/avanza/account"
	"github.com/open-wallstreet/go-avanza/avanza/auth"
	"github.com/open-wallstreet/go-avanza/avanza/client"
)

type AvanzaClient struct {
	client.Client
	Auth    *auth.AuthClient
	Account *account.AccountClient
}

func New(opts ...func(a *AvanzaClient)) *AvanzaClient {
	c := client.New()
	a := &AvanzaClient{
		Client:  c,
		Auth:    &auth.AuthClient{Client: c},
		Account: &account.AccountClient{Client: c},
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
