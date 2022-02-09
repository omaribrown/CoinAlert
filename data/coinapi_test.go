package coinapi

import (
	"bytes"
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
	CalculationChan := make(chan LatestOhlcv, 60)
	restyMock := new(httpMock)
	sampleResponse := "[\n  {\n    \"time_period_start\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:32:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:31:01.0000000Z\",\n    \"time_close\": \"2017-08-09T14:31:46.0000000Z\",\n    \"price_open\": 3255.590000000,\n    \"price_high\": 3255.590000000,\n    \"price_low\": 3244.740000000,\n    \"price_close\": 3244.740000000,\n    \"volume_traded\": 16.903274550,\n    \"trades_count\": 31\n  },\n  {\n    \"time_period_start\": \"2017-08-09T14:30:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:30:05.0000000Z\",\n    \"time_close\": \"2017-08-09T14:30:35.0000000Z\",\n    \"price_open\": 3256.000000000,\n    \"price_high\": 3256.010000000,\n    \"price_low\": 3247.000000000,\n    \"price_close\": 3255.600000000,\n    \"volume_traded\": 58.131397920,\n    \"trades_count\": 33\n  }\n]"
	restyMock.On("Do", mock.Anything).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(sampleResponse)))}, nil)

	// ACT
	testCoinapi := &Coinapi{
		API_KEY: "",
		Client:  restyMock,
	}
	testCoinapi.GetCoinLatest("SYM", "PERIOD", "1", CalculationChan)

	// ASSERT

}

func TestPolygon_GetCoinLatest(t *testing.T) {
	restyMock := new(httpMock)
	sampleResponse := "[\n  {\n    \"time_period_start\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:32:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:31:01.0000000Z\",\n    \"time_close\": \"2017-08-09T14:31:46.0000000Z\",\n    \"price_open\": 3255.590000000,\n    \"price_high\": 3255.590000000,\n    \"price_low\": 3244.740000000,\n    \"price_close\": 3244.740000000,\n    \"volume_traded\": 16.903274550,\n    \"trades_count\": 31\n  },\n  {\n    \"time_period_start\": \"2017-08-09T14:30:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:30:05.0000000Z\",\n    \"time_close\": \"2017-08-09T14:30:35.0000000Z\",\n    \"price_open\": 3256.000000000,\n    \"price_high\": 3256.010000000,\n    \"price_low\": 3247.000000000,\n    \"price_close\": 3255.600000000,\n    \"volume_traded\": 58.131397920,\n    \"trades_count\": 33\n  }\n]"
	restyMock.On("Do", mock.Anything).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(sampleResponse)))}, nil)

	testPolygon := &Polygon{
		API_KEY: "",
		Client:  restyMock,
	}

	testPolygon.GetCoinLatest("ETHUSD", "1", "minute", "3")

}

func TestUnitToRFC(t *testing.T) {
	t.Run("Should convert unix ms to RFC format", func(t *testing.T) {
		testUnix := 1644373895991
		unixToRFC(int64(testUnix))
	})
}

func TestGetTimeFormatted(t *testing.T) {
	t.Run("Should return current time formatted YYYY-MM-DD", func(t *testing.T) {
		getTimeFormatted()
	})
}
