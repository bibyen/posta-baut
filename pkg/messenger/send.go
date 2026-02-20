package messenger

import (
	"context"
	"fmt"
)

func sendToChannel(ctx context.Context, m Messenger, msg SendMessageRequest) error {
	switch {
	case m == nil:
		return fmt.Errorf("messenger cannot be nil")
	case msg.Content == "":
		return fmt.Errorf("message content is empty")
	case msg.Target.Channel == nil:
		return fmt.Errorf("target channel is nil")
	case msg.Target.Channel.ChannelID == "":
		return fmt.Errorf("target channel ID missing")
	case msg.Target.Channel.TeamID == "":
		return fmt.Errorf("target team ID missing")

	default:
		err := m.SendChannelMessage(ctx, msg.Target.Channel.TeamID, msg.Target.Channel.ChannelID, msg.Content)
		if err != nil {
			return fmt.Errorf("error sending channel message: %w", err)
		}
		return nil
	}
}

func sendToChat(ctx context.Context, m Messenger, msg SendMessageRequest) error {
	switch {
	case m == nil:
		return fmt.Errorf("messenger cannot be nil")
	case msg.Content == "":
		return fmt.Errorf("message content is empty")
	case msg.Target.Chat == nil:
		return fmt.Errorf("target channel is nil")
	case msg.Target.Chat.ChatID == "":
		return fmt.Errorf("target channel ID missing")

	default:
		err := m.SendChatMessage(ctx, msg.Target.Chat.ChatID, msg.Content)
		if err != nil {
			return fmt.Errorf("error sending chat message: %w", err)
		}
		return nil
	}
}

func Send(ctx context.Context, mCfg MessengerConfig, msg SendMessageRequest) error {
	m, err := NewMessenger(mCfg)
	if err != nil {
		return fmt.Errorf("error creating messenger: %w", err)
	}

	switch {
	case msg.Target.Channel != nil:
		return sendToChannel(ctx, m, msg)

	case msg.Target.Chat != nil:
		return sendToChat(ctx, m, msg)

	default:
		return fmt.Errorf("message target must be specified")
	}
}
