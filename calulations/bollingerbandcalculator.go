package calulations

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"math"
)

type indicators struct {
	bolBands []bolBandCalculator
}

type bolBandCalculator struct {
	bollingerBandCandle coinapi.LatestOhlcv
	bolUpper            float64
	bolLower            float64
	candles             []coinapi.LatestOhlcv
	size                int
}

type Props struct {
	size int
}

func New(props Props) *bolBandCalculator {
	size := 20
	if props.size != 0 {
		size = props.size
	}
	return &bolBandCalculator{
		size: size,
	}
}

func (b *bolBandCalculator) add(candle coinapi.LatestOhlcv, TriggerChan chan coinapi.LatestOhlcv) {
	b.candles = append(b.candles, candle)
	if len(b.candles) < b.size {
		return
	}
	standardDevs := 2.0
	movingAvg := calcSma(b.candles, b.size)
	stanDevPer := standardDev(b.candles, b.size)
	b.bolUpper = movingAvg + (stanDevPer * standardDevs)
	b.bolLower = movingAvg - (stanDevPer * standardDevs)
	b.bollingerBandCandle = coinapi.LatestOhlcv{
		TimePeriodStart:    candle.TimePeriodStart,
		TimePeriodEnd:      candle.TimePeriodEnd,
		TimeOpen:           candle.TimeOpen,
		TimeClose:          candle.TimeClose,
		PriceOpen:          candle.PriceOpen,
		PriceHigh:          candle.PriceHigh,
		PriceLow:           candle.PriceLow,
		PriceClose:         candle.PriceClose,
		VolumeTraded:       candle.VolumeTraded,
		TradesCount:        candle.TradesCount,
		BollingerBandUpper: b.bolUpper,
		BollingerBandLower: b.bolLower,
	}
	b.candles = b.candles[1:]

	TriggerChan <- b.bollingerBandCandle
}

func standardDev(data []coinapi.LatestOhlcv, size int) float64 {

	closeMean := calcSma(data, size)
	var devLessMean []float64
	var deviation float64
	var addDevs float64
	var avDevs float64
	for _, elem := range data {
		deviation = elem.PriceClose - closeMean
		deviation *= deviation
		devLessMean = append(devLessMean, deviation)
	}
	for _, x := range devLessMean {
		addDevs += x
	}
	avDevs = addDevs / float64(size)
	sqrRoot := math.Sqrt(avDevs)
	return sqrRoot
}
func calcSma(data []coinapi.LatestOhlcv, size int) float64 {
	sum := 0.0
	sma := 0.0
	for _, elem := range data {
		sum += elem.PriceClose
	}
	sma = sum / float64(size)
	return sma
}
