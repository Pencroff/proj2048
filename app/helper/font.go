package helper

import (
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
)

func BoundStringImage(ff font.Face, s string) (rect image.Rectangle, zeroPoint image.Point) {
	int26Rect, _ := font.BoundString(ff, s)
	rect = ToImageRect(int26Rect)
	zeroPoint = image.Point{X: -rect.Min.X, Y: -rect.Min.Y}
	return
}

func ToImageRect(v fixed.Rectangle26_6) image.Rectangle {
	return image.Rect(v.Min.X.Round(), v.Min.Y.Round(), v.Max.X.Round(), v.Max.Y.Round())
}
