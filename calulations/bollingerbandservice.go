package calulations

import coinapi "github.com/omaribrown/coinalert/data"

type Calculations struct {
	CalculationChan chan coinapi.LatestOhlcv
}

func (c *Calculations) SendToCalc(CalculationChan chan coinapi.LatestOhlcv, TriggerChan chan coinapi.LatestOhlcv) {
	bolCalc := New(Props{size: 20})
	for {
		calcData := <-CalculationChan
		bolCalc.add(calcData, TriggerChan)
	}
}
