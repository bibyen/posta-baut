package client

import (
	"testing"

	config "github.com/bibyen/posta-baut/cmd/config"
	"github.com/stretchr/testify/require"
)

const (
	validTenantID     = "tenant-123"
	validClientID     = "client-123"
	validClientSecret = "client-secret-123"
)

func TestNewClient(t *testing.T) {
	type args struct {
		messengerConfig    config.ClientConfig
		lookupClientConfig config.ClientConfig
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		// Success cases
		{
			name: "valid messenger config",
			args: args{
				messengerConfig: config.ClientConfig{
					TenantID:     validTenantID,
					ClientID:     validClientID,
					ClientSecret: validClientSecret,
				},
				lookupClientConfig: config.ClientConfig{
					TenantID:     validTenantID,
					ClientID:     validClientID,
					ClientSecret: validClientSecret,
				},
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		// Failure cases
		{
			name: "invalid messenger config",
			args: args{
				messengerConfig: config.ClientConfig{
					TenantID: "",
					ClientID: validClientID,
				},
				lookupClientConfig: config.ClientConfig{
					TenantID:     validTenantID,
					ClientID:     validClientID,
					ClientSecret: validClientSecret,
				},
			},
			wantErr:    true,
			wantErrMsg: "failed to create messenger",
		},
		{
			name: "invalid lookup client config",
			args: args{
				messengerConfig: config.ClientConfig{
					TenantID:     validTenantID,
					ClientID:     validClientID,
					ClientSecret: validClientSecret,
				},
				lookupClientConfig: config.ClientConfig{
					TenantID: "",
					ClientID: validClientID,
				},
			},
			wantErr:    true,
			wantErrMsg: "failed to create lookup client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.messengerConfig, tt.args.lookupClientConfig)
			if tt.wantErr {
				require.ErrorContainsf(t, err, tt.wantErrMsg, "NewClient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				require.NoErrorf(t, err, "NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
