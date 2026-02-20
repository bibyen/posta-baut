package credentialprovider

import (
	"context"
	"fmt"

	"github.com/bibyen/posta-baut/pkg/messenger"
)

type InMemoryCredentialProvider struct {
	credStore map[string]MessengerCredentials
}

const (
	dummyId = "hello"
)

var InMemCredentialProvider = InMemoryCredentialProvider{
	credStore: map[string]MessengerCredentials{
		"GraphUserAChannelB": {
			GraphCreds: &GraphCredentials{
				ClientID:     dummyId,
				TenantID:     dummyId,
				ClientSecret: dummyId,
			},
		},
		"BotUserCChatD": {
			BotCreds: &BotCredentials{
				AppID:       dummyId,
				AppPassword: dummyId,
				TenantID:    dummyId,
			},
		},
	},
}

func (imProvider *InMemoryCredentialProvider) LookupCredentials(ctx context.Context, userId string, target messenger.MessageTarget) *MessengerCredentials {
	key := ""
	if target.Channel != nil {
		key = fmt.Sprintf("%s%s", userId, target.Channel.ChannelID)
	} else if target.Chat != nil {
		key = fmt.Sprintf("%s%s", userId, target.Chat.ChatID)
	}

	if val, ok := InMemCredentialProvider.credStore[key]; ok {
		return &val
	}
	return nil
}
