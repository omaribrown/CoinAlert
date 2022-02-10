package coinapi

type DataService interface {
	GetCoinLatest() []LatestOhlcv
}
