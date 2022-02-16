package triggers

import (
	"encoding/json"
	"fmt"
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"log"
)
import _ "github.com/joho/godotenv/autoload"

type SlackTrigger struct {
	message          slack.SlackMessage
	candle           coinapi.Candle
	triggeredCandles []coinapi.Candle
	NotifChan        chan coinapi.Candle
	SlackService     *slack.SlackService
}

func (s *SlackTrigger) SendSignal() {
	for {
		s.triggeredCandles = append(s.triggeredCandles, <-s.NotifChan)
		slackData := <-s.NotifChan
		fmt.Println("Slacktrigger received NotifChan running...")
		stringData, err := json.Marshal(slackData)
		if err != nil {
			log.Fatal(err)
		}
		slackMessage := slack.GenerateNewMessage(string(stringData), "Lower Bol Band Breakout")
		s.SlackService.SendSlackMessage(slackMessage)
	}

}
