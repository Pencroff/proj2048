package system

import (
	"github.com/pencroff/ai-agents/proj2048/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Remover struct{}

func (d *Remover) Update(w engine.World) {
	view := w.View(component.TilePropWrap{})
	for _, e := range view.Filter() {
		var prop *component.TilePropWrap
		e.Get(&prop)

		if prop.Ptr.Removed {
			w.RemoveEntity(e)
		}
	}
}
