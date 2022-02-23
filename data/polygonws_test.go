package coinapi

import "testing"

func TestOpenWebsocket(t *testing.T) {
	t.Run("Should start logging", func(t *testing.T) {
		openWebsocket("")
	})
}
