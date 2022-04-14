package goavanza

/*
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

type SubscriptionEvent struct {
	Channel      string `json:"channel"`
	ID           string `json:"id"`
	Subscription string `json:"subscription"`
	Successful   bool   `json:"successful"`
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

type QuoteMessage struct {
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
		Updated                    int64       `json:"updated"`
		VolumeWeightedAveragePrice interface{} `json:"volumeWeightedAveragePrice"`
		TotalVolumeTraded          int         `json:"totalVolumeTraded"`
		TotalValueTraded           float64     `json:"totalValueTraded"`
		LastUpdated                int64       `json:"lastUpdated"`
		ChangePercentNumber        float64     `json:"changePercentNumber"`
		UpdatedDisplay             string      `json:"updatedDisplay"`
	} `json:"data"`
	Channel string `json:"channel"`
}

type OrderDepthsMessage struct {
	Data struct {
		OrderbookID  string `json:"orderbookId"`
		ReceivedTime string `json:"receivedTime"`
		TotalLevel   struct {
			BuySide struct {
				Price         float64 `json:"price"`
				Volume        float64 `json:"volume"`
				VolumePercent int     `json:"volumePercent"`
			} `json:"buySide"`
			SellSide struct {
				Price         float64 `json:"price"`
				Volume        float64 `json:"volume"`
				VolumePercent int     `json:"volumePercent"`
			} `json:"sellSide"`
		} `json:"totalLevel"`
		Levels []struct {
			BuySide struct {
				Price         float64 `json:"price"`
				Volume        int     `json:"volume"`
				VolumePercent int     `json:"volumePercent"`
			} `json:"buySide"`
			SellSide struct {
				Price         float64 `json:"price"`
				Volume        int     `json:"volume"`
				VolumePercent int     `json:"volumePercent"`
			} `json:"sellSide"`
		} `json:"levels"`
		MarketMakerLevelAsk interface{} `json:"marketMakerLevelAsk"`
		MarketMakerLevelBid interface{} `json:"marketMakerLevelBid"`
	} `json:"data"`
	Channel string `json:"channel"`
}
*/
