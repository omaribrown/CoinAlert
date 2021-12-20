package coinapi

import (
	"encoding/json"
	"log"

	"gopkg.in/resty.v0"
)

type IResty interface {
	Get(string) (*resty.Response, error)
	SetHeader(header, value string) *resty.Request
}

// API Data
type LatestOhlcv struct {
	TimePeriodStart string  `json:"time_period_start"`
	TimePeriodEnd   string  `json:"time_period_end"`
	TimeOpen        string  `json:"time_open"`
	TimeClose       float64 `json:"time_close"`
	PriceOpen       float64 `json:"price_open"`
	PriceHigh       float64 `json:"price_high"`
	PriceLow        float64 `json:"price_low"`
	PriceClose      float64 `json:"price_close"`
	VolumeTraded    float64 `json:"volume_traded"`
	TradesCount     int64   `json:"trades_count"`
}

type Coinapi struct {
	API_KEY string
	Resty   IResty
}

func (c *Coinapi) GetCoinLatest(symbol string, period string, limit string) []LatestOhlcv {

	resp, err := c.Resty.
		SetHeader("X-CoinAPI-Key", c.API_KEY).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=" + limit)

	if err != nil {
		log.Fatal(err)
	}

	var Newstruct []LatestOhlcv
	json.Unmarshal(resp.Body, &Newstruct)
	return Newstruct

}
