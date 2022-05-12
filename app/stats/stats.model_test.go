package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameStat_GetFieldState(t *testing.T) {
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	endField := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	gameStat := GameStat{}
	_, e := gameStat.GetFieldStateData()
	assert.Error(t, e)
	gameStat.FieldStateJson = `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":null}`
	fieldState, e := gameStat.GetFieldStateData()
	assert.Equal(t, fieldState, FieldState{startField, nil})
	assert.NoError(t, e)
	gameStat.FieldStateJson = `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]}`
	fieldState, e = gameStat.GetFieldStateData()
	assert.Equal(t, fieldState, FieldState{startField, endField})
	assert.NoError(t, e)
}

func TestGameStat_SetFieldState(t *testing.T) {
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	endField := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	gameStat := GameStat{}
	fs := FieldState{}
	gameStat.SetFieldStateJson(fs)
	assert.Equal(t, gameStat.FieldStateJson, `{"Start":null,"End":null}`)
	fs.Start = startField
	assert.Equal(t, gameStat.FieldStateJson, `{"Start":null,"End":null}`)
	gameStat.SetFieldStateJson(fs)
	assert.Equal(t, gameStat.FieldStateJson, `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":null}`)
	fs.End = endField
	gameStat.SetFieldStateJson(fs)
	assert.Equal(t, gameStat.FieldStateJson, `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]}`)
}

func TestGameStep_GetFieldData(t *testing.T) {
	gameStep := GameStep{}
	_, e := gameStep.GetFieldData()
	assert.Error(t, e)
	gameStep.FieldJson = `[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]`
	field, e := gameStep.GetFieldData()
	assert.Equal(t, field, []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2})
	assert.NoError(t, e)
}

func TestGameStep_SetFieldJson(t *testing.T) {
	gameStep := GameStep{}
	gameStep.SetFieldJson([]int{})
	assert.Equal(t, gameStep.FieldJson, `[]`)
	gameStep.SetFieldJson([]int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2})
	assert.Equal(t, gameStep.FieldJson, `[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]`)
}

func TestFieldState_ToJson(t *testing.T) {
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	endField := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	fieldState := FieldState{startField, nil}
	j, e := fieldState.ToJson()
	assert.NoError(t, e)
	assert.Equal(t, j, `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":null}`)
	fieldState = FieldState{startField, endField}
	j, e = fieldState.ToJson()
	assert.NoError(t, e)
	assert.Equal(t, j, `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]}`)
}

func TestFieldState_FromJson(t *testing.T) {
	startField := []int{2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	startJson := `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":null}`
	endField := []int{8, 128, 2, 4, 4, 8, 16, 2, 8, 2, 64, 32, 2, 4, 8, 2}
	fullJson := `{"Start":[2,0,0,0,4,0,0,0,0,0,0,0,0,0,0,0],"End":[8,128,2,4,4,8,16,2,8,2,64,32,2,4,8,2]}`
	fieldState := FieldState{}
	e := fieldState.FromJson(startJson)
	assert.NoError(t, e)
	assert.Equal(t, fieldState.Start, startField)
	fieldState = FieldState{}
	e = fieldState.FromJson(fullJson)
	assert.NoError(t, e)
	assert.Equal(t, fieldState.Start, startField)
	assert.Equal(t, fieldState.End, endField)
}

func TestFieldState_IncorrectInput(t *testing.T) {
	fieldState := FieldState{}
	e := fieldState.FromJson("{json}")
	assert.Error(t, e)
}
