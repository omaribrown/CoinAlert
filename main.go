package main

import (
	"fmt"
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"github.com/robfig/cron"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world")
	fmt.Println("starting cron job")

	slackService := &slack.SlackService{
		SlackToken:     os.Getenv("SLACK_AUTH_TOKEN"),
		SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
	}

	Viperenv := os.Getenv("API_KEY")
	coinapi := &coinapi.Coinapi{
		API_KEY: Viperenv,
		Client:  &http.Client{},
	}

	c := cron.New()

	c.AddFunc("@every 1m", func() {

		ohlvcLatest := coinapi.GetCoinLatest("BTC/USD", "1MIN", "1")

		stringData := cast.ToString(ohlvcLatest)
		fmt.Println("Crypto Data: ", ohlvcLatest)

		slackService.SendSlackMessage(slack.SlackMessage{
			Pretext: "Incoming crypto data...",
			Text:    stringData,
		})
	})
	c.Start()
	select {}
}
func main() {
	//Viperenv := envVariables.ViperEnvVariable("API_KEY")
	//coinapi := &coinapi.Coinapi{
	//	API_KEY: Viperenv,
	//	Client:  &http.Client{},
	//}

	port := os.Getenv("PORT")
	http.HandleFunc("/", helloWorld)
	log.Print("Listening on port  :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	//Viperenv := envVariables.ViperEnvVariable("API_KEY")
	//coinapi := &coinapi.Coinapi{API_KEY: Viperenv}
	//

	//log.Info("Create new cron")

}

// Period ID's:

// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
