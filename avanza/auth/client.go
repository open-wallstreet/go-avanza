package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"github.com/pquerna/otp/totp"
)

const (
	UserCredentialsPath = "/_api/authentication/sessions/usercredentials"
	TOTPPath            = "/_api/authentication/sessions/totp"
)
const MinInactiveMinutes = 30
const MaxInactiveMinutes = 60 * 24

// AuthClient defines a REST Client for Authentication API:s
type AuthClient struct {
	*client.Client
	reAuthenticateTimer   *time.Timer
	username              string
	password              string
	totpSecret            string
	authenticationTimeout int
}

func NewAuthClient(c *client.Client) *AuthClient {
	return &AuthClient{
		Client:                c,
		authenticationTimeout: MaxInactiveMinutes,
	}
}

// Authenticate will authenticate against Avanza using a TOTP secret for 2FA. Currently only supported version.
// It will also set up a refresh every MaxInactiveMinutes - 1 minute to keep the session up.
//
// Make sure not to save the totpSecret into your code
func (a *AuthClient) Authenticate(ctx context.Context, username, password, totpSecret string, options ...models.RequestOption) (*models.AuthenticateTOTPResponse, error) {
	if !(a.authenticationTimeout <= MaxInactiveMinutes || a.authenticationTimeout >= MinInactiveMinutes) {
		return nil, fmt.Errorf("session timeout (%d) not in range  %d - %d minutes", a.authenticationTimeout, MinInactiveMinutes, MaxInactiveMinutes)
	}

	res := &models.UserCredentialsResponse{}
	a.totpSecret = totpSecret
	a.username = username
	a.password = password
	params := &models.UserCredentialsParams{
		Username:           username,
		Password:           password,
		MaxInactiveMinutes: a.authenticationTimeout,
	}
	err := a.Call(ctx, http.MethodPost, UserCredentialsPath, params, res, options...)
	if res.TwoFactorLogin.Method != "TOTP" {
		return nil, fmt.Errorf("TwoFactorLogin method: %s is not supported", res.TwoFactorLogin.Method)
	}
	authenticateTotp, err := a.authenticateTotp(ctx)
	if err != nil {
		return authenticateTotp, err
	}
	reAuthDuration := time.Duration(a.authenticationTimeout-1) * time.Minute
	a.reAuthenticateTimer = time.AfterFunc(reAuthDuration, a.reAuthenticate)
	return authenticateTotp, err
}

func (a *AuthClient) authenticateTotp(ctx context.Context, options ...models.RequestOption) (*models.AuthenticateTOTPResponse, error) {
	totpCode, err := totp.GenerateCode(a.totpSecret, time.Now())
	if err != nil {
		return nil, err
	}
	res := &models.AuthenticateTOTPResponse{}
	params := &models.AuthenticateTOTPParams{
		TOTPCode: totpCode,
		Method:   "TOTP",
	}
	err = a.Call(ctx, http.MethodPost, TOTPPath, params, res, options...)
	a.Client.AuthTokens = &models.AuthSessionTokens{
		AuthenticationSession: res.AuthenticationSession,
		PushSubscriptionId:    res.PushSubscriptionId,
	}
	return res, err
}

func (a *AuthClient) reAuthenticate() {
	_, err := a.Authenticate(context.Background(), a.username, a.password, a.totpSecret)
	if err != nil {
		log.Println(fmt.Sprintf("failed to reauthenticate %v", err))
	}
}

// Close Will close the reAuthenticateTimer timer if it has been initialized. Called from parent Close method.
// No need to call it manually
func (a *AuthClient) Close() {
	if a.reAuthenticateTimer != nil {
		a.reAuthenticateTimer.Stop()
	}
}
