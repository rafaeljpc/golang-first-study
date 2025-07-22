package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/rafaeljpc/golang-first-study/internal/util"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func NewPostgresContainer(t *testing.T) *postgres.PostgresContainer {
	t.Helper()

	ctr, err := postgres.Run(t.Context(), "postgres:16-alpine",
		postgres.WithDatabase("go_test"),
		postgres.WithUsername("go_test"),
		postgres.WithPassword("a"),
		postgres.BasicWaitStrategies(),
		testcontainers.WithTmpfs(map[string]string{"/var/lib/postgresql/data": "rw"}),
	)
	if err != nil {
		log.Fatalln("failed to start container", err)
		return nil
	}

	return ctr
}

func GetTestConnection(t *testing.T, ctr *postgres.PostgresContainer) (*sql.DB, error) {
	t.Helper()

	if !ctr.IsRunning() {
		return nil, fmt.Errorf("container is not running")
	}

	connStr, err := ctr.ConnectionString(t.Context(), "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to get connection string: %w", err)
	}

	log.Default().Printf("Connecting to database at %s\n", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db, err
}

func StartTestDb(t *testing.T, db *sql.DB, withData bool) error {
	t.Helper()

	err := runScript(t, db, "script/sql/create_tables.sql")
	if err != nil {
		return err
	}
	err = runScript(t, db, "script/sql/insert_products.sql")
	if err != nil {
		return err
	}

	return nil
}

func runScript(t *testing.T, db *sql.DB, filename string) error {
	t.Helper()

	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s script: %w", filename, err)
	}

	script := strings.Split(string(b), ";\n")
	for _, stmt := range script {
		trimmedStmt := strings.TrimSpace(stmt)
		log.Default().Println("Executing statement: ", util.Substring(trimmedStmt, 0, 20))
		_, err = db.Exec(trimmedStmt)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}
	return nil
}
