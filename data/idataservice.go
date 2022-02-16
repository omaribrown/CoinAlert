package coinapi

type IDataService interface {
	GetCandles(params Params) []Candle
}

type Params struct {
	Symbol string
	Period string
	Limit  string
}
