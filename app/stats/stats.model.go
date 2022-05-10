package stats

import "encoding/json"

type GameStat struct {
	Id         int    `db:"id"`
	GameId     int    `db:"game_id"`
	AgentId    string `db:"agent_id"`
	Mode       string `db:"mode"`
	FieldState string `db:"field_state"`
	Score      int    `db:"score"`
	StepCnt    int    `db:"step_counter"`
	NoMoveCnt  int    `db:"no_move_counter"`
	MaxTile    int    `db:"max_tile"`
}

type FieldState struct {
	Start []int
	End   []int
}

func (s *FieldState) ToJson() (j string, err error) {
	b, err := json.Marshal(s)
	j = string(b)
	return
}

func (s *FieldState) FromJson(j string) (err error) {
	err = json.Unmarshal([]byte(j), s)
	return
}
