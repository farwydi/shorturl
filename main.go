package main

import (
	"context"
	"github.com/farwydi/shorturl/domain"
	"log"

	"github.com/drone/signal"
	"github.com/farwydi/shorturl/endpoint/web"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("fail init logger: %v\n", err)
	}

	logger = logger.Named("main")

	cfg := domain.Config{}

	app, cleanup, err := setup(cfg, logger)
	if err != nil {
		logger.Fatal("fail setup application",
			zap.Error(err))
	}
	defer cleanup()

	ctx := signal.WithContext(context.Background())
	eg := errgroup.Group{}

	eg.Go(func() error {
		logger.Info("Run web server")
		return app.web.Run(ctx)
	})

	if err := eg.Wait(); err != nil {
		logger.Fatal("app terminated",
			zap.Error(err))
	}
}

func newApplication(web *web.Server) application {
	return application{
		web: web,
	}
}

type application struct {
	web *web.Server
}

func initLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
