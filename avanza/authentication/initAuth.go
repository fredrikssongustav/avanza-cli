package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/fredrikssongustav/avanza-cli/avanza/avanzaConstants"
	"github.com/fredrikssongustav/avanza-cli/avanza/makeRequest"
	"io/ioutil"
)


type TwoFactorLogin struct {
	TransactionId string `json:"transactionId"`
	Method        string `json:"method"`
}

type Initiate2FAResponse struct {
	TwoFactorLogin TwoFactorLogin `json:"twoFactorLogin"`
}

func InitAuth(userName string, userPassword string) *Initiate2FAResponse {

	requestBody := fmt.Sprintf(`{"username":"%s", "password":"%s"}`, userName, userPassword)

	opt := makeRequest.RequestOptionsProps{
		Method: "POST",
		APIPath: avanzaConstants.URL_AUTHENTICATION_PATH,
		Body:    []byte(requestBody)}

	res := makeRequest.MakeRequest(opt, nil)

	initiate2FAResponse := &Initiate2FAResponse{}

	responseBody, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(responseBody, initiate2FAResponse)

	return initiate2FAResponse
}
