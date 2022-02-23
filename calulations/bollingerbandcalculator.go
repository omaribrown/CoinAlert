package calulations

import (
	"encoding/csv"
	"fmt"
	coinapi "github.com/omaribrown/coinalert/data"
	"go.uber.org/zap"
	"log"
	"math"
	"os"
)

type indicators struct {
	bolBands []bolBandCalculator
}

type bolBandCalculator struct {
	bollingerBandCandle  coinapi.Candle
	bollingerBandCandles []coinapi.Candle
	bolUpper             float64
	bolLower             float64
	candles              []coinapi.Candle
	size                 int
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

func (b *bolBandCalculator) add(candle coinapi.Candle, TriggerChan chan coinapi.Candle) {
	b.candles = append(b.candles, candle)
	if len(b.candles) < b.size {
		zap.S().Warnf("Not enough candles to calculate Bollinger Bands. Length required: %+v, Actual length: %v.", b.size, len(b.candles))
		b.bollingerBandCandles = append(b.bollingerBandCandles, candle)
		return
	}

	standardDevs := 2.0
	movingAvg := calcSma(b.candles, b.size)

	stanDevPer := standardDev(b.candles, b.size)
	b.bolUpper = movingAvg + (stanDevPer * standardDevs)
	//fmt.Println("Upper: ", b.bolUpper)
	b.bolLower = movingAvg - (stanDevPer * standardDevs)
	//fmt.Println("Lower: ", b.bolLower)

	b.bollingerBandCandle = coinapi.Candle{
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

	b.bollingerBandCandles = append(b.bollingerBandCandles, b.bollingerBandCandle)
	zap.S().Info("Sending Bollinger Band Candle to TriggerChan ==> ", b.bollingerBandCandle)
	//csvData(b.bollingerBandCandles)

	b.candles = b.candles[1:]

	TriggerChan <- b.bollingerBandCandle

}

func csvData(candles []coinapi.Candle) {
	multiDim := [][]string{}
	multiDim = append(multiDim, []string{"Time Open", "Price Open", "Price High", "Price Low", "Price Close", "Upper BB", "Lower BB", "Moving Average", "Standard Deviation/Period"})
	for _, candle := range candles {
		row := []string{candle.TimePeriodStart, fmt.Sprintf("%f", candle.PriceOpen), fmt.Sprintf("%f", candle.PriceHigh), fmt.Sprintf("%f", candle.PriceLow), fmt.Sprintf("%f", candle.PriceClose), fmt.Sprintf("%f", candle.BollingerBandUpper), fmt.Sprintf("%f", candle.BollingerBandLower)}
		multiDim = append(multiDim, row)
	}

	f, err := os.Create("data.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(multiDim) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func standardDev(data []coinapi.Candle, size int) float64 {

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
func calcSma(data []coinapi.Candle, size int) float64 {
	sum := 0.0
	sma := 0.0
	for _, elem := range data {
		sum += elem.PriceClose
	}
	sma = sum / float64(size)
	return sma
}
