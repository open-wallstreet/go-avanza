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
)

const SocketUrl string = "wss://www.avanza.se/_push/cometd"

type websocketConnection struct {
	clientId           string
	socketConnected    bool
	pushSubscriptionId string
	socketMessageCount int
	socket             *gowebsocket.Socket
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

func (a *api) Listen() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New(SocketUrl)
	/*socket.ConnectionOptions = gowebsocket.ConnectionOptions{
		//Proxy: gowebsocket.BuildProxy("http://example.com"),
		UseSSL:         false,
		UseCompression: false,
		Subprotocols:   []string{"chat", "superchat"},
	}*/
	a.websocketConnection.socket = &socket

	/*socket.RequestHeader.Set("Accept-Encoding", "gzip, deflate, sdch")
	socket.RequestHeader.Set("Accept-Language", "en-US,en;q=0.8")
	socket.RequestHeader.Set("Pragma", "no-cache")
	socket.RequestHeader.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")
	*/
	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Recieved connect error ", err)
	}
	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
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
				SubscriptionId: &a.websocketConnection.pushSubscriptionId,
			},
			ID:                       &a.websocketConnection.socketMessageCount,
			MinimumVersion:           StringPtr("1.0"),
			SupportedConnectionTypes: []*string{StringPtr("websocket"), StringPtr("long-polling"), StringPtr("callback-polling")},
			Version:                  StringPtr("1.0"),
		}})
		a.sendSocketMessage(bin)
	}
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message  " + message)
		var anyJson []map[string]interface{}
		b := []byte(message)
		err := json.Unmarshal(b, &anyJson)
		if err != nil {
			a.logger.Error(err)
		}
		for _, message := range anyJson {
			channel := message["channel"].(string)
			a.logger.Infof("channel: %s", message["channel"].(string))
			switch channel {
			case "/meta/handshake":
				a.handleHandshakeMessage(message)
				break
			case "/meta/connect":
				a.handleConnectMessage(message)
				a.SubscribeOrders([]string{"*"})
				break
			}

		}

	}
	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)
	}
	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		a.logger.Infof("Disconnected from server %v", err)

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

func (a *api) sendSocketMessage(data []byte) error {
	a.logger.Infof(fmt.Sprintf("sending %s", data))
	a.websocketConnection.socket.SendText(string(data))
	a.websocketConnection.socketMessageCount += 1
	return nil
}

func (a *api) handleHandshakeMessage(message map[string]interface{}) {
	var handshake HandshakeEvent
	err := mapstructure.Decode(message, &handshake)
	if err != nil {
		a.logger.Errorf("failed to decode: %v", err)
	}
	a.websocketConnection.clientId = handshake.ClientID
	if handshake.Successful {
		bin, _ := json.Marshal([]ConnectMessage{{
			Advice: ConnectMessageAdvice{
				Timeout: 0,
			},
			Channel:        "/meta/connect",
			ClientId:       a.websocketConnection.clientId,
			ConnectionType: "websocket",
			ID:             a.websocketConnection.socketMessageCount,
		}})
		a.sendSocketMessage(bin)
	} else {
		a.logger.Errorf("failed to connect to websocket %v", err)
	}
}
func (a *api) handleConnectMessage(message map[string]interface{}) {
	var msg ConnectEvent
	a.logger.Info(message)
	err := mapstructure.Decode(message, &msg)
	if err != nil {
		a.logger.Errorf("failed to decode: %v", err)
	}
	a.logger.Info(msg)
}

func (a *api) SubscribeOrders(ids []string) error {
	if len(a.websocketConnection.pushSubscriptionId) == 0 {
		return fmt.Errorf("need to be authenticated to subscribe to socket")
	}
	idString := strings.Join(ids, ",")

	subscribeString := fmt.Sprintf("/orders/%s", idString)

	bin, _ := json.Marshal([]OrderSubscribeMessage{{
		Subscription: subscribeString,
		Channel:      "/meta/subscribe",
		ClientID:     a.websocketConnection.clientId,
		ID:           a.websocketConnection.socketMessageCount,
	}})
	a.sendSocketMessage(bin)
	return nil
}
