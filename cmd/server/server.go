package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/bibyen/posta-baut/internal/pb/v1/pbv1connect"
	pbconnect "github.com/bibyen/posta-baut/internal/pb/v1/pbv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server encapsulates the HTTP server for handling gRPC requests.
type Server struct {
	server *http.Server
}

// NewServer creates and configures a new Server instance.
func NewServer(teamsSvc pbv1connect.TeamsServiceHandler) (Server, error) {
	if teamsSvc == nil {
		return Server{}, errors.New("teams service is nil")
	}

	mux := http.NewServeMux()

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(pbconnect.TeamsServiceName),
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(pbconnect.TeamsServiceName),
	))

	path, handler := pbconnect.NewTeamsServiceHandler(teamsSvc)
	mux.Handle(path, handler)

	// --- HTTP Server Configuration ---
	addr := fmt.Sprintf(":%d", 8080)
	httpServer := &http.Server{
		Addr: addr,
		// The `h2c` handler enables "HTTP/2 Cleartext", allowing gRPC
		// to function over an unencrypted HTTP/2 connection. In a production environment,
		// you would either configure TLS directly here or terminate TLS at a load balancer.
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	return Server{
		server: httpServer,
	}, nil
}

// Run starts the server and blocks until a shutdown signal is received.
func (s *Server) Run() error {
	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// --- Server Startup ---
	go func() {
		log.Printf("Server is listening on %s\n", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server failed to start: %v\n", err)
			stop() // Trigger shutdown if server fails to start
		}
	}()

	<-shutdownCtx.Done() // Block here until a shutdown signal is received
	log.Println("Shutdown signal received, shutting down server...")
	shutdownTimeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(shutdownTimeoutCtx); err != nil {
		return fmt.Errorf("server graceful shutdown failed: %w", err)
	}

	return nil
}
