package control

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
)

type Badge struct {
	Control
	AlignH        HorizontalAlign
	Label         string
	LabelFontFace font.Face
	Value         string
	ValueFontFace font.Face
}

func (b *Badge) Draw() image.Image {
	if b.AlignH == "" {
		b.AlignH = Left
	}
	dc := gg.NewContext(b.Size.X, b.Size.Y)
	dc.SetColor(b.Paper)
	dc.DrawRoundedRectangle(0, 0, float64(b.Size.X), float64(b.Size.Y), float64(b.BorderRadius))
	dc.Fill()
	dc.SetColor(b.Ink)
	dc.SetFontFace(b.LabelFontFace)
	if b.AlignH == Left {
		dc.DrawStringAnchored(b.Label, float64(b.Padding.X), float64(b.Padding.Y), 0, 0.5)
	} else if b.AlignH == Center {
		dc.DrawStringAnchored(b.Label, float64(b.Size.X/2), float64(b.Padding.Y), 0.5, 0.5)
	} else if b.AlignH == Right {
		dc.DrawStringAnchored(b.Label, float64(b.Size.X-b.Padding.X), float64(b.Padding.Y), 1, 0.5)
	}

	dc.SetFontFace(b.ValueFontFace)
	if b.AlignH == Left {
		dc.DrawStringAnchored(b.Value, float64(b.Padding.X), float64(b.Size.Y-b.Padding.Y), 0, 0)
	} else if b.AlignH == Center {
		dc.DrawStringAnchored(b.Value, float64(b.Size.X/2), float64(b.Size.Y-b.Padding.Y), 0.5, 0)
	} else if b.AlignH == Right {
		dc.DrawStringAnchored(b.Value, float64(b.Size.X-b.Padding.X), float64(b.Size.Y-b.Padding.Y), 1, 0)
	}
	return dc.Image()
}
