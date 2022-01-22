package triggers

import (
	"encoding/json"
	"fmt"
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/omaribrown/coinalert/slack"
	"log"
)

type bolBandTrigger struct {
	bollingerBandCandle coinapi.LatestOhlcv
	bbCandles           []coinapi.LatestOhlcv
}

func (b *bolBandTrigger) spotBreakout(bollBand chan coinapi.LatestOhlcv) {
	// take in a candle from channel, add it to slice
	b.bbCandles = append(b.bbCandles, <-bollBand)
	bbSlackData := make(chan coinapi.LatestOhlcv, 2)
	bbSlackMessage := make(chan slack.SlackMessage, 2)
	slackTrigger := new(slackTrigger)
	// compare the close price to any previous bol high or bol close
	for _, candle := range b.bbCandles {
		// slackTrigger breakout candle through channel to slack file for slack notification
		if candle.PriceClose > candle.BollingerBandUpper {
			fmt.Println("Price broke upper band at close: ", candle.PriceClose, " Sending data to slack...")
			bbSlackData <- candle
			marshal, err := json.Marshal(candle)
			if err != nil {
				log.Fatal(err)
			}
			bbSlackMessage <- slack.SlackMessage{
				Title:   "Bollinger Band Breakout!",
				Pretext: string(marshal),
				Text:    "Target Entry: , Stop Loss: , Price Target: ",
			}

			go slackTrigger.sendSignal(bbSlackData, bbSlackMessage)

		} else if candle.PriceClose < candle.BollingerBandLower {
			bbSlackData <- candle
			fmt.Println("broke lower band at close: ", candle.PriceClose, " Sending data to slack...")
			marshal, err := json.Marshal(candle)
			if err != nil {
				log.Fatal(err)
			}
			bbSlackMessage <- slack.SlackMessage{
				Title:   "Bollinger Band Breakout!",
				Pretext: string(marshal),
				Text:    "Target Entry: , Stop Loss: , Price Target: ",
			}

			go slackTrigger.sendSignal(bbSlackData, bbSlackMessage)

		}
	}
}
