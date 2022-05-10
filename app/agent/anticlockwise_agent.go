package agent

import (
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/stats"
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

func NewAnticlockwiseAgent(gameId int, recorder *stats.StatRecorder) Agent {
	return &AnticlockwiseAgent{
		GenericAgent: NewGenericAgent("anticlockwise_agent", "Anticlockwise Agent", false, gameId, recorder),
		directionMap: map[int]common.Direction{
			0: common.Up,
			1: common.Left,
			2: common.Down,
			3: common.Right,
		},
	}
}
