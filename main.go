package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", RootHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	coinToSlack()
	//log.Fatal(http.ListenAndServe(":8080", nil))

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(body))
}

func coinToSlack() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	CoinAPIKey := os.Getenv("API_KEY")
	coinapi := &coinapi.Coinapi{
		API_KEY: CoinAPIKey,
		Client:  &http.Client{},
	}

	slackService := &slack.SlackService{
		SlackToken:     os.Getenv("SLACK_AUTH_TOKEN"),
		SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
	}
	c := cron.New()
	fmt.Println("starting cron job")

	c.AddFunc("@every 1m", func() {
		ohlvcLatest := coinapi.GetCoinLatest("BTC/USD", "1MIN", "1")

		remarshal, err := json.Marshal(ohlvcLatest)
		if err != nil {
			panic(err)
		}
		stringData := string(remarshal)

		slackService.SendSlackMessage(slack.SlackMessage{
			Pretext: "Incoming crypto data...",
			Text:    stringData,
		})
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
