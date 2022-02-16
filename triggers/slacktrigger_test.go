package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendSignal(t *testing.T) {
	t.Run("Should send message to slack", func(t *testing.T) {
		testBollBand := make(chan coinapi.Candle, 20)
		testBollBand <- coinapi.Candle{
			PriceClose:         100,
			BollingerBandUpper: 95,
			BollingerBandLower: 70,
		}
		testMessage := make(chan slack.SlackMessage, 20)
		testMessage <- slack.SlackMessage{
			Title:   "Did this shit work?",
			Pretext: "fuck",
			Text:    "yeah",
		}

		testSend := new(SlackTrigger)
		SlackService := &slack.SlackService{
			SlackToken:     "123",
			SlackChannelID: "123",
		}
		testSend.SendSignal(testBollBand, SlackService)

		assert.Equal(t, "pre", testMessage)
	})
}
