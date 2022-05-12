package stats

import (
	"context"
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pencroff/proj2048/app/common"
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
	compose, db, m, err := SetupEnv()
	defer compose.Down()
	if err != nil {
		log.Fatalf("[ERROR] Can't setup env: %v", err)
	}

	v := &[]string{}
	db.Select(v, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	assert.Equal(t, len(*v), 3)
	assert.Contains(t, *v, "game_step")
	assert.Contains(t, *v, "game_stat")
	m.Down()
	db.Select(v, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	assert.Equal(t, len(*v), 1)
}

func TestDataRecord_GameStat_StartGame(t *testing.T) {
	compose, db, _, err := SetupEnv()
	defer compose.Down()
	if err != nil {
		log.Fatalf("[ERROR] Can't setup env: %v", err)
	}
	rec := NewStatRecorder(db)
	agent_id := "agent_id"
	mode := "simulation"
	game_id := 5
	fs := FieldState{
		Start: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		End:   nil,
	}
	id, err := rec.StartGame(agent_id, mode, game_id, fs.Start)
	assert.NoError(t, err)
	var data GameStat
	err = db.QueryRowx("SELECT id, agent_id, mode, game_id, field_state FROM game_stat WHERE id=$1", id).StructScan(&data)
	assert.NoError(t, err)
	fsDb, err := data.GetFieldStateData()
	assert.NoError(t, err)
	assert.Equal(t, data.Id, id)
	assert.Equal(t, data.AgentId, agent_id)
	assert.Equal(t, data.Mode, mode)
	assert.Equal(t, data.GameId, game_id)
	assert.Equal(t, fsDb, fs)
}

func TestDataRecord_GameStat_FinishGame(t *testing.T) {
	compose, db, _, err := SetupEnv()
	defer compose.Down()
	if err != nil {
		log.Fatalf("[ERROR] Can't setup env: %v", err)
	}
	rec := NewStatRecorder(db)
	agent_id := "agent_id"
	mode := "simulation"
	game_id := 5
	fs := FieldState{
		Start: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		End:   []int{128, 64, 32, 16, 8, 4, 2, 1, 1, 2, 4, 8, 16, 32, 64, 128},
	}
	score := 768
	stepCnt := 16
	noMoveCnt := 3
	maxTile := 128

	id, _ := rec.StartGame(agent_id, mode, game_id, fs.Start)
	err = rec.FinishGame(id, fs.End, score, stepCnt, noMoveCnt, maxTile)
	assert.NoError(t, err)
	var data GameStat
	err = db.QueryRowx("SELECT * FROM game_stat WHERE id=$1", id).StructScan(&data)
	assert.NoError(t, err)
	fsDb, err := data.GetFieldStateData()
	assert.NoError(t, err)
	assert.Equal(t, data.Id, id)
	assert.Equal(t, data.AgentId, agent_id)
	assert.Equal(t, data.Mode, mode)
	assert.Equal(t, data.GameId, game_id)
	assert.Equal(t, fsDb, fs)
	assert.Equal(t, data.Score, score)
	assert.Equal(t, data.StepCnt, stepCnt)
	assert.Equal(t, data.NoMoveCnt, noMoveCnt)
	assert.Equal(t, data.MaxTile, maxTile)
}

func TestDataRecord_GameStat_AddStep(t *testing.T) {
	compose, db, _, err := SetupEnv()
	defer compose.Down()
	if err != nil {
		log.Fatalf("[ERROR] Can't setup env: %v", err)
	}
	rec := NewStatRecorder(db)
	agent_id := "agent_id"
	mode := "simulation"
	game_id := 5
	stepNum := 32
	score := 512
	noMove := true
	gameFs := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	direction := common.Right
	id, _ := rec.StartGame(agent_id, mode, game_id, []int{})
	err = rec.AddStep(id, stepNum, score, noMove, gameFs, direction)
	assert.NoError(t, err)
	var data GameStep
	err = db.QueryRowx("SELECT * FROM game_step WHERE game_stat_id=$1 AND step=$2", id, stepNum).StructScan(&data)
	assert.NoError(t, err)
	fsDb, err := data.GetFieldData()
	assert.NoError(t, err)
	assert.Equal(t, data.GameStatId, id)
	assert.Equal(t, data.Step, stepNum)
	assert.Equal(t, data.Score, score)
	assert.Equal(t, data.NoMove, noMove)
	assert.Equal(t, fsDb, gameFs)
}

func SetupEnv() (compose *ComposeTestEnv, db *sqlx.DB, m *migrate.Migrate, err error) {
	compose = NewComposeTestEnv([]string{"./../../docker-compose.e2e.yml"})
	compose.SetEnv(envMap)
	// Uncomment row below to show container logs
	//compose.SetStrategy("db", PgStrategy{})
	compose.Up()
	//defer compose.Down()
	d := time.Second * 3
	time.Sleep(d)
	log.Printf("[Info] Connecting to database")
	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", envMap["PG_PORT"],
		envMap["PG_USER"], envMap["PG_PASSWORD"],
		envMap["PG_DB_NAME"])
	db, err = sqlx.Connect("postgres", pgConnStr)
	if err != nil {
		log.Printf("[ERROR] Can't connect to db")
		return
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Printf("[ERROR] Can't make driver")
		return
	}
	m, err = migrate.NewWithDatabaseInstance(
		"file://./../../migrations",
		"postgres", driver)
	if err != nil {
		log.Printf("[ERROR] Can't make migration")
		return
	}
	m.Up()
	return
}
