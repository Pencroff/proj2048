package scene

import (
	"github.com/jmoiron/sqlx"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/entity"
	"github.com/pencroff/proj2048/app/helper"
	"github.com/pencroff/proj2048/app/stats"
	"github.com/pencroff/proj2048/app/system"
	"github.com/sedyh/mizu/pkg/engine"
	"time"
)

type Game struct {
	db *sqlx.DB
}

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

	recorder := stats.NewStatRecorder(g.db)

	fieldPtr, boardPtr := entity.NewBoard(recorder)

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

func (g *Game) SetDb(db *sqlx.DB) {
	g.db = db
}
