package stats

import "encoding/json"

type GameStat struct {
	Id             int    `db:"id"`
	GameId         int    `db:"game_id"`
	AgentId        string `db:"agent_id"`
	Mode           string `db:"mode"`
	FieldStateJson string `db:"field_state"`
	Score          int    `db:"score"`
	StepCnt        int    `db:"step_counter"`
	NoMoveCnt      int    `db:"no_move_counter"`
	MaxTile        int    `db:"max_tile"`
}

func (s *GameStat) GetFieldStateData() (FieldState, error) {
	fieldStateLocal := FieldState{}
	err := fieldStateLocal.FromJson(s.FieldStateJson)
	return fieldStateLocal, err
}

func (s *GameStat) SetFieldStateJson(state FieldState) (err error) {
	s.FieldStateJson, err = state.ToJson()
	return err
}

type GameStep struct {
	GameStatId int    `db:"game_stat_id"`
	Step       int    `db:"step"`
	Score      int    `db:"score"`
	NoMove     bool   `db:"no_move"`
	FieldJson  string `db:"field"`
	Direction  string `db:"direction"`
}

func (s *GameStep) GetFieldData() ([]int, error) {
	var field []int
	err := json.Unmarshal([]byte(s.FieldJson), &field)
	return field, err
}

func (s *GameStep) SetFieldJson(v []int) error {
	byteData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	s.FieldJson = string(byteData)
	return nil
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
