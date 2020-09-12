// +build wireinject

package main

import (
	"github.com/farwydi/shorturl/domain"
	"github.com/farwydi/shorturl/endpoint/web"
	"github.com/farwydi/shorturl/gateway/self"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func setup(domain.Config, *zap.Logger) (application, func(), error) {
	panic(wire.Build(
		self.NewSelfDataGateway,
		web.NewGinEngine,
		web.NewRouter,
		web.NewWebServer,
		newApplication,
	))
}
