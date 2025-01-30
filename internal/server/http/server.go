package http

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"
	"transaction/internal/config"
	"transaction/internal/logger"
	"transaction/internal/server/http/handler"
	"transaction/internal/server/http/router"
)

var (
	onceInit   sync.Once
	httpServer *Server
)

type Server struct {
	server *http.Server
	config *config.HttpServerConfig
}

func NewServer(ctx context.Context, config *config.HttpServerConfig, accountHandler *handler.AccountHandler, transactionHandler *handler.TransactionHandler) *Server {
	onceInit.Do(func() {
		ginEngine := router.ConfigureRoutes(ctx, accountHandler, transactionHandler)
		httpServer = &Server{
			server: &http.Server{
				Addr:    ":" + config.Port,
				Handler: ginEngine,
			},
			config: config,
		}
	})
	return httpServer
}

func (s *Server) Start(ctx context.Context) {
	logger.WithContext(ctx).Infof("Starting http server on port: %s", s.config.Port)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.WithContext(ctx).Errorf("Error starting http server, port: %s, with error: %v", s.config.Port, err)
		panic(err)
	}
}

func (s *Server) ShutDownGracefully() {
	s.server.SetKeepAlivesEnabled(false)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(500))
	defer cancel()

	if s.server != nil {
		if err := s.server.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Fatalf("HTTP Server Shutdown Failed, err: %v", err)
		}
	}
	logger.WithContext(ctx).Info("HTTP Server stopped")
}
