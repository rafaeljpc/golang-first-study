package test

import (
	"log"
	"testing"

	"github.com/rafaeljpc/golang-first-study/internal/adapter/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func Test_E2E_ListProducts(t *testing.T) {
	// Given
	// ctx := t.Context()
	postgresContainer := NewPostgresContainer(t)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Fatalf("failed to terminate container: %v", err)
		}
	}()

	db, err := GetTestConnection(t, postgresContainer)
	assert.NoError(t, err)

	err = StartTestDb(t, db, true)
	assert.NoError(t, err)

	repository := postgres.NewPostgresRepository(db)

	// When
	products, err := repository.ListProducts()

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}
