package main

import (
	"encoding/json"
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

type Latest_OHLCV struct {
	Time_Period_Start string `json:"time_period_start"`
	Time_Period_End   string `json:"time_period_end"`
	Time_Open         string `json:"time_open"`
	Time_Close        string `json:"time_close"`
	Price_Open        string `json:"price_open"`
	Price_High        string `json:"price_high"`
	Price_Low         string `json:"price_low"`
	Price_Close       string `json:"price_close"`
	Volume_Traded     string `json:"volume_traded"`
	Trades_Count      string `json:"trades_count"`
}

type Keys struct {
	API_KEY string
}

func (k Keys) GetCoinLatest(symbol string, period string) string {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", k.API_KEY).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=1")

	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body)
	// data, _ := json.Unmarshal(resp.Body)
	// fmt.Println(string(data))
	// return string(data)
}

func main() {

	viperenv := viperEnvVariable("API_KEY")
	k := Keys{viperenv}

	x := k.GetCoinLatest("BTC/USD", "1DAY")

	data := Latest_OHLCV{}
	json.Unmarshal([]byte(x), &data)

	fmt.Println(data)

	// coinLatest := []byte(k.GetCoinLatest("BTC/USD", "1DAY"))
	// var response []Latest_OHLCV
	// err := json.Unmarshal(coinLatest, &response)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(response)
}

// Period ID's:

// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
