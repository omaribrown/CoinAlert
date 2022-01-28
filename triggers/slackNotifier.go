package triggers

import coinapi "github.com/omaribrown/coinalert/data"

func slackNotifier(triggerChan chan coinapi.LatestOhlcv) {
	slackNotifierChan := make(chan coinapi.LatestOhlcv)
	for {

		slackNotifierChan <- triggerChan
	}
}
