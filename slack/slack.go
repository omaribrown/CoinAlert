package slack

import (
	"context"
	"errors"
	"fmt"
	envVariables "github.com/omaribrown/coinalert/envvar"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
	"strings"
	"time"
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
	s.SlackToken = envVariables.ViperEnvVariable("SLACK_AUTH_TOKEN")
	s.SlackChannelID = envVariables.ViperEnvVariable("SLACK_CHANNEL_ID")

	client := slack.New(s.SlackToken, slack.OptionDebug(true))

	// Message
	//message.Pretext = "pretext"
	//message.Text = "Text"
	pretext := message.Pretext
	fmt.Printf(pretext)
	text := message.Text
	fmt.Println(text)

	attachment := slack.Attachment{
		Pretext: pretext,
		Text:    text,
	}

	//attachment := slack.Attachment{
	//	Pretext: "Test Bot Message",
	//	Text:    s2,
	//	Color:   "#36a64f",
	//	//Fields: []slack.AttachmentField{
	//	//	{
	//	//		Title: "Date",
	//	//		Value: time.Now().String(),
	//	//	},
	//	//},
	//}

	_, timestamp, err := client.PostMessage(
		s.SlackChannelID,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
	//SlackTest := "Hello, World"
	//fmt.Println(SlackToken, SlackChannelID)
}

func PostToSlack() {
	SlackToken := envVariables.ViperEnvVariable("SLACK_AUTH_TOKEN")
	SlackChannelID := envVariables.ViperEnvVariable("SLACK_CHANNEL_ID")

	testMessage := "test"

	client := slack.New(SlackToken, slack.OptionDebug(true))

	// Message
	attachment := slack.Attachment{
		Pretext: "Test Bot Message",
		Text:    testMessage,
		Color:   "#36a64f",
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

func SlackSocket() {
	SlackToken := envVariables.ViperEnvVariable("SLACK_AUTH_TOKEN")
	AppToken := envVariables.ViperEnvVariable("SLACK_APP_TOKEN")

	client := slack.New(SlackToken, slack.OptionDebug(true), slack.OptionAppLevelToken(AppToken))
	// go-slack comes with a SocketMode package that we need to use that accepts a Slack client and outputs a Socket mode client instead
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		// Option to set a custom logger
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	// Create a context that can be used to cancel goroutine
	ctx, cancel := context.WithCancel(context.Background())
	// Make this cancel called properly in a real program , graceful shutdown etc
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		// Create a for loop that selects either the context cancellation or the events incomming
		for {
			select {
			// inscase context cancel is called exit the goroutine
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				// We have a new Events, let's type switch the event
				// Add more use cases here if you want to listen to other events.
				switch event.Type {
				// handle EventAPI events
				case socketmode.EventTypeEventsAPI:
					// The Event sent on the channel is not the same as the EventAPI events so we need to type cast it
					eventsAPIEvent, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
						continue
					}
					// We need to send an Acknowledge to the slack server
					socketClient.Ack(*event.Request)
					// Now we have an Events API event, but this event type can in turn be many types, so we actually need another type switch
					err := handleEventMessage(eventsAPIEvent, client)
					if err != nil {
						// Replace with actual err handeling
						log.Fatal(err)
					}

					// handle slash command events
				case socketmode.EventTypeSlashCommand:
					command, ok := event.Data.(slack.SlashCommand)
					if !ok {
						log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)
						continue
					}

					payload, err := handleSlashCommand(command, client)
					if err != nil {
						log.Fatal(err)
					}
					socketClient.Ack(*event.Request, payload)
				// handle interaction events
				case socketmode.EventTypeInteractive:
					interaction, ok := event.Data.(slack.InteractionCallback)
					if !ok {
						log.Printf("Could not type cast the message to an Interaction callback: %v\n", interaction)
						continue
					}
					err := handleInteractionEvent(interaction, client)
					if err != nil {
						log.Fatal(err)
					}
					socketClient.Ack(*event.Request)
				}

			}
		}
	}(ctx, client, socketClient)

	socketClient.Run()
}

func handleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := handleAppMentionEvent(ev, client)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

func handleAppMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {

	user, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}
	text := strings.ToLower(event.Text)

	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: user.Name,
		},
	}
	if strings.Contains(text, "hello") {
		attachment.Text = fmt.Sprintf("Hello %s", user.Name)
		attachment.Pretext = "Greetings"
		attachment.Color = "#4af030"
	} else {
		attachment.Text = fmt.Sprintf("How can I help you %s?", user.Name)
		attachment.Pretext = "How can I be of service"
		attachment.Color = "#3d3d3d"
	}
	_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}
	return nil
}

// handles response to questionnaire
func handleInteractionEvent(interaction slack.InteractionCallback, client *slack.Client) error {
	log.Printf("The action called is: %s\n", interaction.ActionID)
	log.Printf("The response was of type: %s\n", interaction.Type)

	switch interaction.Type {
	case slack.InteractionTypeBlockActions:
		for _, action := range interaction.ActionCallback.BlockActions {
			log.Printf("%+v", action)
			//log.Printf("Selected option: ", action.SelectedOption)
		}
	default:

	}
	return nil
}

// Slash command parent
func handleSlashCommand(command slack.SlashCommand, client *slack.Client) (interface{}, error) {
	switch command.Command {
	case "/hello":
		return nil, handleHelloCommand(command, client)
	case "/was-this-article-useful":
		return handleArticleGood(command, client)
	}
	return nil, nil
}

// Slash Command children
func handleHelloCommand(command slack.SlashCommand, client *slack.Client) error {
	attachment := slack.Attachment{}

	attachment.Fields = []slack.AttachmentField{
		{
			Title: "Date",
			Value: time.Now().String(),
		}, {
			Title: "Initializer",
			Value: command.UserName,
		}, {
			Title: "Function:",
			Value: "handleHelloCommand",
		},
	}
	attachment.Text = fmt.Sprintf("Hello %s", command.Text)
	attachment.Color = "#4af030"

	_, _, err := client.PostMessage(command.ChannelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// // Creates questionnaire
func handleArticleGood(command slack.SlashCommand, client *slack.Client) (interface{}, error) {
	attachment := slack.Attachment{}

	checkbox := slack.NewCheckboxGroupsBlockElement("answer",
		slack.NewOptionBlockObject("yes", &slack.TextBlockObject{Text: "Yes", Type: slack.MarkdownType}, &slack.TextBlockObject{Text: "Did you enjoy it?", Type: slack.MarkdownType}),
		slack.NewOptionBlockObject("no", &slack.TextBlockObject{Text: "No", Type: slack.MarkdownType}, &slack.TextBlockObject{Text: "Did you dislike it?", Type: slack.MarkdownType}),
	)

	accessory := slack.NewAccessory(checkbox)

	attachment.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewSectionBlock(
				&slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "Did you think this article was helpful?",
				},
				nil,
				accessory,
			),
		},
	}
	attachment.Text = "Rate the tutorial"
	attachment.Color = "#4af030"
	return attachment, nil
}
