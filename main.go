package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/resty.v0"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

type Keys struct {
	API_KEY string
}

func (k Keys) CoinLatest(symbol string, period string) string {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", k.API_KEY).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=1")

	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Body)
}

func main() {

	viperenv := viperEnvVariable("API_KEY")
	k := Keys{viperenv}

	fmt.Println(k.CoinLatest("BTC/USD", "1DAY"))

}

// Period ID's:

// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
