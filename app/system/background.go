package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/ai-agents/proj2048/helper"
	"github.com/sedyh/mizu/pkg/engine"
)

type BackgroundSystem struct{}

func (s *BackgroundSystem) Draw(_ engine.World, screen *ebiten.Image) {
	c, _ := helper.ColorStrToColor("#1c1917")
	screen.Fill(c)
}
