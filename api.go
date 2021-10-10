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

type Client struct {
	username            string
	password            string
	totpSecret          string
	xSecurityToken      string
	IsAuthenticated     bool
	totpSession         TOTPAuthentication
	reAuthenticateTimer *time.Timer
	logger              *zap.SugaredLogger
	pushSubscriptionId  string
}

func (a *Client) Close() {
	if a.reAuthenticateTimer != nil {
		a.reAuthenticateTimer.Stop()
	}
}

func (a *Client) PushSubscriptionId() string {
	return a.pushSubscriptionId
}

func NewClient(logger *zap.SugaredLogger) *Client {
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
	return &Client{
		logger:     logger,
		totpSecret: totpSecret,
		username:   username,
		password:   password,
	}
}
