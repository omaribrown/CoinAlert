package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"go.uber.org/zap"
)

type BolBandTriggers struct {
	TriggerChan chan coinapi.Candle
	NotifChan   chan coinapi.Candle
}

func (b *BolBandTriggers) LowerBbBreakout() {
	for {
		bbData := <-b.TriggerChan
		if bbData.PriceClose < bbData.BollingerBandLower {
			zap.S().Infof("Candle at (%v) closed outside Lower Bollinger Band", bbData.TimePeriodStart)
			b.NotifChan <- bbData
		}
	}
}
