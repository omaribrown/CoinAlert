package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
)

type BolBandTriggers struct {
	bollingerBandCandle coinapi.LatestOhlcv
}

func (b *BolBandTriggers) LowerBbBreakout(TriggerChan chan coinapi.LatestOhlcv, NotifChan chan coinapi.LatestOhlcv) {
	for {
		bbData := <-TriggerChan
		if bbData.PriceClose < bbData.BollingerBandLower {
			NotifChan <- bbData
		}
	}
}
