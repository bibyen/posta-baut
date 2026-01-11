package client

import (
	"fmt"

	config "github.com/bibyen/posta-baut/cmd/config"
	lookup "github.com/bibyen/posta-baut/cmd/shared/lookup"
	msgraph "github.com/bibyen/posta-baut/pkg/messenger/msgraph"
)

// NewClient creates a new Client with the provided Messenger and LookupClient configurations.
func NewClient(messengerConfig config.ClientConfig, lookupClientConfig config.ClientConfig) (Client, error) {
	messenger, err := msgraph.NewGraphMessenger(messengerConfig)
	if err != nil {
		return Client{}, fmt.Errorf("failed to create messenger: %w", err)
	}

	lookupClient, err := lookup.NewMSGraphLookupClient(lookupClientConfig)
	if err != nil {
		return Client{}, fmt.Errorf("failed to create lookup client: %w", err)
	}

	return Client{
		Messenger:    messenger,
		LookupClient: lookupClient,
	}, nil
}
