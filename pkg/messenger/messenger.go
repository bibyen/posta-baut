//go:generate mockgen -source=messenger.go -destination=mocks/messenger.go -package=mocks
package messenger

import (
	"context"
	"fmt"

	"github.com/bibyen/posta-baut/pkg/messenger/msgraph"
)

// Messenger defines methods for sending messages to chats and channels.
type Messenger interface {
	SendChannelMessage(
		ctx context.Context,
		teamID string,
		channelID string,
		msg string,
	) error
	SendChatMessage(
		ctx context.Context,
		chatID string,
		msg string,
	) error
}

func NewMessenger(cfg MessengerConfig) (Messenger, error) {
	switch cfg.Type {
	case MessengerTypeGraph:
		if cfg.GraphConfig == nil {
			return nil, fmt.Errorf("graph config required")
		}
		return msgraph.NewGraphMessenger(*cfg.GraphConfig)
	case MessengerTypeBot:
		// TODO: Write NewBotFrameworkMessenger(config)
		return nil, fmt.Errorf("unimplemented messenger type: %s", cfg.Type)
	default:
		return nil, fmt.Errorf("unknown messenger type: %s", cfg.Type)
	}
}
