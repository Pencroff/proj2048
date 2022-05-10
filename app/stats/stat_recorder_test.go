package stats

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/OneOfOne/xxhash"
	"github.com/jmoiron/sqlx"
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
	j, err := fieldState.ToJson()
	assert.NoError(t, err)
	assert.NoError(t, err)
	gameStatId := recorder.GetGameStatId(agentId, mode, gameId)
	mock.ExpectExec("INSERT INTO game_stat").
		WithArgs(gameStatId, 1, "agent_id", "mode", j).
		WillReturnResult(sqlmock.NewResult(int64(gameStatId), 1))

	id, err := recorder.StartGame(agentId, mode, gameId, startField)
	assert.NoError(t, err)
	assert.Equal(t, gameStatId, id)
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
