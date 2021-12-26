package slack

import (
	"testing"
)

func TestSlackFrame_SendSlackMessage(t *testing.T) {
	type fields struct {
		SlackToken     string
		SlackChannelID string
	}
	type args struct {
		message SlackMessage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SlackKeys{
				SlackToken:     tt.fields.SlackToken,
				SlackChannelID: tt.fields.SlackChannelID,
			}
			s.SendSlackMessage(tt.args.message)
		})
	}
}
