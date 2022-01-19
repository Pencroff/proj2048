package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type FieldProps struct {
	Position image.Point
	Sprite   *ebiten.Image
}
