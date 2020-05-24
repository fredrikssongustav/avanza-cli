package authentication

import (
	"github.com/xlzd/gotp"
)

type AuthenticationResponse struct {
	AuthenticationSession string `json:"authenticationSession"`
	PushSubscriptionId    string `json:"pushSubscriptionId"`
	CustomerId            string `json:"customerId"`
	SecurityToken         string `json:"securityToken"`
}

type AvanzaCredentials struct {
	Username string
	Password string
	Secret   string
}

func Authentication(credentials AvanzaCredentials) *AuthenticationResponse {
	// Authenticating with 2FA require two API calls:
	initAuthResponse := InitAuth(credentials.Username, credentials.Password)

	totp := gotp.NewDefaultTOTP(credentials.Secret)
	body, headers, _ := FinalizeAuth(initAuthResponse.TwoFactorLogin.TransactionId, totp.Now())

	authenticationResponse := &AuthenticationResponse{
		body.AuthenticationSession,
		body.PushSubscriptionId,
		body.CustomerId,
		headers.Get("x-securitytoken")}

	return authenticationResponse
}

