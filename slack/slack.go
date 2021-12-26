package slack

import (
	"fmt"
	"github.com/slack-go/slack"
)

type SlackService interface {
	SendSlackMessage(string)
}

type SlackKeys struct {
	SlackToken     string
	SlackChannelID string
}

type SlackMessage struct {
	Pretext string `json:"pretext,omitempty"`
	Text    string `json:"text,omitempty"`
}

func GenerateNewMessage(s1 string, s2 string) *SlackMessage {
	NewMessage := &SlackMessage{
		Pretext: s1,
		Text:    s2,
	}
	return NewMessage
}

func (s SlackKeys) SendSlackMessage(message SlackMessage) {

	//s.SlackChannelID = envVariables.ViperEnvVariable("SLACK_CHANNEL_ID")

	client := slack.New(s.SlackToken, slack.OptionDebug(true))

	pretext := message.Pretext
	text := message.Text

	attachment := slack.Attachment{
		Pretext: pretext,
		Text:    text,
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
