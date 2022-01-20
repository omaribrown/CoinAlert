package triggers

import (
	"fmt"
	coinapi "github.com/omaribrown/coinalert/data"
)

type bolBandTrigger struct {
	bollingerBandCandle coinapi.LatestOhlcv
	bbCandles           []coinapi.LatestOhlcv
}

func (b *bolBandTrigger) spotBreakout(bollBand chan coinapi.LatestOhlcv) {
	// take in a candle from channel
	// add it to slice
	b.bbCandles = append(b.bbCandles, <-bollBand)
	bbToSlack := make(chan coinapi.LatestOhlcv, 2)
	// compare the close price to any previous bol high or bol close
	for _, candle := range b.bbCandles {
		if candle.PriceClose > candle.BollingerBandUpper {
			bbToSlack <- candle
			fmt.Println("broke upper band at close: ", bbToSlack)
		} else if candle.PriceClose < candle.BollingerBandLower {
			bbToSlack <- candle
			fmt.Println("broke lower band at close: ", bbToSlack)
		}
	}
	// send breakout candle through channel to slack file for slack notification
}
