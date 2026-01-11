// Package svc provides service-level utilities for the Posta Baut application.
package svc

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/bibyen/posta-baut/cmd/client"
	pbv1 "github.com/bibyen/posta-baut/internal/pb/v1"
	mocks "github.com/bibyen/posta-baut/pkg/messenger/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
)

const (
	validChatID     = "chat-id"
	validMsgContent = "Hello!"
)

var (
	validReq = connect.NewRequest(
		&pbv1.SendMessageRequest{
			Content: validMsgContent,
			Target: &pbv1.MessageTarget{
				Target: &pbv1.MessageTarget_Chat{
					Chat: &pbv1.TeamsChatTarget{
						ChatId: validChatID,
					},
				},
			},
		},
	)
)

func TestNewTeamsServiceClient(t *testing.T) {
	tsClient := NewTeamsService(&client.Client{})
	require.Equal(t, tsClient, &teamsService{
		Client: &client.Client{},
	})
}

func Test_teamsService_SendMessage_Failure(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[pbv1.SendMessageRequest]
	}
	tests := []struct {
		name       string
		s          *teamsService
		args       args
		want       *connect.Response[pbv1.SendMessageResponse]
		wantErrMsg string
	}{
		{
			name: "fail to convert request to message",
			s:    NewTeamsService(&client.Client{}),
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&pbv1.SendMessageRequest{}),
			},
			wantErrMsg: "failed to convert request to message",
			want:       nil,
		},
		{
			name: "fail to send message",
			s: NewTeamsService(&client.Client{
				Messenger: nil,
			}),
			args: args{
				ctx: context.Background(),
				req: validReq,
			},
			wantErrMsg: "failed to send message",
			want:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SendMessage(tt.args.ctx, tt.args.req)
			if tt.wantErrMsg != "" {
				require.ErrorContainsf(t, err, tt.wantErrMsg, "teamsService.SendMessage() error = %v, err should contain %v", err, tt.wantErrMsg)
			} else {
				require.NoErrorf(t, err, "teamsService.SendMessage() error = %v, wantErr %v", err, tt.wantErrMsg)
			}

			require.Equal(t, got, tt.want, "teamsService.SendMessage() = %v, want %v", got, tt.want)
		})
	}
}

func Test_teamsService_SendMessage_Success(t *testing.T) {
	t.Run("successfully send message", func(t *testing.T) {
		// Setup mock messenger
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msgr := mocks.NewMockMessenger(ctrl)
		msgr.EXPECT().SendChatMessage(gomock.Any(), validChatID, validMsgContent).Return(nil).Times(1)

		teamService := NewTeamsService(&client.Client{
			Messenger: msgr,
		})

		want := &pbv1.SendMessageResponse{
			MessageId: "success", // TODO: Return idempotent messageid in response https://github.com/bibyen/posta-baut/issues/15
		}

		got, err := teamService.SendMessage(context.Background(), validReq)
		require.NoError(t, err)
		// assert.Equal(t, got.Msg.MessageId, want.MessageId)
		proto.Equal(got.Msg, want)
	})
}
