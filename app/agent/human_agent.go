package agent

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/ai-agents/proj2048/common"
)

type HumanAgent struct {
	GenericAgent
	keyPressed int
}

func (a *HumanAgent) MakeMove(_ int, _ int, _ bool, _ []int) (direction common.Direction) {
	delay := 5
	if a.keyPressed > 0 {
		a.keyPressed -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		a.keyPressed = delay
		direction = common.Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		a.keyPressed = delay
		direction = common.Right
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		a.keyPressed = delay
		direction = common.Down
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		a.keyPressed = delay
		direction = common.Left
	} else {
		direction = common.NoDirection
	}
	return
}

func NewHumanAgent(gameId int64) Agent {
	return &HumanAgent{
		GenericAgent: NewGenericAgent("haman_agent", "Human Agent", true, gameId),
	}
}
