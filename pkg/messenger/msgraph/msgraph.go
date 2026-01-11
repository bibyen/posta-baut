package msgraph

import (
	"context"
	"fmt"

	config "github.com/bibyen/posta-baut/cmd/config"
	msgraph "github.com/bibyen/posta-baut/cmd/shared/msgraph"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GraphMessenger implements the Messenger interface using Microsoft Graph API.
type GraphMessenger struct {
	client msgraphsdk.GraphServiceClient
}

// SendChannelMessage sends a message to a specified channel in a team.
// Returns an error if the operation fails.
func (gm GraphMessenger) SendChannelMessage(ctx context.Context, teamID string, channelID string, msg string) error {
	requestBody := graphmodels.NewChatMessage()
	body := graphmodels.NewItemBody()
	content := msg
	body.SetContent(&content)
	requestBody.SetBody(body)

	_, err := gm.client.Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		Post(context.Background(), requestBody, nil)

	if err != nil {
		return fmt.Errorf(
			"failed to send channel message to channel id %s belonging to team id %s: %w",
			channelID,
			teamID,
			err,
		)
	}
	return nil
}

// SendChatMessage sends a message to a specified chat.
// Returns an error if the operation fails.
func (gm GraphMessenger) SendChatMessage(ctx context.Context, chatID string, msg string) error {
	requestBody := graphmodels.NewChatMessage()
	body := graphmodels.NewItemBody()
	content := msg
	body.SetContent(&content)
	requestBody.SetBody(body)

	_, err := gm.client.Chats().
		ByChatId(chatID).
		Messages().
		Post(context.Background(), requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to send chat message to chat-id \"%s\": %w", chatID, err)
	}
	return nil
}

// NewGraphMessenger creates a new GraphMessenger with the provided configuration.
func NewGraphMessenger(cfg config.ClientConfig) (GraphMessenger, error) {
	msGraphClient, err := msgraph.NewMSGraphClient(cfg.TenantID, cfg.ClientID, cfg.ClientSecret)
	if err != nil {
		return GraphMessenger{}, fmt.Errorf("failed to create MS Graph Messenger: %w", err)
	}

	return GraphMessenger{
		client: *msGraphClient,
	}, nil
}
