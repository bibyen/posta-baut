package messenger

import (
	config "github.com/bibyen/posta-baut/cmd/config"
)

type MessengerType string

const (
	MessengerTypeGraph MessengerType = "graph"
	MessengerTypeBot   MessengerType = "bot"
)

type MessengerConfig struct {
	Type MessengerType // "graph" or "bot"

	GraphConfig *config.ClientConfig
	// TODO: Add BotConfig type
}
