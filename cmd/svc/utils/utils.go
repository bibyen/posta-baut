// Package utils provides service-level utilities for the Posta Baut application.
package utils

import (
	"fmt"

	pbv1 "github.com/bibyen/posta-baut/internal/pb/v1"
	"github.com/bibyen/posta-baut/pkg/messenger"
)

// ReqToMsg converts the protobuf SendMessageRequest into the domain SendMessageRequest.
func ReqToMsg(pbReq *pbv1.SendMessageRequest) (messenger.SendMessageRequest, error) {
	if pbReq == nil {
		return messenger.SendMessageRequest{}, fmt.Errorf("request cannot be nil")
	}

	if pbReq.Target == nil {
		return messenger.SendMessageRequest{}, fmt.Errorf("target is required")
	}

	var target messenger.MessageTarget

	switch t := pbReq.Target.Target.(type) {
	case *pbv1.MessageTarget_Channel:
		if t.Channel == nil {
			return messenger.SendMessageRequest{}, fmt.Errorf("channel target is nil")
		}
		target.Channel = &messenger.TeamsChannelTarget{
			TeamID:    t.Channel.GetTeamId(),
			ChannelID: t.Channel.GetChannelId(),
			ThreadID:  t.Channel.GetThreadId(),
		}
	case *pbv1.MessageTarget_Chat:
		if t.Chat == nil {
			return messenger.SendMessageRequest{}, fmt.Errorf("chat target is nil")
		}
		target.Chat = &messenger.ChatTarget{
			ChatID:           t.Chat.GetChatId(),
			ReplyToMessageID: t.Chat.GetReplyToMessageId(),
		}
	default:
		return messenger.SendMessageRequest{}, fmt.Errorf("unknown message target type: %T", t)
	}

	if pbReq.Content == "" {
		return messenger.SendMessageRequest{}, fmt.Errorf("content cannot be empty")
	}

	domainReq := messenger.SendMessageRequest{
		Target:  target,
		Content: pbReq.GetContent(),
	}

	return domainReq, nil
}
