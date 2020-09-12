package web

import (
	"context"
	"net/http"
	"time"

	"github.com/farwydi/shorturl/domain"
	"golang.org/x/sync/errgroup"
)

func NewWebServer(config domain.Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:    config.Endpoint.Web.Addr,
			Handler: handler,
		},
	}
}

type Server struct {
	srv *http.Server
}

func (s *Server) Run(ctx context.Context) error {
	eg := errgroup.Group{}
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return s.srv.Shutdown(ctx)
		}
	})
	eg.Go(func() error {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return eg.Wait()
}
