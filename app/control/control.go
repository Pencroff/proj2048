package control

import (
	"image"
	"image/color"
)

type HorizontalAlign string

const (
	Left   HorizontalAlign = "left"
	Center                 = "center"
	Right                  = "right"
)

// box-sizing: border-box

type Control struct {
	BorderRadius int
	Padding      image.Point // Support just equal padding? should it be top/bottom, left/right, same for margin below
	Size         image.Point
	Ink          color.Color
	Paper        color.Color
}
