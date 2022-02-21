package slack

import (
	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

type ISlackService interface {
	SendSlackMessage(string)
}

type SlackService struct {
	SlackToken     string
	SlackChannelID string
}

type SlackMessage struct {
	Title   string `json:"title,omitempty"`
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
		Title:   message.Title,
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
	zap.S().Infof("Message sent at %s", timestamp)
}
