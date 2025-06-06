package di

import (
	"context"
	"mygogo/hello/internal/adapter/http"
	"mygogo/hello/internal/domain/service"
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

func (c *Container) init()  {
	ctx := context.Background()
	c.service = service.NewService()

	c.ApiServer = http.NewHttpServer(ctx)

	c.ApiServer.Start()
}