package server

import (
	"context"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, r http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
			Handler: r,
		},
	}
}

func (srv *Server) StartServer() error {
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) ShutdownServer(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
