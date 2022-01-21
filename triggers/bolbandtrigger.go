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
	// compare the close price to any previous bol high or bol close
	for _, candle := range b.bbCandles {
		// send breakout candle through channel to slack file for slack notification
		if candle.PriceClose > candle.BollingerBandUpper {
			marshal, err := json.Marshal(candle)
			if err != nil {
				log.Fatal(err)
			}
			bbSlackData <- candle
			bbSlackMessage <- slack.SlackMessage{
				Title:   "Bollinger Band Breakout!",
				Pretext: string(marshal),
				Text:    "Target Entry: , Stop Loss: , Price Target: ",
			}
			fmt.Println("broke upper band at close: ", bbSlackData)
		} else if candle.PriceClose < candle.BollingerBandLower {
			marshal, err := json.Marshal(candle)
			if err != nil {
				log.Fatal(err)
			}
			bbSlackData <- candle
			bbSlackMessage <- slack.SlackMessage{
				Title:   "Bollinger Band Breakout!",
				Pretext: string(marshal),
				Text:    "Target Entry: , Stop Loss: , Price Target: ",
			}
			fmt.Println("broke lower band at close: ", bbSlackData)
		}
	}
}
