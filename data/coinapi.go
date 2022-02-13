package coinapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IResty interface {
	Do(req *http.Request) (*http.Response, error)
}

// API Data
type LatestOhlcv struct {
	TimePeriodStart    string  `json:"time_period_start"`
	TimePeriodEnd      string  `json:"time_period_end"`
	TimeOpen           string  `json:"time_open"`
	TimeClose          float64 `json:"time_close"`
	PriceOpen          float64 `json:"price_open"`
	PriceHigh          float64 `json:"price_high"`
	PriceLow           float64 `json:"price_low"`
	PriceClose         float64 `json:"price_close"`
	VolumeTraded       float64 `json:"volume_traded"`
	TradesCount        int64   `json:"trades_count"`
	BollingerBandUpper float64
	BollingerBandLower float64
}

type Coinapi struct {
	API_KEY string
	Client  IResty
}

//symbol string, period string, limit string, CalculationChan chan LatestOhlcv
func (c *Coinapi) GetCoinLatest(params Params) []LatestOhlcv {

	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/ohlcv/"+params.Symbol+"/latest?period_id="+params.Period+"&limit="+params.Limit, nil)
	req.Header.Set("X-CoinAPI-Key", c.API_KEY)
	resp, err := c.Client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Newstruct []LatestOhlcv
	json.Unmarshal(body, &Newstruct)
	fmt.Println(Newstruct)
	for v, _ := range Newstruct {
		params.CalculationChan <- Newstruct[v]
	}
	return Newstruct

}
