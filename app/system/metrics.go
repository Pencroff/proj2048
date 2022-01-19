package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pencroff/ai-agents/proj2048/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Metrics struct {
	*component.Metrics
}

func (m *Metrics) Update(w engine.World) {
	select {
	case <-m.Ticker.C:
		m.Component = w.Components()
		m.Entities = w.Entities()
		m.Systems = w.Systems()
		m.Tps = ebiten.CurrentTPS()
		m.Fps = ebiten.CurrentFPS()
	default:
	}
}

func (m *Metrics) Draw(w engine.World, screen *ebiten.Image) {
	str := fmt.Sprintf(
		"%s | %dx%d | TPS: %2.f | FPS: %3.f | Objects: %d:%d:%d",
		m.Gpu, w.Bounds().Dx(), w.Bounds().Dy(),
		m.Tps, m.Fps, m.Component, m.Entities, m.Systems,
	)

	ebitenutil.DebugPrintAt(screen, str, 10, 590)
}
