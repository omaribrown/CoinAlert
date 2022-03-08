package slack

import (
	"testing"
)

func TestSlackService(t *testing.T) {
	t.Run("should send real message to slack", func(t *testing.T) {
		s := SlackService{
			SlackToken:     "",
			SlackChannelID: "",
		}
		message := GenerateNewMessage("New Message", "Lets get these bands")
		s.SendSlackMessage(message)
	})
}
