package goavanza

import (
	"os"
	"time"

	"go.uber.org/zap"
)

const Url = "https://www.avanza.se"
const UserAgent = "Avanza API client"
const MaxInactiveMinutes = 24
const layoutISO = "2006-01-02"

type AvanzaApi interface {
	SubscribeOrders(ids []string) error
	Listen() error
	PlaceOrder(*OrderOptions) (*OrderActionResponse, error)
	EditOrder(instrumentType Instrument, orderId string, options *OrderOptions) (*OrderActionResponse, error)
	DeleteOrder(accountId string, orderId string) (*OrderActionResponse, error)
	GetPositions() (*Positions, error)
	GetOverview() (*Overview, error)
	GetAccountOverview(accountId string) (*AccountOverview, error)
	GetDealsAndOrders() (*DealsAndOrders, error)
	GetTransactions(accountOrTransactionType string, options TransactionOptions) (*Transactions, error)
	Search(query string, instrumentType Instrument) (*Search, error)
	Authenticate() error
	Close()
}

type api struct {
	username            string
	password            string
	totpSecret          string
	xSecurityToken      string
	IsAuthenticated     bool
	totpSession         TOTPAuthentication
	reAuthenticateTimer *time.Timer
	logger              *zap.SugaredLogger
	websocketConnection *websocketConnection
}

func (a *api) Close() {
	if a.reAuthenticateTimer != nil {
		a.reAuthenticateTimer.Stop()
	}
	if a.websocketConnection.socket != nil {
		a.websocketConnection.socket.Close()
	}
}

func NewApi(logger *zap.SugaredLogger) AvanzaApi {
	totpSecret := os.Getenv("AVANZA_TOTP_SECRET")
	if totpSecret == "" {
		logger.Fatalf("AVANZA_TOTP_SECRET environment variable not set")
	}
	username := os.Getenv("AVANZA_USERNAME")
	if username == "" {
		logger.Fatalf("AVANZA_USERNAME environment variable not set")
	}
	password := os.Getenv("AVANZA_PASSWORD")
	if password == "" {
		logger.Fatalf("AVANZA_PASSWORD environment variable not set")
	}
	return &api{
		logger:     logger,
		totpSecret: totpSecret,
		username:   username,
		password:   password,
		websocketConnection: &websocketConnection{
			socketMessageCount: 1,
		},
	}
}
