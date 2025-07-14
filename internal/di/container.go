package di

import (
	"context"
	"database/sql"
	"errors"
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
	sqlConnection, err := c.createPostgresDBConnection()
	if err != nil {
		return
	}	
	defer func() {
		err = errors.Join(sqlConnection.Close())		
	}()
	repository := postgres.NewPostgresRepository(sqlConnection)

	c.service = service.NewService(repository)

	c.ApiServer = http.NewHttpServer(ctx)

	handler := handlers.NewHttpServiceHandler(c.service)
	handler.RegisterRoutes(c.ApiServer.Server)

	err = c.ApiServer.Start()
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}

func (c *Container) createPostgresDBConnection() (*sql.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db, err
}
