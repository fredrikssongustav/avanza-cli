package operations

import (
	"encoding/json"
	"fmt"
	"github.com/fredrikssongustav/avanza-cli/avanza/avanzaConstants"
	"github.com/fredrikssongustav/avanza-cli/avanza/makeRequest"
	"io/ioutil"
)

type ChartResponse struct {
	DataPoints         [][]float64   `json:"dataPoints"`
	TrendSeries        [][]float64   `json:"trendSeries"`
	AllowedResolutions []string      `json:"allowedResolutions"`
	DefaultResolution  string        `json:"defaultResolution"`
	OwnersPoints       []interface{} `json:"ownersPoints"`
	ChangePercent      float64       `json:"changePercent"`
	High               float64       `json:"high"`
	LastPrice          float64       `json:"lastPrice"`
	Low                float64       `json:"low"`
}

func GetChartPosition(orderBookId string) *ChartResponse {
	requestBody := fmt.Sprintf(`{"orderbookId": %s,"chartType": "AREA","widthOfPlotContainer": 0,"chartResolution": "HOUR","timePeriod": "week","ta": [],"compareIds": []}`, orderBookId)

	opt := makeRequest.RequestOptionsProps{Method: "GET",
		APIPath: avanzaConstants.URL_CHARTDATA_PATH, Body: []byte(requestBody)}

	res := makeRequest.MakeRequest(opt, []makeRequest.Header{{
		Prop: "Content-Type",
		Val:  "application/json",
	},{
		Prop: "Accept",
		Val:  "*/*",
	}})

	ChartResponse := &ChartResponse{}

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, ChartResponse)

	return ChartResponse
}
