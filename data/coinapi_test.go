package coinapi

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/resty.v0"
	"testing"
)

type RestyMock struct {
	mock.Mock
}

func (_m *RestyMock) Get(url string) (*resty.Response, error) {
	ret := _m.Called(url)

	return ret.Get(0).(*resty.Response), ret.Error(1)
}

func (_m *RestyMock) SetHeader(header string, value string) *resty.Request {
	ret := _m.Called(header, value)

	return ret.Get(0).(*resty.Request)
}

func TestCoinapi_GetCoinLatest(t *testing.T) {
	// ARRANGE
	restyMock := new(RestyMock)
	sampleResponse := "[\n  {\n    \"time_period_start\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:32:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:31:01.0000000Z\",\n    \"time_close\": \"2017-08-09T14:31:46.0000000Z\",\n    \"price_open\": 3255.590000000,\n    \"price_high\": 3255.590000000,\n    \"price_low\": 3244.740000000,\n    \"price_close\": 3244.740000000,\n    \"volume_traded\": 16.903274550,\n    \"trades_count\": 31\n  },\n  {\n    \"time_period_start\": \"2017-08-09T14:30:00.0000000Z\",\n    \"time_period_end\": \"2017-08-09T14:31:00.0000000Z\",\n    \"time_open\": \"2017-08-09T14:30:05.0000000Z\",\n    \"time_close\": \"2017-08-09T14:30:35.0000000Z\",\n    \"price_open\": 3256.000000000,\n    \"price_high\": 3256.010000000,\n    \"price_low\": 3247.000000000,\n    \"price_close\": 3255.600000000,\n    \"volume_traded\": 58.131397920,\n    \"trades_count\": 33\n  }\n]"
	restyMock.On("Get", mock.Anything).Return(&resty.Response{Body: []byte(sampleResponse)})

	// ACT
	testCoinapi := &Coinapi{
		API_KEY: "",
		Resty: restyMock,
	}
	testCoinapi.GetCoinLatest("SYM", "PERIOD", "1")

	// ASSERT

}
