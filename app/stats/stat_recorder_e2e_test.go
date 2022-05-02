package stats

import (
	"context"
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"os"
	"testing"
	"time"
)

var envMap = map[string]string{
	"PG_IMAGE_TAG": "14.2-alpine",
	"PG_PORT":      "54321",
	"PG_USER":      "tst_user",
	"PG_PASSWORD":  "tst_password",
	"PG_DB_NAME":   "test_db",
}

type PgStrategy struct {
}

func (s PgStrategy) WaitUntilReady(ctx context.Context, str wait.StrategyTarget) error {
	log.Printf("[Info] Ctx: %v\n", ctx)
	log.Printf("[Info] Strategy: %v\n", str)
	r, e := str.Logs(ctx)
	if e != nil {
		log.Fatalf("[ERROR] Can't get logs: %v", e)
	}
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r, buf); err != nil {
		log.Fatalf("CopyBuf err: %v\n", err)
	}
	return nil
}

func TestMigrations(t *testing.T) {
	compose := NewComposeTestEnv([]string{"./../../docker-compose.e2e.yml"})
	compose.SetEnv(envMap)
	compose.SetStrategy("db", PgStrategy{})
	compose.Up()
	defer compose.Down()
	d := time.Second * 3
	time.Sleep(d)
	log.Printf("[Info] Connecting to database")
	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", envMap["PG_PORT"],
		envMap["PG_USER"], envMap["PG_PASSWORD"],
		envMap["PG_DB_NAME"])
	db, err := sqlx.Connect("postgres", pgConnStr)
	if err != nil {
		log.Fatalf("[ERROR] Can't connect to db: %v", err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./../../migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("[ERROR] Can't make migration: %v", err)
	}
	m.Up()
	v := &[]string{}
	db.Select(v, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	assert.Equal(t, len(*v), 3)
	assert.Contains(t, *v, "game_step")
	assert.Contains(t, *v, "game_stat")
	m.Down()
	db.Select(v, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	assert.Equal(t, len(*v), 1)
}
