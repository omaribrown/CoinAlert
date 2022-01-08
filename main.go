package main

import (
	"fmt"
	"github.com/joho/godotenv"
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {
	//p1 := func(w http.ResponseWriter, _ *http.Request) {
	//	io.WriteString(w, "Hello")
	//}

	//port := os.Getenv("PORT")
	http.HandleFunc("/", RootHandler)
	//log.Print("Listening on port  :" + port)
	http.HandleFunc("/cointoslack", coinToSlack)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my content.")
	fmt.Fprintln(w, r.Header)

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(body))
}
func mapToStringSlice(slice interface{}, mapFunc func(interface{}) string) []string {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("mapToStringSlice() given a non-slice type")
	}
	ret := make([]string, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = mapFunc(s.Index(i).Interface())
	}

	return ret
}
func coinToSlack(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "Running cointoslack. Manually added a 200 Status OK.")

	// * get local env's
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	Viperenv := os.Getenv("API_KEY")
	coinapi := &coinapi.Coinapi{
		API_KEY: Viperenv,
		Client:  &http.Client{},
		// * For local testing. Not validated.
	}
	fmt.Println(Viperenv)
	c := cron.New()
	fmt.Fprintf(w, "Starting cron job")
	fmt.Println("starting cron job")

	c.AddFunc("@every 1m", func() {
		ohlvcLatest := coinapi.GetCoinLatest("BTC/USD", "1MIN", "1")

		// * Stringify for slack

		//stringData := cast.ToString(ohlvcLatest)
		//fmt.Println(stringData)
		//fmt.Fprintf(w, stringData)
		fmt.Println("Crypto Data: ", ohlvcLatest)

		slackService := &slack.SlackService{
			SlackToken:     os.Getenv("SLACK_AUTH_TOKEN"),
			SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
		}

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
