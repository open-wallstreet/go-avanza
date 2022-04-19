package models

type GetMarketDataParams struct {
	OrderBookID string `path:"orderbookID" validate:"required"`
}

type GetMarketDataResponse struct {
	Quote struct {
		Buy               float64 `json:"buy"`
		Sell              float64 `json:"sell"`
		Last              float64 `json:"last"`
		Highest           float64 `json:"highest"`
		Lowest            float64 `json:"lowest"`
		Change            float64 `json:"change"`
		ChangePercent     float64 `json:"changePercent"`
		TimeOfLast        string  `json:"timeOfLast"`
		TotalValueTraded  float64 `json:"totalValueTraded"`
		TotalVolumeTraded int     `json:"totalVolumeTraded"`
		Updated           string  `json:"updated"`
	} `json:"quote"`
	OrderDepth struct {
		ReceivedTime Millis `json:"receivedTime"`
		Levels       []struct {
			BuySide struct {
				Price  float64 `json:"price"`
				Volume int     `json:"volume"`
			} `json:"buySide"`
			SellSide struct {
				Price  float64 `json:"price"`
				Volume int     `json:"volume"`
			} `json:"sellSide"`
		} `json:"levels"`
		MarketMakerExpected bool `json:"marketMakerExpected"`
	} `json:"orderDepth"`
	Trades []struct {
		Buyer           string  `json:"buyer"`
		Seller          string  `json:"seller"`
		DealTime        int64   `json:"dealTime"`
		Price           float64 `json:"price"`
		Volume          int     `json:"volume"`
		MatchedOnMarket bool    `json:"matchedOnMarket"`
		Cancelled       bool    `json:"cancelled"`
	} `json:"trades"`
}

type GetInstrumentParams struct {
	Instrument InstrumentType `path:"instrument" validate:"required"`
	ID         string         `path:"id" validate:"required"`
}

type GetInstrumentResponse struct {
	PriceThreeMonthsAgo     float64 `json:"priceThreeMonthsAgo"`
	PriceOneWeekAgo         float64 `json:"priceOneWeekAgo"`
	PriceOneMonthAgo        float64 `json:"priceOneMonthAgo"`
	PriceSixMonthsAgo       float64 `json:"priceSixMonthsAgo"`
	PriceAtStartOfYear      float64 `json:"priceAtStartOfYear"`
	PriceOneYearAgo         float64 `json:"priceOneYearAgo"`
	PriceThreeYearsAgo      float64 `json:"priceThreeYearsAgo"`
	PriceFiveYearsAgo       float64 `json:"priceFiveYearsAgo"`
	MarketPlace             string  `json:"marketPlace"`
	MarketList              string  `json:"marketList"`
	QuoteUpdated            string  `json:"quoteUpdated"`
	HasInvestmentFees       bool    `json:"hasInvestmentFees"`
	MorningStarFactSheetURL string  `json:"morningStarFactSheetUrl"`
	Currency                string  `json:"currency"`
	BuyPrice                float64 `json:"buyPrice"`
	LowestPrice             float64 `json:"lowestPrice"`
	HighestPrice            float64 `json:"highestPrice"`
	TotalVolumeTraded       float64 `json:"totalVolumeTraded"`
	SellPrice               float64 `json:"sellPrice"`
	Isin                    string  `json:"isin"`
	LastPrice               float64 `json:"lastPrice"`
	LastPriceUpdated        string  `json:"lastPriceUpdated"`
	Change                  float64 `json:"change"`
	ChangePercent           float64 `json:"changePercent"`
	TotalValueTraded        float64 `json:"totalValueTraded"`
	ShortSellable           bool    `json:"shortSellable"`
	Tradable                bool    `json:"tradable"`
	TickerSymbol            string  `json:"tickerSymbol"`
	FlagCode                string  `json:"flagCode"`
	LoanFactor              float64 `json:"loanFactor"`
	Name                    string  `json:"name"`
	ID                      string  `json:"id"`
	Country                 string  `json:"country"`
	KeyRatios               struct {
		PriceEarningsRatio float64 `json:"priceEarningsRatio"`
		Volatility         float64 `json:"volatility"`
		DirectYield        float64 `json:"directYield"`
	} `json:"keyRatios"`
	NumberOfOwners      int  `json:"numberOfOwners"`
	SuperLoan           bool `json:"superLoan"`
	NumberOfPriceAlerts int  `json:"numberOfPriceAlerts"`
	PushPermitted       bool `json:"pushPermitted"`
	Dividends           []struct {
		ExDate         string  `json:"exDate"`
		AmountPerShare float64 `json:"amountPerShare"`
		PaymentDate    string  `json:"paymentDate"`
		Currency       string  `json:"currency"`
	} `json:"dividends"`
	RelatedStocks []struct {
		PriceOneYearAgo float64 `json:"priceOneYearAgo,omitempty"`
		LastPrice       float64 `json:"lastPrice"`
		FlagCode        string  `json:"flagCode"`
		Name            string  `json:"name"`
		ID              string  `json:"id"`
	} `json:"relatedStocks"`
	Company struct {
		Sector string `json:"sector"`
		Stocks []struct {
			TotalNumberOfShares int    `json:"totalNumberOfShares"`
			Name                string `json:"name"`
		} `json:"stocks"`
		Chairman              string `json:"chairman"`
		TotalNumberOfShares   int    `json:"totalNumberOfShares"`
		Description           string `json:"description"`
		MarketCapital         int    `json:"marketCapital"`
		MarketCapitalCurrency string `json:"marketCapitalCurrency"`
		Name                  string `json:"name"`
		ID                    string `json:"id"`
		Ceo                   string `json:"CEO"`
	} `json:"company"`
	OrderDepthLevels []struct {
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
		Buyer           string  `json:"buyer"`
		Seller          string  `json:"seller"`
		MatchedOnMarket bool    `json:"matchedOnMarket"`
		Cancelled       bool    `json:"cancelled"`
		Price           float64 `json:"price"`
		Volume          int     `json:"volume"`
		DealTime        string  `json:"dealTime"`
	} `json:"latestTrades"`
	MarketTrades bool `json:"marketTrades"`
	Positions    []struct {
		AccountName          string  `json:"accountName"`
		AccountType          string  `json:"accountType"`
		AccountID            string  `json:"accountId"`
		Volume               int     `json:"volume"`
		AverageAcquiredPrice float64 `json:"averageAcquiredPrice"`
		ProfitPercent        float64 `json:"profitPercent"`
		AcquiredValue        float64 `json:"acquiredValue"`
		Profit               float64 `json:"profit"`
		Value                float64 `json:"value"`
	} `json:"positions"`
	PositionsTotalValue float64 `json:"positionsTotalValue"`
	AnnualMeetings      []struct {
		EventDate string `json:"eventDate"`
		Extra     bool   `json:"extra"`
	} `json:"annualMeetings"`
	CompanyReports []struct {
		EventDate  string `json:"eventDate"`
		ReportType string `json:"reportType"`
	} `json:"companyReports"`
	BrokerTradeSummary struct {
		OrderbookID string `json:"orderbookId"`
		Items       []struct {
			NetBuyVolume int    `json:"netBuyVolume"`
			BuyVolume    int    `json:"buyVolume"`
			SellVolume   int    `json:"sellVolume"`
			BrokerCode   string `json:"brokerCode"`
		} `json:"items"`
	} `json:"brokerTradeSummary"`
	CompanyOwners struct {
		List []struct {
			Name    string  `json:"name"`
			Capital float64 `json:"capital"`
			Votes   float64 `json:"votes"`
		} `json:"list"`
		Updated string `json:"updated"`
	} `json:"companyOwners"`
}

type SearchParams struct {
	Instrument InstrumentType `path:"instrument"`
	Query      string         `query:"query" validate:"required"`
	Limit      *int           `query:"limit"`
}

type SearchResponse struct {
	TotalNumberOfHits int `json:"totalNumberOfHits"`
	Hits              []struct {
		InstrumentType InstrumentType `json:"instrumentType"`
		NumberOfHits   int            `json:"numberOfHits"`
		TopHits        []struct {
			Currency      string  `json:"currency"`
			LastPrice     float64 `json:"lastPrice"`
			ChangePercent float64 `json:"changePercent"`
			FlagCode      string  `json:"flagCode"`
			Tradable      bool    `json:"tradable"`
			TickerSymbol  string  `json:"tickerSymbol"`
			Name          string  `json:"name"`
			ID            string  `json:"id"`
		} `json:"topHits,omitempty"`
	} `json:"hits"`
}
