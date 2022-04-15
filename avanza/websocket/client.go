package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"go.uber.org/atomic"
)

const (
	DefaultWebsocketUrl = "wss://www.avanza.se/_push/cometd"
	PingChannel         = "/meta/connect"
)

type WebsocketClient struct {
	*client.Client
	Dialer       *websocket.Dialer
	messageCount int
	clientID     string
}

func (w *WebsocketClient) Connect(params string) (*Conn, error) {
	conn, err := w.dial()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to websocket server: %v", err)
	}
	if err := conn.authenticate(); err != nil {
		return nil, fmt.Errorf("failed to authenticate: %v", err)
	}
	if err := conn.subscribe(params); err != nil {
		return nil, fmt.Errorf("failed to subscribe to feed %s: %v", params, err)
	}
	return conn, nil
}

func (w *WebsocketClient) dial() (*Conn, error) {
	if w.Client.AuthTokens == nil {
		return nil, fmt.Errorf("you need to authenticate before using websockets")
	}
	conn, _, err := w.Dialer.Dial(DefaultWebsocketUrl, nil)
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn:               conn,
		messageCount:       atomic.NewInt64(1),
		pushSubscriptionID: w.Client.AuthTokens.PushSubscriptionId,
	}, nil
}

type ChanStreamQuotes chan models.StreamQuote

func (w *WebsocketClient) StreamQuotes(ctx context.Context, params string) (*Conn, ChanStreamQuotes, error) {
	conn, err := w.Connect(params)
	if err != nil {
		return nil, nil, err
	}
	data := make(chan []byte, 10000)
	go conn.Collect(data)

	var quotes = make(ChanStreamQuotes, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(quotes)
				return
			case msgBytes := <-data:
				var quote []models.StreamQuote
				err := json.Unmarshal(msgBytes, &quote)
				if err != nil {
					println("failed to unmarshal")
					continue // ignore malformed data
				}
				for _, q := range quote {
					println(q.Channel)
					quotes <- q
				}
			}
		}
	}()
	return conn, quotes, nil
}
