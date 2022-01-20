package slack

import (
	"fmt"
	"github.com/slack-go/slack"
)

type ISlackService interface {
	SendSlackMessage(string)
}

type SlackService struct {
	SlackToken     string
	SlackChannelID string
}

type SlackMessage struct {
	Pretext string `json:"pretext,omitempty"`
	Text    string `json:"text,omitempty"`
}

func GenerateNewMessage(Pretext string, Text string) SlackMessage {
	NewMessage := SlackMessage{
		Pretext: Pretext,
		Text:    Text,
	}
	return NewMessage
}

func (s SlackService) SendSlackMessage(message SlackMessage) {
	client := slack.New(s.SlackToken, slack.OptionDebug(true))

	attachment := slack.Attachment{
		Pretext: message.Pretext,
		Text:    message.Text,
	}

	_, timestamp, err := client.PostMessage(
		s.SlackChannelID,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
}
