package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/proj2048/app/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type RenderField struct {
	*component.FieldProps
}

func (r *RenderField) Draw(_ engine.World, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.Position.X), float64(r.Position.Y))
	screen.DrawImage(r.Sprite, op)
}
