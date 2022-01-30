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
}

func (s *SlackTrigger) SendSignal(NotifChan chan coinapi.LatestOhlcv, SlackService *slack.SlackService) {
	// Store triggered candles
	for {
		fmt.Println("Slacktrigger received NotifChan running...")
		s.triggeredCandles = append(s.triggeredCandles, <-NotifChan)
		slackData := <-NotifChan
		stringData, err := json.Marshal(slackData)
		if err != nil {
			log.Fatal(err)
		}
		slackMessage := slack.GenerateNewMessage(string(stringData), "Lower Bol Band Breakout")
		SlackService.SendSlackMessage(slackMessage)
	}

}
