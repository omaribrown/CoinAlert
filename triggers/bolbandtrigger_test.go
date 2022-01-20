package triggers

import (
	coinapi "github.com/omaribrown/coinalert/data"
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

		bolBandTrigger.spotBreakout(testBollBand)

	})
}
