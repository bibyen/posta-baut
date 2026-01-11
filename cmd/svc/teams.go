// Package svc provides service-level utilities for the Posta Baut application.
package svc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/bibyen/posta-baut/cmd/client"
	utils "github.com/bibyen/posta-baut/cmd/svc/utils"
	pb "github.com/bibyen/posta-baut/internal/pb/v1"
	"github.com/bibyen/posta-baut/internal/pb/v1/pbv1connect"
	"github.com/bibyen/posta-baut/pkg/messenger"
)

// teamsService implements the TeamsService defined in the protobuf.
type teamsService struct {
	pbv1connect.UnimplementedTeamsServiceHandler
	Client *client.Client
}

// NewTeamsService creates a new TeamsService with the provided client.
func NewTeamsService(client *client.Client) *teamsService {
	return &teamsService{
		Client: client,
	}
}

// SendMessage handles sending messages to Microsoft Teams.
func (s *teamsService) SendMessage(ctx context.Context, req *connect.Request[pb.SendMessageRequest]) (*connect.Response[pb.SendMessageResponse], error) {
	// Convert the request to domain-level Message - messenger.Message
	msg, err := utils.ReqToMsg(req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request to message: %w", err)
	}

	// Use the client's Messenger to send the message
	err = messenger.Send(ctx, s.Client.Messenger, msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to send message: %w", err))
	}

	resp := &pb.SendMessageResponse{
		MessageId: "success", // TODO: Return idempotent messageid in response https://github.com/bibyen/posta-baut/issues/15
	}
	return connect.NewResponse(resp), nil
}
