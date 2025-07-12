package di

import (
	"context"

	"github.com/rafaeljpc/golang-first-study/internal/adapter/dummy"
	"github.com/rafaeljpc/golang-first-study/internal/adapter/http"
	"github.com/rafaeljpc/golang-first-study/internal/adapter/http/handlers"
	"github.com/rafaeljpc/golang-first-study/internal/domain/service"
)

type Container struct {
	service *service.Service

	ApiServer *http.HttpServer
}

func NewContainer() *Container {
	container := &Container{}
	container.init()

	return container
}

func (c *Container) init() {
	ctx := context.Background()
	repository := dummy.NewDummyRepository()

	c.service = service.NewService(repository)

	c.ApiServer = http.NewHttpServer(ctx)
	
	handler := handlers.NewHttpServiceHandler(c.service)
	handler.RegisterRoutes(c.ApiServer.Server)

	c.ApiServer.Start()
}
