package resources

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/ai-agents/proj2048/common"
	. "github.com/pencroff/ai-agents/proj2048/helper"
	"image"
)

func RenderField(size image.Point, rect image.Rectangle,
	color common.ColorPair) *ebiten.Image {
	spacePercent := Config.FieldSpace
	xSpacePx := RoundToInt(float64(rect.Dx()) * spacePercent)
	xSpaceAllPx := (size.X + 1) * xSpacePx

	ySpacePx := RoundToInt(float64(rect.Dy()) * spacePercent)
	ySpaceAllPx := (size.Y + 1) * ySpacePx

	tileSizeW := RoundToInt(float64(rect.Dx()-xSpaceAllPx) / float64(size.X))
	tileSizeH := RoundToInt(float64(rect.Dy()-ySpaceAllPx) / float64(size.Y))

	field := ebiten.NewImage(rect.Dx(), rect.Dy())
	field.Fill(color.Paper)

	tile := ebiten.NewImage(tileSizeW, tileSizeH)
	tile.Fill(color.Ink)

	for y := 0; y < size.X; y++ {
		for x := 0; x < size.Y; x++ {
			op := &ebiten.DrawImageOptions{}
			offsetX := float64(xSpacePx + x*(xSpacePx+tileSizeW))
			offsetY := float64(ySpacePx + y*(ySpacePx+tileSizeH))
			op.GeoM.Translate(offsetX, offsetY)
			field.DrawImage(tile, op)
		}
	}
	return field
}
