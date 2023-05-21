package server

import (
	"context"
	"net"
	"net/http"
	"new/pkg/e"
	"time"
)

type Server struct {
	S http.Server
}

func (s *Server) Start(host *string, router *http.ServeMux, ctx context.Context) error {
	n := net.ListenConfig{}
	lis, err := n.Listen(ctx, "tcp", *host)
	if err != nil {
		return e.Err("cen't stsrt server", err)
	}

	s.S = http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	if err = s.S.Serve(lis); err != nil {
		return e.Err("cen't stsrt server", err)
	}
	return nil
}
