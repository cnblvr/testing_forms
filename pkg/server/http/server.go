package serverHttp

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

type Config interface {
	HttpServerPort() string
	HttpReadTimeout() time.Duration
	HttpWriteTimeout() time.Duration
}

func New(cfg Config, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + cfg.HttpServerPort(),
			Handler:      handler,
			ReadTimeout:  cfg.HttpReadTimeout(),
			WriteTimeout: cfg.HttpWriteTimeout(),
		},
	}
}

func (s *Server) Run() error {
	err := s.server.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return errors.WithStack(err)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
