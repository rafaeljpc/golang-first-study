package di

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/rafaeljpc/golang-first-study/internal/adapter/http"
	"github.com/rafaeljpc/golang-first-study/internal/adapter/http/handlers"
	"github.com/rafaeljpc/golang-first-study/internal/adapter/postgres"
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
	repository := postgres.NewPostgresRepository(c.createPostgresDBConnection())

	c.service = service.NewService(repository)

	c.ApiServer = http.NewHttpServer(ctx)

	handler := handlers.NewHttpServiceHandler(c.service)
	handler.RegisterRoutes(c.ApiServer.Server)

	err := c.ApiServer.Start()
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}

func (c *Container) createPostgresDBConnection() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db
}
