package goavanza

type HandshakeEvent struct {
	MinimumVersion           string   `json:"minimumVersion"`
	ClientID                 string   `json:"clientId"`
	SupportedConnectionTypes []string `json:"supportedConnectionTypes"`
	Advice                   struct {
		Interval  int    `json:"interval"`
		Timeout   int    `json:"timeout"`
		Reconnect string `json:"reconnect"`
	} `json:"advice"`
	Channel    string `json:"channel"`
	ID         string `json:"id"`
	Version    string `json:"version"`
	Successful bool   `json:"successful"`
}

type ConnectMessageAdvice struct {
	Timeout int `json:"timeout"`
}

type ConnectMessage struct {
	Advice         ConnectMessageAdvice `json:"advice"`
	Channel        string               `json:"channel"`
	ClientId       string               `json:"clientId"`
	ConnectionType string               `json:"connectionType"`
	ID             int                  `json:"id"`
}

type ConnectEvent struct {
	Advice struct {
		Interval  int    `json:"interval"`
		Timeout   int    `json:"timeout"`
		Reconnect string `json:"reconnect"`
	} `json:"advice"`
	Channel    string `json:"channel"`
	ID         string `json:"id"`
	Successful bool   `json:"successful"`
}

type OrderSubscribeMessage struct {
	ID           int    `json:"id"`
	Channel      string `json:"channel"`
	Subscription string `json:"subscription"`
	ClientID     string `json:"clientId"`
}
