package triggers

//func TestSendSignal(t *testing.T) {
//	t.Run("Should send message to slack", func(t *testing.T) {
//		testBollBand := make(chan coinapi.LatestOhlcv, 20)
//		testBollBand <- coinapi.LatestOhlcv{
//			PriceClose:         100,
//			BollingerBandUpper: 95,
//			BollingerBandLower: 70,
//		}
//		testMessage := make(chan slack.SlackMessage, 20)
//		testMessage <- slack.SlackMessage{
//			Title:   "Did this shit work?",
//			Pretext: "fuck",
//			Text:    "yeah",
//		}
//
//		testSend := new(SlackTrigger)
//
//		testSend.sendSignal(testBollBand, testMessage)
//
//		assert.Equal(t, "pre", testMessage)
//	})
//}
