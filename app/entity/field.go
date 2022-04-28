package entity

import (
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/resources"
	"image"
)

type Field struct {
	component.FieldProps
}

func NewField(size image.Point, rect image.Rectangle, color common.ColorPair) (field *Field) {
	field = &Field{
		component.FieldProps{
			Position: rect.Min,
			Sprite:   resources.RenderField(size, rect, color),
		},
	}
	return
}
