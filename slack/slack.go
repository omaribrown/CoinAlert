package slack

import (
	"fmt"
	envVariables "github.com/omaribrown/coinalert/envvar"
	"github.com/slack-go/slack"
	"time"
)

func PostToSlack() {
	SlackToken := envVariables.ViperEnvVariable("SLACK_AUTH_TOKEN")
	SlackChannelID := envVariables.ViperEnvVariable("SLACK_CHANNEL_ID")

	testMessage := "DataData"

	client := slack.New(SlackToken, slack.OptionDebug(true))

	// Message
	attachment := slack.Attachment{
		Pretext: "Test Bot Message",
		Text:    testMessage,
		Color:   "red",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := client.PostMessage(
		SlackChannelID,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
	//SlackTest := "Hello, World"
	//fmt.Println(SlackToken, SlackChannelID)
}
