package goavanza

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sacOO7/gowebsocket"
	"go.uber.org/zap"
)

const SocketUrl string = "wss://www.avanza.se/_push/cometd"

type AvanzaWebsocket struct {
	client             *Client
	logger             *zap.SugaredLogger
	clientId           string
	socketConnected    bool
	socketMessageCount int
	socket             *gowebsocket.Socket
	options            *AvanzaWebsocketOptions
	subscriptions      map[string]string
}

type AvanzaWebsocketOptions struct {
	OnError      func(error)
	OnConnected  func()
	OnDisconnect func(error)
	OnQuote      func(QuoteMessage)
}

type SocketMessage struct {
	Advice *struct {
		Timeout  int `json:"timeout"`
		Interval int `json:"interval"`
	} `json:"advice"`
	Channel string `json:"channel"`
	Ext     *struct {
		SubscriptionId *string `json:"subscriptionId"`
	} `json:"ext"`
	ID                       *int      `json:"id"`
	MinimumVersion           *string   `json:"minimumVersion"`
	SupportedConnectionTypes []*string `json:"supportedConnectionTypes"`
	Version                  *string   `json:"version"`
	ClientId                 *string   `json:"clientId"`
	ConnectionType           *string   `json:"connectionType"`
}

func StringPtr(s string) *string {
	return &s
}

func (ws *AvanzaWebsocket) Listen() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New(SocketUrl)

	ws.socket = &socket

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		if ws.options.OnError != nil {
			ws.options.OnError(err)
		}
	}
	socket.OnConnected = func(socket gowebsocket.Socket) {
		bin, _ := json.Marshal([]SocketMessage{{
			Advice: &struct {
				Timeout  int "json:\"timeout\""
				Interval int "json:\"interval\""
			}{
				Timeout:  60000,
				Interval: 0,
			},
			Channel: "/meta/handshake",
			Ext: &struct {
				SubscriptionId *string "json:\"subscriptionId\""
			}{
				SubscriptionId: StringPtr(ws.client.PushSubscriptionId()),
			},
			ID:                       &ws.socketMessageCount,
			MinimumVersion:           StringPtr("1.0"),
			SupportedConnectionTypes: []*string{StringPtr("websocket"), StringPtr("long-polling"), StringPtr("callback-polling")},
			Version:                  StringPtr("1.0"),
		}})
		ws.sendSocketMessage(bin)
		if ws.options.OnConnected != nil {
			ws.options.OnConnected()
		}
	}
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		var anyJson []map[string]interface{}
		b := []byte(message)
		err := json.Unmarshal(b, &anyJson)
		if err != nil {
			ws.logger.Error(err)
		}
		for _, message := range anyJson {
			channel := message["channel"].(string)
			println(channel)
			switch channel {
			case "/meta/handshake":
				ws.handleHandshakeMessage(message)
			case "/meta/connect":
				ws.handleConnectMessage(message)
			case "/meta/subscribe":
				ws.onSubscribeMessage(message)
			case "/meta/unsubscribe":
				ws.onUnsubscribeMessage(message)
			default:
				switch {
				case strings.HasPrefix(channel, "/quotes/"):
					ws.onQuotesMessage(message)
				default:
					ws.logger.Warn("got unhandled channel message %s", message["channel"].(string))
				}
			}

		}

	}
	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		if ws.options.OnDisconnect != nil {
			ws.options.OnDisconnect(err)
		}
	}
	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return nil
		}
	}
}

func (ws *AvanzaWebsocket) onUnsubscribeMessage(message map[string]interface{}) {
	var subscription SubscriptionEvent
	err := mapstructure.Decode(message, &subscription)
	if err != nil {
		ws.logger.Errorf("failed to decode: %v", err)
	}
	if subscription.Successful {
		delete(ws.subscriptions, subscription.Subscription)
	}
}

func (ws *AvanzaWebsocket) onQuotesMessage(message map[string]interface{}) {
	var quote QuoteMessage
	err := mapstructure.Decode(message, &quote)
	if err != nil {
		ws.logger.Errorf("failed to decode: %v", err)
	}
	if ws.options.OnQuote != nil {
		ws.options.OnQuote(quote)
	}
}
func (ws *AvanzaWebsocket) onSubscribeMessage(message map[string]interface{}) {
	var subscription SubscriptionEvent
	err := mapstructure.Decode(message, &subscription)
	if err != nil {
		ws.logger.Errorf("failed to decode: %v", err)
	}
	if subscription.Successful {
		ws.subscriptions[subscription.Subscription] = ws.clientId
	} else {
		ws.logger.Errorf("failed to subscribe %v", subscription)
	}
}

func (ws *AvanzaWebsocket) sendSocketMessage(data []byte) error {
	if ws.socket == nil {
		return fmt.Errorf("no websocket connection has been established")
	}
	ws.socket.SendText(string(data))
	ws.socketMessageCount += 1
	return nil
}

func (ws *AvanzaWebsocket) handleHandshakeMessage(message map[string]interface{}) {
	var handshake HandshakeEvent
	err := mapstructure.Decode(message, &handshake)
	if err != nil {
		ws.logger.Errorf("failed to decode: %v", err)
	}
	ws.clientId = handshake.ClientID
	if handshake.Successful {
		bin, _ := json.Marshal([]ConnectMessage{{
			Advice: ConnectMessageAdvice{
				Timeout: 0,
			},
			Channel:        "/meta/connect",
			ClientId:       ws.clientId,
			ConnectionType: "websocket",
			ID:             ws.socketMessageCount,
		}})
		ws.sendSocketMessage(bin)
	} else {
		ws.logger.Errorf("failed to connect to websocket %v", err)
	}
}
func (ws *AvanzaWebsocket) handleConnectMessage(message map[string]interface{}) {
	var msg ConnectEvent
	// ws.logger.Info(message)
	err := mapstructure.Decode(message, &msg)
	if err != nil {
		ws.logger.Errorf("failed to decode: %v", err)
	}
	if msg.Successful {
		subscriptionIds := make([]string, 0, len(ws.subscriptions))
		for key, v := range ws.subscriptions {
			ws.logger.Infof("sub to %s", key)
			if v != ws.clientId {
				subscriptionIds = append(subscriptionIds, key)
			}
		}
		ws.Subscribe(subscriptionIds)
	}
}

func (ws *AvanzaWebsocket) Subscribe(ids []string) error {
	if len(ws.client.PushSubscriptionId()) == 0 {
		return fmt.Errorf("need to be authenticated to subscribe to socket")
	}
	idString := strings.Join(ids, ",")

	bin, _ := json.Marshal([]OrderSubscribeMessage{{
		Subscription: idString,
		Channel:      "/meta/subscribe",
		ClientID:     ws.clientId,
		ID:           ws.socketMessageCount,
	}})
	for _, id := range ids {
		ws.subscriptions[id] = ""
	}
	ws.sendSocketMessage(bin)
	return nil
}

func (ws *AvanzaWebsocket) Unsubscribe(ids []string) error {
	if len(ws.client.PushSubscriptionId()) == 0 {
		return fmt.Errorf("need to be authenticated to subscribe to socket")
	}
	idString := strings.Join(ids, ",")

	bin, _ := json.Marshal([]OrderSubscribeMessage{{
		Subscription: idString,
		Channel:      "/meta/unsubscribe",
		ClientID:     ws.clientId,
		ID:           ws.socketMessageCount,
	}})
	ws.sendSocketMessage(bin)
	return nil
}

func NewAvanzaWebsocketOptions() *AvanzaWebsocketOptions {
	return &AvanzaWebsocketOptions{
		OnError:      func(e error) {},
		OnConnected:  func() {},
		OnDisconnect: func(e error) {},
		OnQuote:      func(q QuoteMessage) {},
	}
}

func NewWebsocket(client *Client, logger *zap.SugaredLogger, options *AvanzaWebsocketOptions) *AvanzaWebsocket {
	return &AvanzaWebsocket{
		socketMessageCount: 1,
		client:             client,
		logger:             logger,
		options:            options,
		subscriptions:      make(map[string]string),
	}
}
