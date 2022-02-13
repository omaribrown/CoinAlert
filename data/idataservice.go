package coinapi

type IDataService interface {
	GetCoinLatest() []LatestOhlcv
}

type Params struct {
	Symbol          string
	Period          string
	Limit           string
	CalculationChan chan LatestOhlcv
}
