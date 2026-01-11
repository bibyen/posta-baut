package messenger_test

import (
	"testing"

	"github.com/bibyen/posta-baut/pkg/messenger"
)

const (
	validChatID    = "chat-123"
	validMessageID = "messg-123"
	validTeamID    = "team-123"
	validChannelID = "channel-123"
	validContent   = "Hello, World!"
)

var (
	testChannelMessage = messenger.SendMessageRequest{
		Target: messenger.MessageTarget{
			Channel: &messenger.TeamsChannelTarget{
				ChannelID: validChannelID,
				TeamID:    validTeamID,
				ThreadID:  validContent,
			},
		},
		Content: validContent,
	}

	testChannelReply = messenger.SendMessageRequest{
		Target: messenger.MessageTarget{
			Channel: &messenger.TeamsChannelTarget{
				ChannelID: validChannelID,
				TeamID:    validTeamID,
			},
		},
		Content: validContent,
	}

	testChatMessage = messenger.SendMessageRequest{
		Target: messenger.MessageTarget{
			Chat: &messenger.ChatTarget{
				ChatID: validChatID,
			},
		},
		Content: validContent,
	}

	testChatReply = messenger.SendMessageRequest{
		Target: messenger.MessageTarget{
			Chat: &messenger.ChatTarget{
				ChatID:           validChatID,
				ReplyToMessageID: validMessageID,
			},
		},
		Content: validContent,
	}
)

func TestSend_Fail(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		m          messenger.Messenger
		msg        messenger.SendMessageRequest
		wantErr    bool
		wantErrMsg string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSend_Success(t *testing.T) {
	tests := []struct {
		name string
		msg  messenger.SendMessageRequest
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
