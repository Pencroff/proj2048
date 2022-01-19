package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type TileProp struct {
	Value       int
	Position    image.Point
	EndPosition image.Point
	Speed       image.Point
	IsMoving    bool
	Removed     bool
	Sprite      *ebiten.Image
}

type TilePropWrap struct {
	Ptr *TileProp
}
