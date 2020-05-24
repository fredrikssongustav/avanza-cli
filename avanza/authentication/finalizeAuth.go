package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/fredrikssongustav/avanza-cli/avanza/avanzaConstants"
	"github.com/fredrikssongustav/avanza-cli/avanza/makeRequest"
	"io/ioutil"
	"net/http"
)


type Final2FAResponse struct {
	AuthenticationSession string `json:"authenticationSession"`
	PushSubscriptionId    string `json:"pushSubscriptionId"`
	RegistrationComplete  bool   `json:"registrationComplete"`
	CustomerId            string `json:"customerId"`
}

func FinalizeAuth(TransactionId2FA string, totpSecret string) (*Final2FAResponse, http.Header, error) {

	requestBody := fmt.Sprintf(`{"method":"TOTP", "totpCode":"%s"}`, totpSecret)

	headers := []makeRequest.Header{
		{Prop: "Cookie",
			Val: fmt.Sprintf(`AZAMFATRANSACTION=%s`, TransactionId2FA)},
	}

	opt := makeRequest.RequestOptionsProps{Method: "POST",
		APIPath: avanzaConstants.URL_TOTP_PATH,
		Body:    []byte(requestBody)}

	response := makeRequest.MakeRequest(opt, headers)

	final2FAResponse := &Final2FAResponse{}

	responseHeaders := response.Header

	responseBody, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseBody, final2FAResponse)

	return final2FAResponse, responseHeaders, nil
}
