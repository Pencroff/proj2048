package control

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
)

type Button struct {
	Control
	AlignH       HorizontalAlign
	Text         string
	TextFontFace font.Face
}

func (b *Button) Draw() image.Image {
	dc := gg.NewContext(b.Size.X, b.Size.Y)
	dc.SetColor(b.Paper)
	dc.DrawRoundedRectangle(0, 0, float64(b.Size.X), float64(b.Size.Y), float64(b.BorderRadius))
	dc.Fill()
	dc.SetColor(b.Ink)
	dc.SetFontFace(b.TextFontFace)
	if b.AlignH == Left {
		dc.DrawStringAnchored(b.Text, float64(b.Padding.X), float64(b.Size.Y/2), 0, 0.25)
	} else if b.AlignH == Center {
		dc.DrawStringAnchored(b.Text, float64(b.Size.X/2), float64(b.Size.Y/2), 0.5, 0.25)
	} else if b.AlignH == Right {
		dc.DrawStringAnchored(b.Text, float64(b.Size.X-b.Padding.X), float64(b.Size.Y/2), 1, 0.25)
	}
	//dc.SetColor(colornames.Red)
	//dc.DrawPoint(float64(b.Size.X/2), float64(b.Size.Y/2), 2)
	//dc.Fill()
	return dc.Image()
}

func (b *Button) DrawOnContext(ctx *gg.Context, x, y int) {
	img := b.Draw()
	ctx.DrawImage(img, x, y)
}
