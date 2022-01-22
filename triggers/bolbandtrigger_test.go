package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpotBreakout(t *testing.T) {
	t.Run("Should receive & store candles while looking for any breakouts", func(t *testing.T) {
		testBollBand := make(chan coinapi.LatestOhlcv, 20)
		testBollBand <- coinapi.LatestOhlcv{
			PriceClose:         100,
			BollingerBandUpper: 95,
			BollingerBandLower: 70,
		}

		testBol := new(bolBandTrigger)
		testBol.spotBreakout(testBollBand)
		//bolBandTrigger.spotBreakout(testBollBand)

		testResult := testBol.bollingerBandCandle
		assert.Equal(t, 0.0, testResult)
	})
}
