package coinapi

import (
	"encoding/json"
	"log"

	"gopkg.in/resty.v0"
)

// API Data
type Latest_OHLCV struct {
	Time_Period_Start string  `json:"time_period_start"`
	Time_Period_End   string  `json:"time_period_end"`
	Time_Open         string  `json:"time_open"`
	Time_Close        float64 `json:"time_close"`
	Price_Open        float64 `json:"price_open"`
	Price_High        float64 `json:"price_high"`
	Price_Low         float64 `json:"price_low"`
	Price_Close       float64 `json:"price_close"`
	Volume_Traded     float64 `json:"volume_traded"`
	Trades_Count      int64   `json:"trades_count"`
}

type KEYS struct {
	api_key string
}

type Coinapi struct {
	API_KEY string
}

func (c *Coinapi) GetCoinLatest(symbol string, period string, limit string) []Latest_OHLCV {

	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", c.API_KEY).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=" + limit)

	if err != nil {
		log.Fatal(err)
	}

	var Newstruct []Latest_OHLCV
	json.Unmarshal([]byte(resp.Body), &Newstruct)
	return Newstruct

}
