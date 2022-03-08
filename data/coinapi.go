package coinapi

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type IResty interface {
	Do(req *http.Request) (*http.Response, error)
}

type Candle struct {
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

type Props struct {
	API_KEY string
	Client  IResty
}

func New(props Props) *Coinapi {
	return &Coinapi{
		API_KEY: props.API_KEY,
		Client:  props.Client,
	}
}

func (c *Coinapi) GetCandles(params Params) []Candle {

	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/ohlcv/"+formatSymbol(params.Symbol)+"/latest?period_id="+params.Period+"&limit="+params.Limit, nil)
	req.Header.Set("X-CoinAPI-Key", c.API_KEY)
	zap.S().Debug("Trying to get response at url ==> ", req)
	resp, err := c.Client.Do(req)
	if err != nil {
		zap.S().Fatal("Failed to get response ==> ", err)
	} else {
		zap.S().Debug("Got response status ==> ", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.S().Fatal("Failed to read body ==> ", err)

	}

	var Newstruct []Candle
	json.Unmarshal(body, &Newstruct)
	return Newstruct

}

func formatSymbol(symbol string) string {
	index := 3
	insertSlash := symbol[:index] + "/" + symbol[index:]
	return insertSlash
}
