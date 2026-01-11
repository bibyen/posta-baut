package lookup

import (
	"reflect"
	"testing"

	config "github.com/bibyen/posta-baut/cmd/config"
)

func TestNewMSGraphLookupClient(t *testing.T) {
	type args struct {
		cfg config.ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *MSGraphLookup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMSGraphLookupClient(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMSGraphLookupClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMSGraphLookupClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSGraphLookup_TeamNameFromID(t *testing.T) {
	type args struct {
		teamID string
	}
	tests := []struct {
		name    string
		l       *MSGraphLookup
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.TeamNameFromID(tt.args.teamID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("MSGraphLookup.TeamNameFromID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("MSGraphLookup.TeamNameFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSGraphLookup_ChannelNameFromID(t *testing.T) {
	type args struct {
		teamID    string
		channelID string
	}
	tests := []struct {
		name    string
		l       *MSGraphLookup
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.ChannelNameFromID(tt.args.teamID, tt.args.channelID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("MSGraphLookup.ChannelNameFromID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("MSGraphLookup.ChannelNameFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}
