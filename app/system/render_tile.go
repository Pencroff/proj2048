package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/ai-agents/proj2048/component"
	"github.com/pencroff/ai-agents/proj2048/helper"
	"github.com/sedyh/mizu/pkg/engine"
	"image"
)

type RenderTile struct {
	*component.TilePropWrap
}

func (r *RenderTile) Update(_ engine.World) {
	p := r.Ptr
	if p.IsMoving {
		dist := helper.ManhattanDistance(p.Position, p.EndPosition)
		spDist := helper.ManhattanDistance(p.Speed, image.Point{})
		if spDist > dist {
			p.IsMoving = false
			p.Position = p.EndPosition
		} else {
			p.Position = p.Position.Add(p.Speed)
		}
	}
}

func (r *RenderTile) Draw(_ engine.World, screen *ebiten.Image) {
	p := r.Ptr
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.Sprite, op)
}
