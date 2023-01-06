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
	OrderbookID string `json:"orderbookId"`
	Name        string `json:"name"`
	Isin        string `json:"isin"`
	Sectors     []struct {
		SectorID   string `json:"sectorId"`
		SectorName string `json:"sectorName"`
	} `json:"sectors"`
	Tradable string `json:"tradable"`
	Listing  struct {
		ShortName             string `json:"shortName"`
		TickerSymbol          string `json:"tickerSymbol"`
		CountryCode           string `json:"countryCode"`
		Currency              string `json:"currency"`
		MarketPlaceCode       string `json:"marketPlaceCode"`
		MarketPlaceName       string `json:"marketPlaceName"`
		MarketListName        string `json:"marketListName"`
		TickSizeListID        string `json:"tickSizeListId"`
		MarketTradesAvailable bool   `json:"marketTradesAvailable"`
	} `json:"listing"`
	HistoricalClosingPrices struct {
		OneDay      float64 `json:"oneDay"`
		OneWeek     float64 `json:"oneWeek"`
		OneMonth    float64 `json:"oneMonth"`
		ThreeMonths float64 `json:"threeMonths"`
		StartOfYear float64 `json:"startOfYear"`
		OneYear     float64 `json:"oneYear"`
		ThreeYears  float64 `json:"threeYears"`
		FiveYears   float64 `json:"fiveYears"`
		TenYears    float64 `json:"tenYears"`
		Start       float64 `json:"start"`
		StartDate   string  `json:"startDate"`
	} `json:"historicalClosingPrices"`
	KeyIndicators struct {
		NumberOfOwners        int     `json:"numberOfOwners"`
		ReportDate            string  `json:"reportDate"`
		DirectYield           float64 `json:"directYield"`
		Volatility            float64 `json:"volatility"`
		Beta                  float64 `json:"beta"`
		PriceEarningsRatio    float64 `json:"priceEarningsRatio"`
		PriceSalesRatio       float64 `json:"priceSalesRatio"`
		InterestCoverageRatio float64 `json:"interestCoverageRatio"`
		ReturnOnEquity        float64 `json:"returnOnEquity"`
		ReturnOnTotalAssets   float64 `json:"returnOnTotalAssets"`
		EquityRatio           float64 `json:"equityRatio"`
		CapitalTurnover       float64 `json:"capitalTurnover"`
		OperatingProfitMargin float64 `json:"operatingProfitMargin"`
		GrossMargin           float64 `json:"grossMargin"`
		NetMargin             float64 `json:"netMargin"`
		MarketCapital         struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"marketCapital"`
		EquityPerShare struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"equityPerShare"`
		TurnoverPerShare struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"turnoverPerShare"`
		EarningsPerShare struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"earningsPerShare"`
		Dividend struct {
			ExDate       string  `json:"exDate"`
			PaymentDate  string  `json:"paymentDate"`
			Amount       float64 `json:"amount"`
			CurrencyCode string  `json:"currencyCode"`
			ExDateStatus string  `json:"exDateStatus"`
		} `json:"dividend"`
		DividendsPerYear int `json:"dividendsPerYear"`
		NextReport       struct {
			Date       string `json:"date"`
			ReportType string `json:"reportType"`
		} `json:"nextReport"`
		PreviousReport struct {
			Date       string `json:"date"`
			ReportType string `json:"reportType"`
		} `json:"previousReport"`
	} `json:"keyIndicators"`
	Quote struct {
		Last                       float64 `json:"last"`
		Highest                    float64 `json:"highest"`
		Lowest                     float64 `json:"lowest"`
		Change                     float64 `json:"change"`
		ChangePercent              float64 `json:"changePercent"`
		TimeOfLast                 int64   `json:"timeOfLast"`
		TotalValueTraded           float64 `json:"totalValueTraded"`
		TotalVolumeTraded          float64 `json:"totalVolumeTraded"`
		Updated                    int64   `json:"updated"`
		VolumeWeightedAveragePrice float64 `json:"volumeWeightedAveragePrice"`
	} `json:"quote"`
	Type string `json:"type"`
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
