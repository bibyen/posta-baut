// Package svc provides service-level utilities for the Posta Baut application.
package svc

import (
	"context"
	"fmt"
	"log"
	"strings"

	"connectrpc.com/connect"
	"github.com/bibyen/posta-baut/cmd/client"
	msgraph "github.com/bibyen/posta-baut/cmd/config"
	utils "github.com/bibyen/posta-baut/cmd/svc/utils"
	"github.com/bibyen/posta-baut/internal/credentialprovider"
	pb "github.com/bibyen/posta-baut/internal/pb/v1"
	"github.com/bibyen/posta-baut/internal/pb/v1/pbv1connect"
	"github.com/bibyen/posta-baut/pkg/messenger"
	msgr "github.com/bibyen/posta-baut/pkg/messenger"
	"google.golang.org/grpc/metadata"
)

const (
	DummyId = "hello"
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

func extractJwtToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata in context")
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return "", fmt.Errorf("authorization header not provided")
	}

	// Assuming authorization header is: "Bearer <token>"
	parts := strings.SplitN(authHeaders[0], " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

func parseUserIdFromToken(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is empty string")
	}
	return "", nil
}

func fetchCredentials(ctx context.Context, target messenger.MessageTarget) (*credentialprovider.MessengerCredentials, error) {
	token, err := extractJwtToken(ctx)
	if err != nil {
		log.Printf("unauthenticated request attempted. no jwt provided")
		return nil, fmt.Errorf("unauthenticated request")
	}

	userId, err := parseUserIdFromToken(token)
	if err != nil {
		log.Printf("unauthenticated request attempted. missing user id from provided jwt %s: %v", token, err)
		return nil, fmt.Errorf("unauthenticated request")
	}

	creds := credentialprovider.InMemCredentialProvider.LookupCredentials(ctx, userId, target)
	if creds == nil {
		return nil, fmt.Errorf("user id %s does not have permission to send messages to target: %v", userId, target)
	}

	return creds, nil
}

// SendMessage handles sending messages to Microsoft Teams.
func (s *teamsService) SendMessage(ctx context.Context, req *connect.Request[pb.SendMessageRequest]) (*connect.Response[pb.SendMessageResponse], error) {

	// Convert the request to domain-level Message - msgr.Message
	msg, err := utils.ReqToMsg(req.Msg)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request to message: %w", err)
	}

	creds, err := fetchCredentials(ctx, msg.Target)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("unauthenticated request"))
	}

	messengerCfg := msgr.MessengerConfig{
		Type: msgr.MessengerTypeGraph,
		GraphConfig: &msgraph.ClientConfig{
			ClientID:     creds.GraphCreds.ClientID,
			TenantID:     creds.GraphCreds.TenantID,
			ClientSecret: creds.GraphCreds.ClientSecret,
		},
	}

	if creds.GraphCreds == nil {
		return nil, fmt.Errorf("unsupported credential type found, %v", creds)
	}
	err = msgr.Send(ctx, messengerCfg, msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to send message: %w", err))
	}

	resp := &pb.SendMessageResponse{
		MessageId: "success", // TODO: Return idempotent messageid in response https://github.com/bibyen/posta-baut/issues/15
	}
	return connect.NewResponse(resp), nil
}
