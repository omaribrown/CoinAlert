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
	candle           coinapi.LatestOhlcv
	triggeredCandles []coinapi.LatestOhlcv
	NotifChan        chan coinapi.LatestOhlcv
	SlackService     *slack.SlackService
}

func (s *SlackTrigger) SendSignal() {
	for {
		fmt.Println("Slacktrigger received NotifChan running...")
		s.triggeredCandles = append(s.triggeredCandles, <-s.NotifChan)
		slackData := <-s.NotifChan
		stringData, err := json.Marshal(slackData)
		if err != nil {
			log.Fatal(err)
		}
		slackMessage := slack.GenerateNewMessage(string(stringData), "Lower Bol Band Breakout")
		s.SlackService.SendSlackMessage(slackMessage)
	}

}
