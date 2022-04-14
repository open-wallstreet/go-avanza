package models

type DeleteOrderParams struct {
	AccountID string `query:"accountId" validate:"required"`
	OrderID   string `query:"orderId" validate:"required"`
}

type DeleteOrderResponse struct {
	Status    string   `json:"status"`
	Messages  []string `json:"messages"`
	OrderID   string   `json:"orderId"`
	RequestID string   `json:"requestId"`
}

type EditOrderParams struct {
	OrderID        string         `path:"id" validate:"required"`
	InstrumentType InstrumentType `path:"instrument" validate:"required"`
	AccountID      string
	OrderBookID    string
	OrderType      OrderType
	Price          float64
	Volume         int
	ValidUntil     Date
}

type EditOrderResponse struct {
	Status    string   `json:"status"`
	Messages  []string `json:"messages"`
	OrderID   string   `json:"orderId"`
	RequestID string   `json:"requestId"`
}

type PlaceOrderParams struct {
	AccountID   string    `validate:"required"`
	OrderBookID string    `validate:"required"`
	OrderType   OrderType `validate:"required"`
	Price       float64   `validate:"required"`
	Volume      int       `validate:"required"`
	ValidUntil  Date      `validate:"required"`
}

type PlaceOrderResponse struct {
	Status    string   `json:"status"`
	Messages  []string `json:"messages"`
	OrderID   string   `json:"orderId"`
	RequestID string   `json:"requestId"`
}

type GetOrderBooksParams struct {
	OrderBookIDs []string `path:"orderbookIds" validate:"required"`
}

type GetOrderBooksResponse []struct {
	ChangePercentPeriod      float64 `json:"changePercentPeriod"`
	ChangePercentThreeMonths float64 `json:"changePercentThreeMonths"`
	MinMonthlySavingAmount   float64 `json:"minMonthlySavingAmount"`
	ChangePercentOneYear     float64 `json:"changePercentOneYear"`
	Currency                 string  `json:"currency"`
	ManagementFee            float64 `json:"managementFee"`
	Prospectus               string  `json:"prospectus"`
	Rating                   float64 `json:"rating"`
	Risk                     float64 `json:"risk"`
	LastUpdated              string  `json:"lastUpdated"`
	Tradable                 bool    `json:"tradable"`
	InstrumentType           string  `json:"instrumentType"`
	Name                     string  `json:"name"`
	ID                       string  `json:"id"`
}

type GetOrderBookParams struct {
	OrderBookID string         `query:"orderbookId" validate:"required"`
	Instrument  InstrumentType `path:"instrument" validate:"required"`
}

type GetOrderBookResponse struct {
	Customer struct {
		ShowCourtageClassInfoOnOrderPage bool   `json:"showCourtageClassInfoOnOrderPage"`
		CourtageClass                    string `json:"courtageClass"`
	} `json:"customer"`
	Account struct {
		Type         string  `json:"type"`
		TotalBalance float64 `json:"totalBalance"`
		BuyingPower  float64 `json:"buyingPower"`
		Name         string  `json:"name"`
		ID           string  `json:"id"`
	} `json:"account"`
	Orderbook struct {
		BuyPrice          float64 `json:"buyPrice"`
		SellPrice         float64 `json:"sellPrice"`
		Spread            float64 `json:"spread"`
		HighestPrice      float64 `json:"highestPrice"`
		LowestPrice       float64 `json:"lowestPrice"`
		LastPrice         float64 `json:"lastPrice"`
		LastPriceUpdated  string  `json:"lastPriceUpdated"`
		Change            float64 `json:"change"`
		ChangePercent     float64 `json:"changePercent"`
		TotalVolumeTraded int     `json:"totalVolumeTraded"`
		TotalValueTraded  float64 `json:"totalValueTraded"`
		ExchangeRate      float64 `json:"exchangeRate"`
		PositionVolume    int     `json:"positionVolume"`
		Currency          string  `json:"currency"`
		Tradable          bool    `json:"tradable"`
		TradingUnit       int     `json:"tradingUnit"`
		TickerSymbol      string  `json:"tickerSymbol"`
		FlagCode          string  `json:"flagCode"`
		VolumeFactor      float64 `json:"volumeFactor"`
		Name              string  `json:"name"`
		ID                string  `json:"id"`
		Type              string  `json:"type"`
	} `json:"orderbook"`
	FirstTradableDate string   `json:"firstTradableDate"`
	LastTradableDate  string   `json:"lastTradableDate"`
	UntradableDates   []string `json:"untradableDates"`
	OrderDepthLevels  []struct {
		Buy struct {
			Percent float64 `json:"percent"`
			Price   float64 `json:"price"`
			Volume  int     `json:"volume"`
		} `json:"buy"`
		Sell struct {
			Percent float64 `json:"percent"`
			Price   float64 `json:"price"`
			Volume  int     `json:"volume"`
		} `json:"sell"`
	} `json:"orderDepthLevels"`
	MarketMakerExpected    bool   `json:"marketMakerExpected"`
	OrderDepthReceivedTime string `json:"orderDepthReceivedTime"`
	LatestTrades           []struct {
		Cancelled       bool    `json:"cancelled"`
		Buyer           string  `json:"buyer"`
		Seller          string  `json:"seller"`
		MatchedOnMarket bool    `json:"matchedOnMarket"`
		Price           float64 `json:"price"`
		Volume          int     `json:"volume"`
		DealTime        string  `json:"dealTime"`
	} `json:"latestTrades"`
	MarketTrades           bool `json:"marketTrades"`
	HasShortSellKnowledge  bool `json:"hasShortSellKnowledge"`
	HasInstrumentKnowledge bool `json:"hasInstrumentKnowledge"`
	BrokerTradeSummary     struct {
		OrderbookID string `json:"orderbookId"`
		Items       []struct {
			NetBuyVolume int    `json:"netBuyVolume"`
			BuyVolume    int    `json:"buyVolume"`
			SellVolume   int    `json:"sellVolume"`
			BrokerCode   string `json:"brokerCode"`
		} `json:"items"`
	} `json:"brokerTradeSummary"`
	HasInvestmentFees struct {
		Buy  bool `json:"buy"`
		Sell bool `json:"sell"`
	} `json:"hasInvestmentFees"`
	TickSizeRules []struct {
		MinPrice float64 `json:"minPrice"`
		MaxPrice float64 `json:"maxPrice"`
		TickSize float64 `json:"tickSize"`
	} `json:"tickSizeRules"`
}
