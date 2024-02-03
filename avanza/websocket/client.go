package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"
	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"go.uber.org/atomic"
	"gopkg.in/tomb.v2"
	"sync"
	"time"
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

type Client struct {
	*client.Client
	Dialer       *websocket.Dialer
	messageCount int
	clientID     string
	rwTomb       tomb.Tomb
	pTomb        tomb.Tomb
	mutex        sync.Mutex
	backoff      backoff.BackOff
	conn         *Conn
	rQueue       chan json.RawMessage
	wQueue       chan json.RawMessage
	output       chan any
	err          chan error
	shouldClose  bool
}

func NewClient(cl *client.Client) *Client {
	c := &Client{
		Client:  cl,
		backoff: backoff.NewExponentialBackOff(),
		rQueue:  make(chan json.RawMessage, 10000),
		wQueue:  make(chan json.RawMessage, 1000),
		output:  make(chan any, 100_000),
		err:     make(chan error),
	}

	return c
}

func (w *Client) Connect() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.conn != nil {
		return nil
	}

	notify := func(err error, d time.Duration) {
		w.Logger.Errorf("websocket connection failed: %v, retrying in %s", err, d)
	}
	if err := backoff.RetryNotify(w.connect(false), w.backoff, notify); err != nil {
		return err
	}
	return nil
}

func (w *Client) dial() (*Conn, error) {
	w.Logger.Infof("dialing websocket")
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

func (w *Client) Subscribe(path, params string) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	msg := []*models.SubscribeRequest{{
		Subscription: fmt.Sprintf(path, params),
		ChannelMessage: models.ChannelMessage{
			Channel: "/meta/subscribe",
		},
		ClientID: w.conn.clientID,
		ID:       w.conn.messageCount.String(),
	}}
	marshal, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	w.wQueue <- marshal
	return nil
}

type ChanStreamQuotes chan models.StreamQuote

/*
	func (w *Client) StreamQuotes(ctx context.Context, params string) (*Conn, ChanStreamQuotes, error) {
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

	func (w *Client) StreamOrderDepth(ctx context.Context, params string) (*Conn, ChanStreamOrderDepth, error) {
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

	func (w *Client) StreamPositions(ctx context.Context, params string) (*Conn, ChanStreamPositions, error) {
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

	func (w *Client) StreamOrders(ctx context.Context, params string) (*Conn, ChanStreamOrders, error) {
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

	func (w *Client) StreamTrades(ctx context.Context, params string) (*Conn, ChanStreamTrades, error) {
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
*/
func (w *Client) connect(reconnect bool) func() error {
	return func() error {
		conn, err := w.dial()
		if err != nil {
			return err
		}
		w.conn = conn
		w.Logger.Debugf("connected to %s", w.conn.conn.RemoteAddr())
		// reset write queue
		w.wQueue = make(chan json.RawMessage, 10000)
		err = w.conn.authenticate()
		if err != nil {
			return err
		}
		w.Logger.Debugf("authenticated")

		w.rwTomb = tomb.Tomb{}
		w.rwTomb.Go(w.read)
		w.rwTomb.Go(w.write)
		if !reconnect {
			w.pTomb = tomb.Tomb{}
			w.pTomb.Go(w.process)
		}

		return nil
	}
}

func (w *Client) read() error {
	defer func() {
		w.Logger.Debugf("read thread exited")
	}()
	for {
		select {
		case <-w.rwTomb.Dying():
			return nil
		default:
			_, msg, err := w.conn.conn.ReadMessage()
			if err != nil {
				w.Logger.Errorf("read error: %s", err)
				return err
			}
			w.Logger.Debugf("read: %s", string(msg))
			w.rQueue <- msg
		}

	}
}

func (w *Client) write() error {
	defer func() {
		w.Logger.Debugf("write thread exited")
		go w.reconnect()
	}()
	for {
		select {
		case <-w.rwTomb.Dying():
			writeWait := time.Second * 5
			if err := w.conn.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait)); err != nil {
				return fmt.Errorf("failed to gracefully close: %w", err)
			}
			return nil
		case msg := <-w.wQueue:
			w.Logger.Debugf("write: %s", string(msg))
			if err := w.conn.sendJson(msg); err != nil {
				return err
			}
		}
	}
}

func (w *Client) reconnect() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.shouldClose {
		return
	}

	w.Logger.Debugf("unexpected disconnect, reconnecting")
	w.close(true)

	notify := func(err error, _ time.Duration) {
		w.Logger.Errorf(err.Error())
	}
	err := backoff.RetryNotify(w.connect(true), w.backoff, notify)
	if err != nil {
		err = fmt.Errorf("error reconnecting: %w: closing connection", err)
		w.Logger.Errorf(err.Error())
		w.close(false)
		w.err <- err
	}
}

func (w *Client) close(reconnect bool) {
	if w.conn == nil {
		return
	}

	w.rwTomb.Kill(nil)
	if err := w.rwTomb.Wait(); err != nil {
		w.Logger.Errorf("r/w threads closed: %v", err)
	}

	if !reconnect {
		w.pTomb.Kill(nil)
		if err := w.pTomb.Wait(); err != nil {
			w.Logger.Errorf("process thread closed: %v", err)
		}
		w.shouldClose = true
		w.closeOutput()
	}

	if w.conn != nil {
		_ = w.conn.conn.Close()
		w.conn = nil
	}
}

func (w *Client) closeOutput() {
	close(w.output)
	w.Logger.Debugf("output channel closed")
}

func (w *Client) process() (err error) {
	defer func() {
		// this client should close if it hits a fatal error (e.g. auth failed)
		w.Logger.Debugf("process thread closed")
		if err != nil {
			go w.Close()
			w.err <- err
		}
	}()

	for {
		select {
		case <-w.pTomb.Dying():
			return nil
		case data := <-w.rQueue:
			var msgs []json.RawMessage
			if err := json.Unmarshal(data, &msgs); err != nil {
				w.Logger.Errorf("failed to process raw messages: %v", err)
				continue
			}
			if err := w.route(msgs); err != nil {
				return err
			}
		}
	}
}

func (w *Client) Close() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.close(false)
}

func (w *Client) route(msgs []json.RawMessage) error {
	for _, msg := range msgs {
		var ev models.ChannelMessage
		err := json.Unmarshal(msg, &ev)
		if err != nil {
			w.Logger.Errorf("failed to process message: %v", err)
			continue
		}

		switch ev.Channel {
		default:
			w.Logger.Debugf("unknown channel: %s", ev.Channel)

		}
	}

	return nil
}

func (w *Client) Error() chan error {
	return w.err
}

func (w *Client) Output() <-chan any {
	return w.output
}
