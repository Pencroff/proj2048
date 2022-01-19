package system

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/control"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/sedyh/mizu/pkg/engine"
	"image"
)

type GameOverSystem struct {
	*component.BoardProp
	btn     *control.Button
	btnRect *image.Rectangle
}

func (s *GameOverSystem) Update(_ engine.World) {
	p := image.Pt(ebiten.CursorPosition())
	isClickable := s.State == common.GameEnd
	if isClickable && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.In(*s.btnRect) {
		s.State = common.RestartGame
	}
}

func (s *GameOverSystem) Draw(_ engine.World, screen *ebiten.Image) {
	if s.btn == nil {
		s.btn = s.GetButton()
		width, height := s.FieldProps.Sprite.Size()
		x1, y1 := s.FieldProps.Position.X+(width-s.btn.Size.X)/2, s.FieldProps.Position.Y+(height/2)+s.btn.Size.Y/2
		rect := image.Rect(x1, y1, x1+s.btn.Size.X, y1+s.btn.Size.Y)
		s.btnRect = &rect
	}
	if s.IsFinished {
		width, height := s.FieldProps.Sprite.Size()
		dc := gg.NewContext(width, height)
		dc.SetHexColor("#EEE4DABA")
		dc.Clear()
		dc.SetFontFace(resources.CeilOneCharFF)
		dc.SetColor(resources.BoardColor.Ink)
		dc.DrawStringAnchored("Game Over!", float64(width/2), float64(height/2), 0.5, 0)

		s.btn.DrawOnContext(dc, (width-s.btn.Size.X)/2, (height/2)+s.btn.Size.Y/2)

		img := ebiten.NewImageFromImage(dc.Image())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(s.FieldProps.Position.X), float64(s.FieldProps.Position.Y))
		screen.DrawImage(img, op)
	}
}

func (s *GameOverSystem) GetButton() *control.Button {
	return &control.Button{
		Control: control.Control{
			BorderRadius: 4,
			Padding:      image.Point{X: 8, Y: 8},
			Size:         image.Point{X: 128, Y: 64},
			Ink:          resources.BoardColor.Paper,
			Paper:        resources.BoardColor.Ink,
		},
		AlignH:       control.Center,
		Text:         "Try again",
		TextFontFace: resources.TitleFF,
	}
}
