package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"os"
)

type slackTrigger struct {
	message          slack.SlackMessage
	candle           coinapi.LatestOhlcv
	triggeredCandles []coinapi.LatestOhlcv
}

func (s *slackTrigger) sendSignal(candle chan coinapi.LatestOhlcv, message chan slack.SlackMessage) {
	// Store triggered candles
	s.triggeredCandles = append(s.triggeredCandles, <-candle)
	// Create & send slack message
	var slackMessage slack.SlackMessage
	slackMessage = <-message

	slackService := &slack.SlackService{
		SlackToken:     os.Getenv("SLACK_AUTH_TOKEN"),
		SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
	}

	slackService.SendSlackMessage(slackMessage)
}
