package operations

import (
	"encoding/json"
	"github.com/fredrikssongustav/avanza-cli/avanza/authentication"
	"github.com/fredrikssongustav/avanza-cli/avanza/avanzaConstants"
	"github.com/fredrikssongustav/avanza-cli/avanza/makeRequest"
	"io/ioutil"
)

type Position struct {
	AccountName          string  `json:"accountName"`
	AccountType          string  `json:"accountType"`
	Depositable          bool    `json:"depositable"`
	AccountId            float64 `json:"accountId"`
	Volume               float64 `json:"volume"`
	AverageAcquiredPrice float64 `json:"averageAcquiredPrice"`
	ProfitPercent        float64 `json:"profitPercent"`
	AcquiredValue        float64 `json:"acquiredValue"`
	Profit               float64 `json:"profit"`
	Value                float64 `json:"value"`
	Currency             string  `json:"currency"`
	OrderbookId          string  `json:"orderbookId"`
	Tradable             bool    `json:"tradable"`
	LastPrice            float64 `json:"lastPrice"`
	LastPriceUpdated     string  `json:"lastPriceUpdated"`
	Change               float64 `json:"change"`
	ChangePercent        float64 `json:"changePercent"`
	FlagCode             string  `json:"flagCode"`
	Name                 string  `json:"name"`
}

type instruments struct {
	InstrumentType string     `json:"instrumentType"`
	Positions      []Position `json:"positions"`
}

type PositionsResponse struct {
	InstrumentPositions []instruments `json:"instrumentPositions"`
}

func GetPositions(auth authentication.AuthenticationResponse) *PositionsResponse {

	opt := makeRequest.RequestOptionsProps{Method: "GET",
		APIPath: avanzaConstants.URL_POSITIONS_PATH}

	res := makeRequest.MakeRequest(opt, []makeRequest.Header{{
		Prop: "X-SecurityToken",
		Val:  auth.SecurityToken,
	}, {
		Prop: "X-AuthenticationSession",
		Val:  auth.AuthenticationSession,
	}})

	positions := &PositionsResponse{}

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, positions)

	return positions
}
