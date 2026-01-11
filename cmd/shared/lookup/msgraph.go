package lookup

import (
	"context"
	"fmt"
	"log"

	config "github.com/bibyen/posta-baut/cmd/config"
	msgraph "github.com/bibyen/posta-baut/cmd/shared/msgraph"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

type MSGraphLookup struct {
	// graph client
	client msgraphsdk.GraphServiceClient
}

// NewMSGraphLookupClient creates a new MSGraphLookup instance.
func NewMSGraphLookupClient(cfg config.ClientConfig) (*MSGraphLookup, error) {
	NewMSGraphClient, err := msgraph.NewMSGraphClient(cfg.TenantID, cfg.ClientID, cfg.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create MS Graph Client: %w", err)
	}
	client := NewMSGraphClient
	return &MSGraphLookup{
		client: *client,
	}, nil
}

// TeamNameFromID looks up the team name given its ID.
func (l *MSGraphLookup) TeamNameFromID(teamID string) (string, error) {
	teamName := ""
	team, err := l.client.Teams().ByTeamId(teamID).Get(context.Background(), nil)
	if err != nil {
		log.Printf("failed to get team for team ID %s: %v", teamID, err)
		return "", err
	}
	teamName = *team.GetDisplayName()

	return teamName, nil
}

// ChannelNameFromID looks up the channel name given its team ID and channel ID.
func (l *MSGraphLookup) ChannelNameFromID(teamID string, channelID string) (string, error) {
	channelName := ""
	channel, err := l.client.
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Get(context.Background(), nil)
	if err != nil {
		log.Printf("failed to get channel for channel ID %s: %v", channelID, err)
		return "", err
	}
	channelName = *channel.GetDisplayName()

	return channelName, nil
}
