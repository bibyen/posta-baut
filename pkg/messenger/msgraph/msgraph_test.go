package msgraph

import (
	"context"
	"reflect"
	"testing"

	config "github.com/bibyen/posta-baut/cmd/config"
)

func TestGraphMessenger_SendChannelMessage(t *testing.T) {
	type args struct {
		ctx       context.Context
		teamID    string
		channelID string
		msg       string
	}
	tests := []struct {
		name    string
		gm      GraphMessenger
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gm.SendChannelMessage(tt.args.ctx, tt.args.teamID, tt.args.channelID, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("GraphMessenger.SendChannelMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGraphMessenger_SendChatMessage(t *testing.T) {
	type args struct {
		ctx    context.Context
		chatID string
		msg    string
	}
	tests := []struct {
		name    string
		gm      GraphMessenger
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gm.SendChatMessage(tt.args.ctx, tt.args.chatID, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("GraphMessenger.SendChatMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewGraphMessenger(t *testing.T) {
	type args struct {
		cfg config.ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    GraphMessenger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGraphMessenger(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewGraphMessenger() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraphMessenger() = %v, want %v", got, tt.want)
			}
		})
	}
}
