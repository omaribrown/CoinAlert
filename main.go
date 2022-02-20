package main

import (
	"fmt"
	"github.com/omaribrown/coinalert/calulations"
	_ "github.com/omaribrown/coinalert/calulations"
	coinapi "github.com/omaribrown/coinalert/data"
	envVariables "github.com/omaribrown/coinalert/envvar"
	"github.com/omaribrown/coinalert/slack"
	"github.com/omaribrown/coinalert/triggers"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", RootHandler)

	// load env with env.go implementation
	env, err := envVariables.New(envVariables.Props{DotEnvPath: ".env"})
	if err != nil {
		log.Fatal(err)
	}

	go coinToSlack(
		selectDataService("polygon"),
		env.Get("SLACK_AUTH_TOKEN"),
		env.Get("SLACK_CHANNEL_ID"),
	)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func selectDataService(service string) coinapi.IDataService {
	env, err := envVariables.New(envVariables.Props{DotEnvPath: ".env"})
	if err != nil {
		log.Fatal(err)
	}
	var dataService coinapi.IDataService
	switch service {
	case "polygon":
		PolygonAPIKey := env.Get("POLY_API_KEY")
		dataService = &coinapi.Polygon{
			API_KEY: PolygonAPIKey,
			Client:  &http.Client{},
		}
	case "coinapi":
		CoinAPIKey := env.Get("API_KEY")
		dataService = &coinapi.Coinapi{
			API_KEY: CoinAPIKey,
			Client:  &http.Client{},
		}
	}
	return dataService
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(body))
}

func coinToSlack(dataService coinapi.IDataService, slackToken string, slackChannelID string) {
	calculationChan := make(chan coinapi.Candle, 60)
	TriggerChan := make(chan coinapi.Candle)
	notifChan := make(chan coinapi.Candle)

	calculator := &calulations.Calculations{
		CalculationChan: calculationChan,
		TriggerChan:     TriggerChan,
	}
	strategy := &triggers.BolBandTriggers{
		TriggerChan: TriggerChan,
		NotifChan:   notifChan,
	}

	notification := &triggers.SlackTrigger{
		NotifChan: notifChan,
		SlackService: &slack.SlackService{
			SlackToken:     slackToken,
			SlackChannelID: slackChannelID,
		},
	}

	c := cron.New()
	fmt.Println("starting cron job")

	go calculator.SendToCalc()
	go strategy.LowerBbBreakout()
	go notification.SendSignal()

	c.AddFunc("@every 1m", func() {
		//go coinapi.GetCandles("ETH/USD", "1MIN", "60", calculationChan)
		params := coinapi.Params{
			Symbol: "ETHUSD",
			Period: "1MIN",
			Limit:  "60",
		}

		candles := dataService.GetCandles(params)

		for _, candle := range candles {
			calculationChan <- candle
			//fmt.Println("Candle received, sending to calc chan", candle)
		}
	})
	c.Start()
	select {}
}

// Period ID's:
// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
