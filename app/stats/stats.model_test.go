package stats

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	fmt.Printf("%v\n", e)
	assert.Error(t, e)
}
