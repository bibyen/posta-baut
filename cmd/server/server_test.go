package server

import (
	"testing"

	"github.com/bibyen/posta-baut/internal/pb/v1/pbv1connect"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	type fakeTeamsService struct {
		pbv1connect.UnimplementedTeamsServiceHandler
	}

	t.Run("successful server initialisation", func(t *testing.T) {
		svc := &fakeTeamsService{}
		_, err := NewServer(svc)
		require.NoError(t, err)
	})
}

// TODO: TestServerEndpoints tests that all expected endpoints are registered - https://github.com/bibyen/posta-baut/issues/12
func TestServerEndpoints(t *testing.T) {
}

// TODO: TestServerGracefulShutdown tests the graceful shutdown mechanism - https://github.com/bibyen/posta-baut/issues/12
func TestServerGracefulShutdown(t *testing.T) {
}

// TODO: TestServerShutdownTimeout tests shutdown timeout behavior - https://github.com/bibyen/posta-baut/issues/12
func TestServerShutdownTimeout(t *testing.T) {
}

// TODO: TestServerConfiguration tests server configuration parameters - https://github.com/bibyen/posta-baut/issues/12
func TestServerConfiguration(t *testing.T) {
}
