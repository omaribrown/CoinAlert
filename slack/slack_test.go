package slack

import (
	"reflect"
	"testing"
)

func TestGenerateNewMessage(t *testing.T) {
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
			s := SlackKeys{
				SlackToken:     tt.fields.SlackToken,
				SlackChannelID: tt.fields.SlackChannelID,
			}
			s.SendSlackMessage(tt.args.message)
		})
	}
}
