package stats

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/OneOfOne/xxhash"
	"github.com/jmoiron/sqlx"
	"github.com/pencroff/proj2048/app/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatRecorder_GetGameStatId(t *testing.T) {
	hasher := xxhash.New64()
	lst := []struct {
		agentId  string
		mode     string
		gameId   int
		expected int
	}{
		{"", "", 0,
			int(xxhash.ChecksumString64("::0000000000000000") &^ (1 << 63))},
		{"a", "b", 1,
			int(xxhash.ChecksumString64("a:b:0000000000000001") &^ (1 << 63))},
		{"clockwise_agent", "simulation", 65535,
			int(xxhash.ChecksumString64("clockwise_agent:simulation:000000000000ffff") &^ (1 << 63))},
	}
	r := NewStatRecorder(nil)
	for _, el := range lst {
		hasher.Reset()
		actual := r.GetGameStatId(el.agentId, el.mode, el.gameId)
		assert.Equal(t, actual, el.expected)
	}
}

func TestStatRecorder_StartGame(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer mockDB.Close()
	assert.NoError(t, err)
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	recorder := NewStatRecorder(sqlxDB)
	agentId := "agent_id"
	mode := "mode"
	gameId := 1
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fieldState := &FieldState{startField, nil}
	j, _ := fieldState.ToJson()
	gameStatId := recorder.GetGameStatId(agentId, mode, gameId)
	mock.ExpectExec("INSERT INTO game_stat").
		WithArgs(gameStatId, 1, "agent_id", "mode", j).
		WillReturnResult(sqlmock.NewResult(int64(gameStatId), 1))

	id, err := recorder.StartGame(agentId, mode, gameId, startField)
	assert.NoError(t, err)
	assert.Equal(t, gameStatId, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStatRecorder_AddStep(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer mockDB.Close()
	assert.NoError(t, err)
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	recorder := NewStatRecorder(sqlxDB)
	agentId := "agent_id"
	mode := "mode"
	gameId := 1
	step := 24
	score := 768
	noMove := false
	fieldState := []int{2, 0, 0, 0, 4, 0, 2, 64, 0, 16, 0, 0, 0, 32, 0, 0}
	fieldStateParam := "[2,0,0,0,4,0,2,64,0,16,0,0,0,32,0,0]"
	direction := common.Up
	gameStatId := recorder.GetGameStatId(agentId, mode, gameId)
	mock.ExpectExec("INSERT INTO game_step").
		WithArgs(gameStatId, step, score, noMove, fieldStateParam, direction).
		WillReturnResult(sqlmock.NewResult(int64(gameStatId), 1))

	err = recorder.AddStep(gameStatId, step, score, noMove, fieldState, direction)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStatRecorder_FinishGame(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer mockDB.Close()
	assert.NoError(t, err)
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	recorder := NewStatRecorder(sqlxDB)
	agentId := "agent_id"
	mode := "mode"
	gameId := 1
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	endField := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	score := 16384
	stepCnt := 100
	noMoveCnt := 10
	maxTile := 256
	fieldState := &FieldState{startField, nil}
	dbJson, _ := fieldState.ToJson()
	fieldState.End = endField
	updateJson, _ := fieldState.ToJson()
	gameStatId := recorder.GetGameStatId(agentId, mode, gameId)
	rows := sqlmock.NewRows([]string{"id", "field_state"}).
		AddRow(gameStatId, dbJson)
	mock.ExpectQuery("SELECT id, field_state FROM game_stat WHERE id").
		WithArgs(gameStatId).
		WillReturnRows(rows)
	mock.ExpectExec("UPDATE game_stat").
		WithArgs(gameStatId, updateJson, score, stepCnt, noMoveCnt, maxTile).
		WillReturnResult(sqlmock.NewResult(int64(gameStatId), 1))

	err = recorder.FinishGame(gameStatId, endField, score, stepCnt, noMoveCnt, maxTile)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
