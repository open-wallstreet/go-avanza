package models

type UserCredentialsResponse struct {
	TwoFactorLogin struct {
		TransactionID string `json:"transactionId"`
		Method        string `json:"method"`
	} `json:"twoFactorLogin"`
}

type UserCredentialsParams struct {
	Username           string `json:"username" validate:"required"`
	Password           string `json:"password" validate:"required"`
	MaxInactiveMinutes int    `json:"maxInactiveMinutes"`
}

type AuthenticateTOTPResponse struct {
	AuthenticationSession string `json:"authenticationSession"`
	PushSubscriptionId    string `json:"pushSubscriptionId"`
	CustomerId            string `json:"customerId"`
	RegistrationComplete  bool   `json:"registrationComplete"`
}

type AuthenticateTOTPParams struct {
	TOTPCode string `json:"totpCode" validate:"required"`
	Method   string `json:"method" validate:"required"`
}

type AuthSessionTokens struct {
	AuthenticationSession string `json:"authenticationSession"`
	PushSubscriptionId    string `json:"pushSubscriptionId"`
}
