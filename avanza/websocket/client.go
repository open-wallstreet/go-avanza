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
	DefaultWebsocketUrl        = "wss://www.avanza.se/_push/cometd"
	PingChannel                = "/meta/connect"
	QuotesSubscriptionPath     = "/quotes/%s"
	OrderDepthSubscriptionPath = "/orderdepths/%s"
	PositionsSubscriptionPath  = "/positions/%s"
	OrdersSubscriptionPath     = "/orders/%s"
	TradesSubscriptionPath     = "/trades/%s"
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
	conn, err := w.Connect(fmt.Sprintf(QuotesSubscriptionPath, params))
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
					quotes <- q
				}
			}
		}
	}()
	return conn, quotes, nil
}

type ChanStreamOrderDepth chan models.StreamOrderDepth

func (w *WebsocketClient) StreamOrderDepth(ctx context.Context, params string) (*Conn, ChanStreamOrderDepth, error) {
	conn, err := w.Connect(fmt.Sprintf(OrderDepthSubscriptionPath, params))
	if err != nil {
		return nil, nil, err
	}
	data := make(chan []byte, 10000)
	go conn.Collect(data)

	var orderDepth = make(ChanStreamOrderDepth, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(orderDepth)
				return
			case msgBytes := <-data:
				var orderDepths []models.StreamOrderDepth
				err := json.Unmarshal(msgBytes, &orderDepths)
				if err != nil {
					println("failed to unmarshal")
					continue // ignore malformed data
				}
				for _, o := range orderDepths {
					orderDepth <- o
				}
			}
		}
	}()
	return conn, orderDepth, nil
}

type ChanStreamPositions chan models.StreamPositions

func (w *WebsocketClient) StreamPositions(ctx context.Context, params string) (*Conn, ChanStreamPositions, error) {
	conn, err := w.Connect(fmt.Sprintf(PositionsSubscriptionPath, params))
	if err != nil {
		return nil, nil, err
	}
	data := make(chan []byte, 10000)
	go conn.Collect(data)

	var positions = make(ChanStreamPositions, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(positions)
				return
			case msgBytes := <-data:
				var streamPositions []models.StreamPositions
				err := json.Unmarshal(msgBytes, &streamPositions)
				if err != nil {
					println("failed to unmarshal")
					continue // ignore malformed data
				}
				for _, p := range streamPositions {
					positions <- p
				}
			}
		}
	}()
	return conn, positions, nil
}

type ChanStreamOrders chan models.StreamOrders

func (w *WebsocketClient) StreamOrders(ctx context.Context, params string) (*Conn, ChanStreamOrders, error) {
	conn, err := w.Connect(fmt.Sprintf(OrdersSubscriptionPath, params))
	if err != nil {
		return nil, nil, err
	}
	data := make(chan []byte, 10000)
	go conn.Collect(data)

	var orders = make(ChanStreamOrders, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(orders)
				return
			case msgBytes := <-data:
				var streamOrders []models.StreamOrders
				err := json.Unmarshal(msgBytes, &streamOrders)
				if err != nil {
					println("failed to unmarshal")
					continue // ignore malformed data
				}
				for _, o := range streamOrders {
					orders <- o
				}
			}
		}
	}()
	return conn, orders, nil
}

type ChanStreamTrades chan models.StreamTrades

func (w *WebsocketClient) StreamTrades(ctx context.Context, params string) (*Conn, ChanStreamTrades, error) {
	conn, err := w.Connect(fmt.Sprintf(TradesSubscriptionPath, params))
	if err != nil {
		return nil, nil, err
	}
	data := make(chan []byte, 10000)
	go conn.Collect(data)

	var trades = make(ChanStreamTrades, 10000)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(trades)
				return
			case msgBytes := <-data:
				var streamTrades []models.StreamTrades
				err := json.Unmarshal(msgBytes, &streamTrades)
				if err != nil {
					println("failed to unmarshal")
					continue // ignore malformed data
				}
				for _, trade := range streamTrades {
					trades <- trade
				}
			}
		}
	}()
	return conn, trades, nil
}
