package main

import (
	"fmt"

	coinapi "github.com/omaribrown/coinalert/data"
	envVariables "github.com/omaribrown/coinalert/envvar"
)

func main() {
	Viperenv := envVariables.ViperEnvVariable("API_KEY")
	coinapi := &coinapi.Coinapi{API_KEY: Viperenv}

	ohlvcLatest := coinapi.GetCoinLatest("BTC/USD", "1DAY", "3")
	fmt.Println("Negative: ", ohlvcLatest)

}

// Period ID's:

// Second	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
// Minute	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
// Hour	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
// Day	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
// Month	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
// Year	1YRS, 2YRS, 3YRS, 4YRS, 5YRS
