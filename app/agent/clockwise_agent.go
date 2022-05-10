package agent

import (
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/stats"
)

type ClockwiseAgent struct {
	GenericAgent
	directionMap map[int]common.Direction
}

func (a *ClockwiseAgent) MakeMove(step int, _ int, noMove bool, _ []int) (direction common.Direction) {
	if noMove {
		a.noMoveCounter += 1
	}
	idx := (step + a.noMoveCounter) % 4
	direction = a.directionMap[idx]
	return
}

func NewClockwiseAgent(gameId int, recorder *stats.StatRecorder) Agent {
	return &ClockwiseAgent{
		GenericAgent: NewGenericAgent("clockwise_agent", "Clockwise Agent", false, gameId, recorder),
		directionMap: map[int]common.Direction{
			0: common.Up,
			1: common.Right,
			2: common.Down,
			3: common.Left,
		},
	}
}
