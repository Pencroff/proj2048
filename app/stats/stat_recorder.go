package stats

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"github.com/jmoiron/sqlx"
)

func NewStatRecorder(db *sqlx.DB) *StatRecorder {
	return &StatRecorder{db}
}

type StatRecorder struct {
	*sqlx.DB
}

func (r *StatRecorder) GetGameStatId(agentId string, mode string, gameId int) int {
	v := fmt.Sprintf("%s:%s:%016x", agentId, mode, gameId)
	return int(xxhash.ChecksumString64(v) &^ (1 << 63))
}

func (r *StatRecorder) StartGame(agentId string, mode string, gameId int,
	startField []int) (gameStatId int, err error) {
	gameStatId = r.GetGameStatId(agentId, mode, gameId)
	j, err := (&FieldState{startField, nil}).ToJson()
	if err != nil {
		return
	}
	_, err = r.DB.Exec(`
INSERT INTO game_stat 
    (id, game_id, agent_id, mode, field_state)
VALUES
    ($1, $2, $3, $4, $5) 
ON CONFLICT (id) 
DO 
   UPDATE SET game_id = $2, agent_id = $3, mode = $4, field_state = $5;`,
		gameStatId, gameId, agentId, mode, j)
	return
}

func (r *StatRecorder) AddStep(gameStatId int, step int,
	score int, noMove bool, gameField []int,
	direction string) (err error) {
	return
}

func (r *StatRecorder) FinishGame(gameStatId int, endField []int,
	score int, stepCnt int, noMoveCnt int, maxTile int) (err error) {
	return
}
