package msgraph

import (
	"fmt"

	gh "github.com/bibyen/posta-baut/cmd/shared/msgraph/graphhelper"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// NewMSGraphClient returns a GraphServiceClient with app permissions.
func NewMSGraphClient(tenantID string, clientID string, clientSecret string) (*msgraphsdk.GraphServiceClient, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant ID is required")
	}
	if clientID == "" {
		return nil, fmt.Errorf("client ID is required")
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("clientSecret is required")
	}

	graphServiceClient, err := gh.NewGraphHelper().NewGraphForAppAuth(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("error creating graph service client with auth: %w", err)
	}

	return graphServiceClient, nil
}
