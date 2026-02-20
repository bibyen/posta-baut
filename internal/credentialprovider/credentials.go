package credentialprovider

import "context"

type CredentialProvider interface {
	// Returns messenger credentials for a given userID and targetID
	LookupCredentials(ctx context.Context, userID, targetID string) (MessengerCredentials, error)
}

// BotCredentials holds the authentication credentials
// required for the Bot Framework messenger implementation.
type BotCredentials struct {
	AppID       string
	AppPassword string
	TenantID    string
}

// GraphCredentials holds the authentication credentials
// required for the Microsoft Graph API messenger implementation.
type GraphCredentials struct {
	ClientID     string
	ClientSecret string
	TenantID     string
}

// MessengerCredentials represents the union of possible
// credential types needed to authenticate with messaging services.
// Exactly one of BotCreds or GraphCreds should be non-nil.
type MessengerCredentials struct {
	BotCreds   *BotCredentials
	GraphCreds *GraphCredentials
}
