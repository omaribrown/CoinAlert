package coinapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type AggregateBars struct {
	Ticker       string    `json:"ticker"`
	Adjusted     bool      `json:"adjusted"`
	QueryCount   int       `json:"queryCount"`
	RequestId    string    `json:"request_id"`
	ResultsCount int       `json:"resultsCount"`
	Status       string    `json:"status"`
	Results      []Results `json:"results"`
}

type Results struct {
	Close         float64 `json:"c"`
	High          float64 `json:"h"`
	Low           float64 `json:"l"`
	Transactions  int64   `json:"n"`
	Open          float64 `json:"o"`
	Time          int64   `json:"t"`
	VolumeTraded  float64 `json:"v"`
	VolumeAverage float64 `json:"vw"`
}

type Polygon struct {
	API_KEY string
	Client  IResty
}

func (p *Polygon) GetCoinLatest(cryptoTicker string, multiplier string, timespan string, limit string, CalculationChan chan LatestOhlcv) []LatestOhlcv {

	url := "https://api.polygon.io/v2/aggs/ticker/X:" + cryptoTicker + "/range/" + multiplier + "/" + timespan + "/" + getTimeFormatted() + "/" + getTimeFormatted() + "?adjusted=true&sort=desc&limit=" + limit

	polyClient := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)

	}
	req.Header.Set("Authorization", "Bearer "+p.API_KEY)

	res, getErr := polyClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	bars := AggregateBars{}
	jsonErr := json.Unmarshal(body, &bars)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	var candles []LatestOhlcv
	for _, i := range bars.Results {
		candles = append(candles, LatestOhlcv{
			TimePeriodStart:    unixToRFC(i.Time),
			TimePeriodEnd:      "",
			TimeOpen:           "",
			TimeClose:          0,
			PriceOpen:          i.Open,
			PriceHigh:          i.High,
			PriceLow:           i.Low,
			PriceClose:         i.Close,
			VolumeTraded:       i.VolumeTraded,
			TradesCount:        0,
			BollingerBandUpper: 0,
			BollingerBandLower: 0,
		})

	}
	candles = reverseCandles(candles)
	fmt.Println(candles)
	for v, _ := range candles {
		CalculationChan <- candles[v]
	}
	return candles
}

func unixToRFC(unix int64) string {
	sec := unix / 1000
	msec := unix % 1000
	t := time.Unix(sec, msec*int64(time.Millisecond))
	//fmt.Println(t.String())
	return t.String()
}

func getTimeFormatted() string {
	t := time.Now()
	tm := t.Format("2006-01-02")
	//fmt.Println("YYYY-MM-DD : ", tm)
	return tm
}

func reverseCandles(candles []LatestOhlcv) []LatestOhlcv {
	for i := 0; i < len(candles)/2; i++ {
		j := len(candles) - i - 1
		candles[i], candles[j] = candles[j], candles[i]
	}
	return candles
}
