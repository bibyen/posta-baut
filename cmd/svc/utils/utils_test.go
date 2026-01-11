// Package utils provides service-level utilities for the Posta Baut application.
package utils

import (
	"testing"

	pbv1 "github.com/bibyen/posta-baut/internal/pb/v1"
	"github.com/bibyen/posta-baut/pkg/messenger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	validChannelID  = "channel-123"
	validTeamID     = "team-123"
	validThreadID   = "thread-123"
	validMsgContent = "Hello!"
)

var (
	pbReqChannelReply = &pbv1.SendMessageRequest{
		Target: &pbv1.MessageTarget{
			Target: &pbv1.MessageTarget_Channel{
				Channel: &pbv1.TeamsChannelTarget{
					ChannelId: validChannelID,
					TeamId:    validTeamID,
					ThreadId:  validThreadID,
				},
			},
		},
		Content: validMsgContent,
	}
)

func TestReqToMsg(t *testing.T) {
	tests := []struct {
		name       string
		pbReq      *pbv1.SendMessageRequest
		want       *messenger.SendMessageRequest
		wantErrMsg string
	}{
		{
			name:  "success",
			pbReq: pbReqChannelReply,
			want: &messenger.SendMessageRequest{
				Target: messenger.MessageTarget{
					Channel: &messenger.TeamsChannelTarget{
						ChannelID: validChannelID,
						TeamID:    validTeamID,
						ThreadID:  validThreadID,
					},
				},
				Content: validMsgContent,
			},
		},
		{
			name:       "nil request",
			pbReq:      nil,
			wantErrMsg: "request cannot be nil",
		},
		{
			name: "message target is nil",
			pbReq: &pbv1.SendMessageRequest{
				Target: &pbv1.MessageTarget{
					Target: nil,
				},
			},
			wantErrMsg: "unknown message target type:",
		},
		{
			name: "nil target channel",
			pbReq: &pbv1.SendMessageRequest{
				Target: &pbv1.MessageTarget{
					Target: &pbv1.MessageTarget_Channel{
						Channel: nil,
					},
				},
			},
			wantErrMsg: "channel target is nil",
		},
		{
			name: "nil target chat",
			pbReq: &pbv1.SendMessageRequest{
				Target: &pbv1.MessageTarget{
					Target: &pbv1.MessageTarget_Chat{
						Chat: nil,
					},
				},
			},
			wantErrMsg: "chat target is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ReqToMsg(tt.pbReq)
			if tt.wantErrMsg != "" {
				require.ErrorContains(t, gotErr, tt.wantErrMsg)
			} else {
				assert.Equal(t, tt.want, &got)
				require.NoError(t, gotErr)
			}
		})
	}
}
