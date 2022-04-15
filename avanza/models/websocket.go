package models

type ConnectMessage struct {
	ID             string          `json:"id"`
	Channel        string          `json:"channel"`
	ConnectionType string          `json:"connectionType"`
	Advice         WebsocketAdvice `json:"advice,omitempty"`
	ClientID       string          `json:"clientId"`
}

type ConnectResponse struct {
	Id         string `json:"id"`
	Channel    string `json:"channel"`
	Successful bool   `json:"successful"`
	ClientId   string `json:"clientId"`
}

type SubscribeRequest struct {
	ID           string `json:"id"`
	Channel      string `json:"channel"`
	Subscription string `json:"subscription"`
	ClientID     string `json:"clientId"`
}

type SubscribeResponse struct {
	Channel      string `json:"channel"`
	ID           string `json:"id"`
	Subscription string `json:"subscription"`
	Successful   bool   `json:"successful"`
}

type StreamQuote struct {
	Data struct {
		OrderbookID                string      `json:"orderbookId"`
		BuyPrice                   interface{} `json:"buyPrice"`
		SellPrice                  interface{} `json:"sellPrice"`
		Spread                     interface{} `json:"spread"`
		ClosingPrice               float64     `json:"closingPrice"`
		HighestPrice               float64     `json:"highestPrice"`
		LowestPrice                float64     `json:"lowestPrice"`
		LastPrice                  float64     `json:"lastPrice"`
		Change                     float64     `json:"change"`
		ChangePercent              float64     `json:"changePercent"`
		Updated                    Millis      `json:"updated"`
		VolumeWeightedAveragePrice interface{} `json:"volumeWeightedAveragePrice"`
		TotalVolumeTraded          int         `json:"totalVolumeTraded"`
		TotalValueTraded           float64     `json:"totalValueTraded"`
		LastUpdated                Millis      `json:"lastUpdated"`
		ChangePercentNumber        float64     `json:"changePercentNumber"`
		UpdatedDisplay             string      `json:"updatedDisplay"`
	} `json:"data"`
	Channel string `json:"channel"`
}

type HandshakeResponse struct {
	MinimumVersion           string          `json:"minimumVersion"`
	ClientID                 string          `json:"clientId"`
	SupportedConnectionTypes []string        `json:"supportedConnectionTypes"`
	Advice                   WebsocketAdvice `json:"advice"`
	Channel                  string          `json:"channel"`
	ID                       string          `json:"id"`
	Version                  string          `json:"version"`
	Successful               bool            `json:"successful"`
}

type WebsocketExt struct {
	SubscriptionID string `json:"subscriptionId"`
}

type WebsocketAdvice struct {
	Timeout   int    `json:"timeout,omitempty"`
	Interval  int    `json:"interval,omitempty"`
	Reconnect string `json:"reconnect,omitempty"`
}

type HandshakeMessage struct {
	Ext                      WebsocketExt    `json:"ext"`
	ID                       string          `json:"id"`
	Version                  string          `json:"version"`
	MinimumVersion           string          `json:"minimumVersion"`
	Channel                  string          `json:"channel"`
	SupportedConnectionTypes []string        `json:"supportedConnectionTypes"`
	Advice                   WebsocketAdvice `json:"advice"`
}
