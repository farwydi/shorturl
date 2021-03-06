// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/farwydi/shorturl/domain"
	"github.com/farwydi/shorturl/endpoint/web"
	"github.com/farwydi/shorturl/gateway/self"
	"go.uber.org/zap"
)

// Injectors from inject.go:

func setup(config domain.Config, logger *zap.Logger) (application, func(), error) {
	engine := web.NewGinEngine()
	dataGateway, err := self.NewSelfDataGateway(config)
	if err != nil {
		return application{}, nil, err
	}
	handler := web.NewRouter(engine, dataGateway)
	server := web.NewWebServer(config, handler)
	mainApplication := newApplication(server)
	return mainApplication, func() {
	}, nil
}
