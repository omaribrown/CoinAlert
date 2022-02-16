package coinapi

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

type httpMock struct {
	mock.Mock
}

func (_m *httpMock) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	return ret.Get(0).(*http.Response), ret.Error(1)
}

func TestCoinapi_GetCoinLatest(t *testing.T) {
	// ARRANGE
	CalculationChan := make(chan Candle, 60)
	restyMock := new(httpMock)
	sampleResponse := "[\n  {\n    \"time_period_start\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:32:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:31:01.0000000Z\",\n    \"time_close\": \"2017-08-09T14:31:46.0000000Z\",\n    \"price_open\": 3255.590000000,\n    \"price_high\": 3255.590000000,\n    \"price_low\": 3244.740000000,\n    \"price_close\": 3244.740000000,\n    \"volume_traded\": 16.903274550,\n    \"trades_count\": 31\n  },\n  {\n    \"time_period_start\": \"2017-08-09T14:30:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:30:05.0000000Z\",\n    \"time_close\": \"2017-08-09T14:30:35.0000000Z\",\n    \"price_open\": 3256.000000000,\n    \"price_high\": 3256.010000000,\n    \"price_low\": 3247.000000000,\n    \"price_close\": 3255.600000000,\n    \"volume_traded\": 58.131397920,\n    \"trades_count\": 33\n  }\n]"
	restyMock.On("Do", mock.Anything).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(sampleResponse)))}, nil)

	// ACT
	testCoinapi := &Coinapi{
		API_KEY: "",
		Client:  restyMock,
	}
	testCoinapi.GetCandles("SYM", "PERIOD", "1", CalculationChan)

	// ASSERT

}

func TestPolygon_GetCoinLatest(t *testing.T) {
	restyMock := new(httpMock)
	sampleResponse := "[\n  {\n    \"time_period_start\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:32:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:31:01.0000000Z\",\n    \"time_close\": \"2017-08-09T14:31:46.0000000Z\",\n    \"price_open\": 3255.590000000,\n    \"price_high\": 3255.590000000,\n    \"price_low\": 3244.740000000,\n    \"price_close\": 3244.740000000,\n    \"volume_traded\": 16.903274550,\n    \"trades_count\": 31\n  },\n  {\n    \"time_period_start\": \"2017-08-09T14:30:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:30:05.0000000Z\",\n    \"time_close\": \"2017-08-09T14:30:35.0000000Z\",\n    \"price_open\": 3256.000000000,\n    \"price_high\": 3256.010000000,\n    \"price_low\": 3247.000000000,\n    \"price_close\": 3255.600000000,\n    \"volume_traded\": 58.131397920,\n    \"trades_count\": 33\n  }\n]"
	restyMock.On("Do", mock.Anything).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(sampleResponse)))}, nil)
	CalculationChan := make(chan Candle)
	testPolygon := &Polygon{
		API_KEY: "",
		Client:  restyMock,
	}

	params := Params{
		symbol:          "BTCUSD",
		period:          "1MIN",
		limit:           "3",
		CalculationChan: CalculationChan,
	}
	testPolygon.GetCandles(params)

}

func TestUnitToRFC(t *testing.T) {
	t.Run("Should convert unix ms to RFC format", func(t *testing.T) {
		testUnix := 1644373895991
		unixToRFC(int64(testUnix))
	})
}

func TestGetTimeFormatted(t *testing.T) {
	t.Run("Should return current time formatted YYYY-MM-DD", func(t *testing.T) {
		getTodaysDate()
	})
}

func TestReverseCandles(t *testing.T) {
	t.Run("Should reverse a slice of candles", func(t *testing.T) {
		sampleCandles := []Candle{
			{
				PriceHigh: 5,
			},
			{
				PriceHigh: 4,
			},
			{
				PriceHigh: 3,
			},
			{
				PriceHigh: 2,
			},
			{
				PriceHigh: 1,
			},
		}

		reverseCandles(sampleCandles)

		expectedFirstElement := 1.0

		assert.Equal(t, expectedFirstElement, sampleCandles[0].PriceHigh)
	})
}

func TestFormatSymbol(t *testing.T) {
	t.Run("Should insert slash in middle of string", func(t *testing.T) {
		symbol := "BTCUSD"
		formatted := formatSymbol(symbol)

		assert.Equal(t, "BTC/USD", formatted)
	})
}

func TestFormatUnit(t *testing.T) {
	t.Run("Should format time string to 3 digit string", func(t *testing.T) {
		result := formatUnit("1MIN")

		assert.Equal(t, "minute", result)
	})

}
