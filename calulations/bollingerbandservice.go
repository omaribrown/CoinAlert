package calulations

import (
	coinapi "github.com/omaribrown/coinalert/data"
)

type Calculations struct {
	CalculationChan chan coinapi.LatestOhlcv
	TriggerChan     chan coinapi.LatestOhlcv
}

func (c *Calculations) SendToCalc() {

	bolCalc := New(Props{size: 20})
	for {
		calcData := <-c.CalculationChan
		bolCalc.add(calcData, c.TriggerChan)
	}
}
