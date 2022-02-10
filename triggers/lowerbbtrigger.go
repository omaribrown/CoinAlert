package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
)

type BolBandTriggers struct {
	TriggerChan chan coinapi.LatestOhlcv
	NotifChan   chan coinapi.LatestOhlcv
}

func (b *BolBandTriggers) LowerBbBreakout() {
	for {
		bbData := <-b.TriggerChan
		if bbData.PriceClose < bbData.BollingerBandLower {
			b.NotifChan <- bbData
		}
	}
}
