package resources

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/ai-agents/proj2048/common"
	"golang.org/x/image/font"
	"image"
	"strconv"
)

var TileMap map[int]*ebiten.Image

func InitTiles() {
	TileMap = map[int]*ebiten.Image{}
	size := image.Point{X: 100, Y: 100}
	for key, pair := range ColorMap {
		TileMap[key] = RenderTile(key, size, pair)
	}
}

func RenderTile(
	value int, sizePx image.Point, color common.ColorPair,
) *ebiten.Image {
	dc := gg.NewContext(sizePx.X, sizePx.Y)
	dc.SetRGBA255(0, 0, 0, 0)
	dc.Clear()
	dc.SetColor(color.Paper)
	dc.DrawRoundedRectangle(0, 0, float64(sizePx.X), float64(sizePx.Y), 8)
	dc.Fill()
	dc.SetColor(color.Ink)
	valueStr := strconv.Itoa(value)
	dc.SetFontFace(getFontFace(valueStr))
	dc.DrawStringAnchored(valueStr, float64(sizePx.X/2), float64(sizePx.Y/2), 0.5, 0.25)
	dc.Fill()
	img := dc.Image()
	return ebiten.NewImageFromImage(img)
}

func getFontFace(value string) font.Face {
	var res font.Face
	l := len(value)
	if l == 1 {
		res = CeilOneCharFF
	} else if l == 2 {
		res = CeilTwoCharFF
	} else if l == 3 {
		res = CeilThreeCharFF
	} else if l == 4 {
		res = CeilFourCharFF
	} else {
		res = CeilFiveCharFF
	}
	return res
}
