package goavanza

type Authentication struct {
	TwoFactorLogin struct {
		TransactionID string `json:"transactionId"`
		Method        string `json:"method"`
	} `json:"twoFactorLogin"`
}

type ErrorMessage struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Time       string        `json:"time"`
	Errors     []interface{} `json:"errors"`
	Additional struct {
	} `json:"additional"`
}

type OrderActionResponse struct {
	Status    string   `json:"status"`
	Messages  []string `json:"messages"`
	OrderID   string   `json:"orderId"`
	RequestID string   `json:"requestId"`
}

type TOTPAuthentication struct {
	AuthenticationSession string `json:"authenticationSession"`
	PushSubscriptionId    string `json:"pushSubscriptionId"`
	CustomerId            string `json:"customerId"`
	RegistrationComplete  bool   `json:"registrationComplete"`
}

type Positions struct {
	InstrumentPositions []struct {
		InstrumentType string `json:"instrumentType"`
		Positions      []struct {
			AccountName          string  `json:"accountName"`
			AccountType          string  `json:"accountType"`
			Depositable          bool    `json:"depositable"`
			AccountÏd            string  `json:"accountId"`
			Profit               float64 `json:"profit"`
			Volume               float64 `json:"volume"`
			ProfitPercent        float64 `json:"profitPercent"`
			AcquiredValue        float64 `json:"acquiredValue"`
			AverageAcquiredPrice float64 `json:"averageAcquiredPrice"`
			Value                float64 `json:"value"`
			FlagCode             string  `json:"flagCode"`
			Currency             string  `json:"currency"`
			OrderbookId          string  `json:"orderbookId"`
			LastPrice            float64 `json:"lastPrice"`
			LastPriceUpdated     string  `json:"lastPriceUpdated"`
			Change               float64 `json:"change"`
			ChangePercent        float64 `json:"changePercent"`
			Tradable             bool    `json:"tradable"`
			Name                 string  `json:"name"`
		} `json:"positions"`
		TodaysProfitPercent float64 `json:"todaysProfitPercent"`
		TotalValue          float64 `json:"totalValue"`
		TotalProfitValue    float64 `json:"totalProfitValue"`
		TotalProfitPercent  float64 `json:"totalProfitPercent"`
	} `json:"instrumentPositions"`
	TotalProfit        float64 `json:"totalProfit"`
	TotalOwnCapital    float64 `json:"totalOwnCapital"`
	TotalBuyingPower   float64 `json:"totalBuyingPower"`
	TotalBalance       float64 `json:"totalBalance"`
	TotalProfitPercent float64 `json:"totalProfitPercent"`
}

type Overview struct {
	Accounts []struct {
		AccountType        string  `json:"accountType"`
		InterestRate       float64 `json:"interestRate"`
		Depositable        bool    `json:"depositable"`
		Attorney           bool    `json:"attorney"`
		Active             bool    `json:"active"`
		AccountId          string  `json:"accountId"`
		AccountPartlyOwned bool    `json:"accountPartlyOwned"`
		Tradable           bool    `json:"tradable"`
		TotalBalance       float64 `json:"totalBalance"`
		TotalBalanceDue    float64 `json:"totalBalanceDue"`
		OwnCapital         float64 `json:"ownCapital"`
		BuyingPower        float64 `json:"buyingPower"`
		TotalProfitPercent float64 `json:"totalProfitPercent"`
		Performance        float64 `json:"performance"`
		TotalProfit        float64 `json:"totalProfit"`
		PerformancePercent float64 `json:"performancePercent"`
		Name               string  `json:"name"`
		SparkontoPlusType  string  `json:"sparkontoPlusType,omitempty"`
	} `json:"accounts"`
	NumberOfOrders            int     `json:"numberOfOrders"`
	NumberOfDeals             int     `json:"numberOfDeals"`
	TotalBuyingPower          float64 `json:"totalBuyingPower"`
	TotalOwnCapital           float64 `json:"totalOwnCapital"`
	TotalPerformancePercent   float64 `json:"totalPerformancePercent"`
	TotalPerformance          float64 `json:"totalPerformance"`
	TotalBalance              float64 `json:"totalBalance"`
	NumberOfTransfers         int     `json:"numberOfTransfers"`
	NumberOfIntradayTransfers int     `json:"numberOfIntradayTransfers"`
}

type AccountOverview struct {
	CourtageClass                      string  `json:"courtageClass"`
	Depositable                        bool    `json:"depositable"`
	AccountType                        string  `json:"accountType"`
	Withdrawable                       bool    `json:"withdrawable"`
	ClearingNumber                     string  `json:"clearingNumber"`
	InstrumentTransferPossible         bool    `json:"instrumentTransferPossible"`
	InternalTransferPossible           bool    `json:"internalTransferPossible"`
	JointlyOwned                       bool    `json:"jointlyOwned"`
	AccountId                          string  `json:"accountId"`
	AccountTypeName                    string  `json:"accountTypeName"`
	InterestRate                       float64 `json:"interestRate"`
	NumberOfOrders                     int     `json:"numberOfOrders"`
	NumberOfDeals                      int     `json:"numberOfDeals"`
	PerformanceSinceOneWeek            float64 `json:"performanceSinceOneWeek"`
	PerformanceSinceOneMonth           float64 `json:"performanceSinceOneMonth"`
	PerformanceSinceThreeMonths        float64 `json:"performanceSinceThreeMonths"`
	PerformanceSinceSixMonths          float64 `json:"performanceSinceSixMonths"`
	PerformanceSinceOneYear            float64 `json:"performanceSinceOneYear"`
	PerformanceSinceThreeYears         float64 `json:"performanceSinceThreeYears"`
	PerformanceSinceOneWeekPercent     float64 `json:"performanceSinceOneWeekPercent"`
	PerformanceSinceOneMonthPercent    float64 `json:"performanceSinceOneMonthPercent"`
	PerformanceSinceThreeMonthsPercent float64 `json:"performanceSinceThreeMonthsPercent"`
	PerformanceSinceSixMonthsPercent   float64 `json:"performanceSinceSixMonthsPercent"`
	PerformanceSinceOneYearPercent     float64 `json:"performanceSinceOneYearPercent"`
	PerformanceSinceThreeYearsPercent  float64 `json:"performanceSinceThreeYearsPercent"`
	AvailableSuperLoanAmount           float64 `json:"availableSuperLoanAmount"`
	AllowMonthlySaving                 bool    `json:"allowMonthlySaving"`
	TotalProfit                        float64 `json:"totalProfit"`
	CurrencyAccounts                   []struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	} `json:"currencyAccounts"`
	CreditLimit               float64 `json:"creditLimit"`
	ForwardBalance            float64 `json:"forwardBalance"`
	ReservedAmount            float64 `json:"reservedAmount"`
	TotalCollateralValue      float64 `json:"totalCollateralValue"`
	TotalPositionsValue       float64 `json:"totalPositionsValue"`
	BuyingPower               float64 `json:"buyingPower"`
	TotalProfitPercent        float64 `json:"totalProfitPercent"`
	Overdrawn                 bool    `json:"overdrawn"`
	Performance               float64 `json:"performance"`
	AccruedInterest           float64 `json:"accruedInterest"`
	CreditAfterInterest       float64 `json:"creditAfterInterest"`
	PerformancePercent        float64 `json:"performancePercent"`
	OverMortgaged             bool    `json:"overMortgaged"`
	TotalBalance              float64 `json:"totalBalance"`
	OwnCapital                float64 `json:"ownCapital"`
	NumberOfTransfers         int     `json:"numberOfTransfers"`
	NumberOfIntradayTransfers int     `json:"numberOfIntradayTransfers"`
	StandardDeviation         float64 `json:"standardDeviation"`
	SharpeRatio               float64 `json:"sharpeRatio"`
}

type DealsAndOrders struct {
	Orders []interface{} `json:"orders"`
	Deals  []struct {
		Account struct {
			Type string `json:"type"`
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"account"`
		DealID            string `json:"deadId"`
		DealTime          string `json:"dealTime"`
		MarketTransaction bool   `json:"marketTransaction"`
		Orderbook         struct {
			// TODO
		} `json:"orderbook"`
		OrderId string  `json:"orderId"`
		Price   float64 `json:"price"`
		Sum     float64 `json:"sum"`
		Type    string  `json:"type"`
		Volume  int     `json:"Volume"`
	} `json:"deals"`
	Accounts []struct {
		Type string `json:"type"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"accounts"`
	ReservedAmount float64 `json:"reservedAmount"`
}

type Search struct {
	TotalNumberOfHits int `json:"totalNumberOfHits"`
	Hits              []struct {
		InstrumentType string `json:"instrumentType"`
		NumberOfHits   int    `json:"numberOfHits"`
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

type Transactions struct {
	Transactions              []Transaction `json:"transactions"`
	TotalNumberOfTransactions int           `json:"totalNumberOfTransactions"`
}

type Transaction struct {
	Account struct {
		Type string `json:"type"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"account"`
	TransactionType  string  `json:"transactionType"`
	VerificationDate string  `json:"verificationDate"`
	Description      string  `json:"description"`
	Currency         string  `json:"currency"`
	Amount           float64 `json:"amount"`
	ID               string  `json:"id"`
	Sum              float64 `json:"sum,omitempty"`
	Commission       float64 `json:"commission,omitempty"`
	NoteID           string  `json:"noteId,omitempty"`
	CurrencyRate     float64 `json:"currencyRate,omitempty"`
	Orderbook        struct {
		Isin     string `json:"isin"`
		Currency string `json:"currency"`
		Name     string `json:"name"`
		FlagCode string `json:"flagCode"`
		ID       string `json:"id"`
		Type     string `json:"type"`
	} `json:"orderbook,omitempty"`
	Volume float64 `json:"volume,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
