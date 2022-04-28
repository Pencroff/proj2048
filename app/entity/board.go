package entity

import (
	"github.com/Pencroff/fluky"
	"github.com/Pencroff/fluky/rng"
	"github.com/pencroff/proj2048/app/agent"
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/pencroff/proj2048/app/stats"
	"time"
)

type Board struct {
	component.BoardProp
}

func NewBoard(recorder *stats.StatRecorder) (field *Field, board *Board) {
	size := resources.Config.FieldSize
	headerOffset := resources.Config.HeaderRect.Size()
	headerOffset.X = 0
	fieldRect := resources.Config.FieldRect.Add(headerOffset)
	boardRect := resources.Config.BoardRect
	field = NewField(size, fieldRect, resources.FieldColor)
	agentMap := CreateAgentMap(recorder)
	agent := agentMap[common.Manual]

	board = &Board{
		component.BoardProp{
			Name:        "2048",
			Description: agent.GetName(),
			Score:       0,
			Step:        0,
			NoMove:      false,
			Speed:       resources.Config.Speed,
			State:       common.StartGame,
			Color:       resources.BoardColor,
			Size:        size,
			BoardRect:   boardRect,
			Direction:   common.NoDirection,
			FieldProps:  &field.FieldProps,
			List:        make([]*component.TileProp, size.X*size.Y),
			Mode:        common.Manual,
			Agent:       agent,
			Flk:         fluky.NewFluky(rng.NewSmallPrng()),
			AgentMap:    agentMap,
		},
	}
	return
}

func CreateAgentMap(recorder *stats.StatRecorder) map[common.PlayMode]agent.Agent {
	gameId := time.Now().UTC().UnixNano() % (1 << 16)
	m := make(map[common.PlayMode]agent.Agent)
	m[common.Manual] = agent.NewHumanAgent(gameId, recorder)
	m[common.AutoPlay] = agent.NewPoolAgent(gameId, recorder)
	return m
}
