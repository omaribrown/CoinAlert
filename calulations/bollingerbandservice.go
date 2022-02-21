package calulations

import (
	coinapi "github.com/omaribrown/coinalert/data"
	"go.uber.org/zap"
)

type Calculations struct {
	CalculationChan chan coinapi.Candle
	TriggerChan     chan coinapi.Candle
}

func (c *Calculations) SendToCalc() {

	bolCalc := New(Props{size: 20})
	for {
		calcData := <-c.CalculationChan
		zap.S().Info("Calculation Chan received candle... sending to Bollinger Band Calculator")
		bolCalc.add(calcData, c.TriggerChan)
	}
}
