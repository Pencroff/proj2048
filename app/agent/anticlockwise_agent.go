package agent

import (
	"github.com/pencroff/ai-agents/proj2048/common"
)

type AnticlockwiseAgent struct {
	GenericAgent
	directionMap map[int]common.Direction
}

func (a *AnticlockwiseAgent) MakeMove(step int, _ int, noMove bool, _ []int) (direction common.Direction) {
	if noMove {
		a.noMoveCounter += 1
	}
	idx := (step + a.noMoveCounter) % 4
	direction = a.directionMap[idx]
	return
}

func NewAnticlockwiseAgent(gameId int64) Agent {
	return &AnticlockwiseAgent{
		GenericAgent: NewGenericAgent("anticlockwise_agent", "Anticlockwise Agent", false, gameId),
		directionMap: map[int]common.Direction{
			0: common.Up,
			1: common.Left,
			2: common.Down,
			3: common.Right,
		},
	}
}
