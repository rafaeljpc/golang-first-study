package http

import (
	"context"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server interface {
	Start() error
}

type HttpServer struct {
	Server *echo.Echo
	Addr string
}

func NewHttpServer(ctx context.Context) *HttpServer {
	gZipLevel := 5
	addr := os.Getenv("HOST_ADDRESS")

	app := echo.New()
	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gZipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "/swagger")
		},
	}))
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	return &HttpServer{
		Server: app,
		Addr: addr,
	}
}

func (s *HttpServer) Start() error {
	return s.Server.Start(s.Addr)
}