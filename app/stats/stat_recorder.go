package stats

import (
	"encoding/json"
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
   UPDATE SET game_id = $2, agent_id = $3, mode = $4, field_state = $5;
`,
		gameStatId, gameId, agentId, mode, j)
	return
}

func (r *StatRecorder) AddStep(gameStatId int, step int,
	score int, noMove bool, gameField []int,
	direction string) (err error) {
	byteData, err := json.Marshal(gameField)
	if err != nil {
		return
	}
	_, err = r.DB.Exec(`
INSERT INTO game_step
	(game_stat_id, step, score, no_move, field, direction)
VALUES
	($1, $2, $3, $4, $5, $6)
ON CONFLICT (game_stat_id, step) 
DO 
   UPDATE SET score = $3, no_move = $4, field = $5, direction = $6;
`,
		gameStatId, step, score, noMove, string(byteData), direction)
	return
}

func (r *StatRecorder) FinishGame(gameStatId int, endField []int,
	score int, stepCnt int, noMoveCnt int, maxTile int) (err error) {
	var data GameStat
	err = r.DB.QueryRowx("SELECT id, field_state FROM game_stat WHERE id=$1", gameStatId).StructScan(&data)
	if err != nil {
		return
	}
	fieldState, err := data.GetFieldStateData()
	if err != nil {
		return
	}
	fieldState.End = endField
	j, err := fieldState.ToJson()
	if err != nil {
		return
	}
	_, err = r.DB.Exec(`
UPDATE game_stat 
	SET field_state = $2, score = $3, step_counter = $4, no_move_counter = $5, max_tile = $6
WHERE id = $1`,
		gameStatId, j, score, stepCnt, noMoveCnt, maxTile)
	return
}
