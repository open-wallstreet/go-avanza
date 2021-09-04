package goavanza

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func (a *api) Authenticate() error {
	data := map[string]interface{}{
		"username":           a.username,
		"password":           a.password,
		"maxInactiveMinutes": MaxInactiveMinutes * 60,
	}
	body, _, err := a.request("/_api/authentication/sessions/usercredentials", "POST", data, RequestOptions{})
	a.logger.Infof("d %s %s", body, err)
	if err != nil {
		return err
	}
	var authResponse Authentication
	if err := json.Unmarshal([]byte(body), &authResponse); err != nil {
		a.logger.Errorf("failed to marshal %v", err)
		return nil
	}

	if authResponse.TwoFactorLogin.Method != "TOTP" {
		return fmt.Errorf("unsupported sign in type %s", authResponse.TwoFactorLogin.Method)
	}

	a.authenticateTotp(authResponse.TwoFactorLogin.TransactionID)

	return nil
}

func (a *api) authenticateTotp(transactionId string) (TOTPAuthentication, error) {
	totpCode, err := totp.GenerateCode(a.totpSecret, time.Now())
	if err != nil {
		a.logger.Error(err)
	}
	data := map[string]interface{}{
		"totpCode": totpCode,
		"method":   "TOTP",
	}
	body, response, err := a.request("/_api/authentication/sessions/totp", "POST", data, RequestOptions{
		headers: map[string]string{
			"Cookie": fmt.Sprintf("AZAMFATRANSACTION=%s", transactionId),
		},
	})
	if err != nil {
		a.logger.Error(err)
	}

	var totpResponse TOTPAuthentication
	if err := json.Unmarshal([]byte(body), &totpResponse); err != nil {
		a.logger.Errorf("failed to marshal %v", err)
	}

	a.logger.Debug("Logged In")
	a.IsAuthenticated = true
	a.xSecurityToken = response.Header.Get("x-securitytoken")
	a.totpSession = totpResponse

	a.reAuthenticateTimer = time.AfterFunc((MaxInactiveMinutes-1)*time.Minute, a.reAuthenticate)

	return totpResponse, nil
}

func (a *api) reAuthenticate() {
	a.logger.Debug("reAuthenticate to avanza")
	a.Authenticate()
}
