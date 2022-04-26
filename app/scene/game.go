package scene

import (
	"github.com/Pencroff/fluky"
	"github.com/Pencroff/fluky/rng"
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/entity"
	"github.com/pencroff/proj2048/app/helper"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/pencroff/proj2048/app/system"
	"github.com/sedyh/mizu/pkg/engine"
	"image"
	"time"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Metrics{},
		component.TilePropWrap{},
		component.FieldProps{},
		component.BoardProp{},
	)

	w.AddSystems(
		&system.Remover{},
		&system.BackgroundSystem{},
		&system.Board{},
		&system.RenderField{},
		&system.RenderTile{},
		&system.GameOverSystem{},
		&system.Metrics{},
	)

	fieldPtr, boardPtr := NewBoard()

	w.AddEntities(
		&entity.Metrics{
			component.Metrics{
				Ticker: time.NewTicker(500 * time.Millisecond),
				Gpu:    helper.GpuInfo(),
			},
		},
		fieldPtr,
		boardPtr,
	)
}

func NewBoard() (field *entity.Field, board *entity.Board) {
	size := resources.Config.FieldSize
	headerOffset := resources.Config.HeaderRect.Size()
	headerOffset.X = 0
	fieldRect := resources.Config.FieldRect.Add(headerOffset)
	boardRect := resources.Config.BoardRect
	field = NewField(size, fieldRect, resources.FieldColor)
	agnt := resources.PoolAgentInstance
	board = &entity.Board{
		component.BoardProp{
			Name:        "2048",
			Description: agnt.GetName(),
			Score:       0,
			Step:        0,
			NoMove:      false,
			Speed:       resources.Config.Speed,
			State:       common.StartGame,
			Color:       resources.BoardColor,
			Size:        size,
			BoardRect:   boardRect,
			Direction:   common.NoDirection,
			FieldProps:  &field.FieldProps,
			List:        make([]*component.TileProp, size.X*size.Y),
			Agent:       resources.HumanAgentInstance,
			Mode:        common.Manual,
			Flk:         fluky.NewFluky(rng.NewSmallPrng()),
		},
	}
	return
}

func NewField(size image.Point, rect image.Rectangle, color common.ColorPair) (field *entity.Field) {
	field = &entity.Field{
		component.FieldProps{
			Position: rect.Min,
			Sprite:   resources.RenderField(size, rect, color),
		},
	}
	return
}
