package coinapi

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/resty.v0"
)

// API Data
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

type KEYS struct {
	api_key string
}

func NewKey(api_key string) *KEYS {
	key := new(KEYS)
	key.api_key = api_key

	// k := KEYS{viperenv}
	//  := viperEnvVariable("API_KEY")
	// &Keys := Keys{viperenv}
	return key
}

func (k KEYS) GetCoinLatest(symbol string, period string) string {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", k.api_key).
		Get("https://rest.coinapi.io/v1/ohlcv/" + symbol + "/latest?period_id=" + period + "&limit=1")

	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Body)
	// data, _ := json.Unmarshal(resp.Body)
	// fmt.Println(string(data))
	// return string(data)
}

var Test = 1
