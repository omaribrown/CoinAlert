package main

import (
	"fmt"
	"log"

	"gopkg.in/resty.v0"
)

type Keys struct {
	API_KEY string
}

func (k Keys) CoinLatest(symbol string, period string) string {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", k.API_KEY).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=3")

	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Body)
}

func main() {

	k := Keys{"ADE71FB5-9455-4547-9272-9C2D491AA630"}

	fmt.Println(k.CoinLatest("BTC/USD", "1DAY"))

}

// Period ID's:

// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
