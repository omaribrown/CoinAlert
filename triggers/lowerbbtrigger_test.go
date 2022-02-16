package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBolBandTriggers_LowerBbBreakout(t *testing.T) {
	t.Run("Should spot candle where close is below lower bb", func(t *testing.T) {
		TriggerChan := make(chan coinapi.Candle)
		NotifChan := make(chan coinapi.Candle)
		var NotifHolder coinapi.Candle

		TriggerChan <- coinapi.Candle{
			TimePeriodStart:    "",
			TimePeriodEnd:      "",
			TimeOpen:           "",
			TimeClose:          0,
			PriceOpen:          100,
			PriceHigh:          200,
			PriceLow:           10,
			PriceClose:         50,
			VolumeTraded:       0,
			TradesCount:        0,
			BollingerBandUpper: 0,
			BollingerBandLower: 0,
		}

		newSpotter := new(BolBandTriggers)

		newSpotter.LowerBbBreakout(TriggerChan, NotifChan)
		NotifHolder = <-NotifChan

		testVar := NotifHolder.PriceOpen

		assert.Equal(t, 100.0, testVar)

	})
}
