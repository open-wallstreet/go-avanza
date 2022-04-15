package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"go.uber.org/atomic"
	"strings"
	"time"
)

// Conn defines a connection to a WebSocket server.
type Conn struct {
	conn                *websocket.Conn
	clientID            string
	messageCount        *atomic.Int64
	pushSubscriptionID  string
	reAuthenticateTimer *time.Timer
}

func (c *Conn) Collect(data chan []byte) {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			println("failed to read message")
			return
		}
		println(string(msg))
		if strings.Contains(string(msg), PingChannel) {
			var connectionMessages []models.ConnectResponse
			err := json.Unmarshal(msg, &connectionMessages)
			if err != nil {
				println("failed to unmarshal ping message")
				continue
			}
			if len(connectionMessages) == 0 || !connectionMessages[0].Successful {
				println("unsuccessful ping message")
				return
			}
			c.ping()
		} else {
			data <- msg
		}
	}
}

func (c *Conn) authenticate() error {
	err := c.sendJson([]*models.HandshakeMessage{{
		Ext: models.WebsocketExt{
			SubscriptionID: c.pushSubscriptionID,
		},
		ID:                       c.messageCount.String(),
		Version:                  "1.0",
		MinimumVersion:           "1.0",
		Channel:                  "/meta/handshake",
		SupportedConnectionTypes: []string{"websocket", "long-polling", "callback-polling"},
		Advice: models.WebsocketAdvice{
			Timeout:  60000,
			Interval: 0,
		},
	}})
	if err != nil {
		return err
	}
	var res []models.HandshakeResponse
	err = c.conn.ReadJSON(&res)
	if err != nil {
		return err
	}
	if len(res) == 0 || !res[0].Successful {
		return fmt.Errorf("failed to handshake")
	}
	c.clientID = res[0].ClientID

	err = c.sendJson([]*models.ConnectMessage{{
		ID:             c.messageCount.String(),
		ConnectionType: "websocket",
		Channel:        "/meta/connect",
		Advice: models.WebsocketAdvice{
			Timeout: 0,
		},
		ClientID: c.clientID,
	}})
	if err != nil {
		return err
	}
	var connectionMessages []models.ConnectResponse
	err = c.conn.ReadJSON(&connectionMessages)
	if err != nil {
		return err
	}
	if len(connectionMessages) == 0 || !connectionMessages[0].Successful {
		return fmt.Errorf("unsuccessful ping message")
	}
	return err
}

func (c *Conn) Close() {
	if c.reAuthenticateTimer != nil {
		c.reAuthenticateTimer.Stop()
	}
}

func (c *Conn) ping() {
	err := c.sendJson([]*models.ConnectMessage{{
		ID:             c.messageCount.String(),
		ConnectionType: "websocket",
		Channel:        "/meta/connect",
		ClientID:       c.clientID,
	}})
	if err != nil {
		println("failed to ping")
	}
}

func (c *Conn) subscribe(params string) error {
	marshal, _ := json.Marshal(&models.SubscribeRequest{
		Subscription: params,
		Channel:      "/meta/subscribe",
		ClientID:     c.clientID,
		ID:           c.messageCount.String(),
	})
	println(string(marshal))
	msg := []*models.SubscribeRequest{{
		Subscription: params,
		Channel:      "/meta/subscribe",
		ClientID:     c.clientID,
		ID:           c.messageCount.String(),
	}}
	err := c.sendJson(msg)
	if err != nil {
		return err
	}
	var res []models.SubscribeResponse
	err = c.conn.ReadJSON(&res)
	if err != nil {
		return err
	}
	if len(res) == 0 || !res[0].Successful {
		return fmt.Errorf("failed to subscribe")
	}
	return nil
}

func (c *Conn) sendJson(req interface{}) error {
	err := c.conn.WriteJSON(&req)
	if err != nil {
		return err
	}
	c.messageCount.Inc()
	return nil
}
