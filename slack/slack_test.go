package slack

import (
	"testing"
)

func TestSlackService (t *testing.T) {
	t.Run("should send real message to slack", func(t *testing.T) {
		s := SlackService{
			SlackToken:     "xoxb-2865687350037-2865731248501-uP1yg8TYTPFLG0HUmKfnTADf",
			SlackChannelID: "C02RJK5EXD1",
		}
		message := GenerateNewMessage("New Message", "Lets get these bands")
		s.SendSlackMessage(message)
	})
}
/*func TestGenerateNewMessage(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want *SlackMessage
	}{
		{
			name: "normal",
			args: args{
				"one",
				"two",
			},
			want: &SlackMessage{
				"one",
				"two",
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNewMessage(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateNewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlackKeys_SendSlackMessage(t *testing.T) {
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
			s := SlackService{
				SlackToken:     tt.fields.SlackToken,
				SlackChannelID: tt.fields.SlackChannelID,
			}
			s.SendSlackMessage(tt.args.message)
		})
	}
}*/
