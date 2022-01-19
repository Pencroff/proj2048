package resources

import "image"

type ConfigObj struct {
	Speed      int
	Padding    int
	HeaderRect image.Rectangle
	FieldRect  image.Rectangle
	BoardRect  image.Rectangle
	Tile       image.Point
	FieldSpace float64
	FieldSize  image.Point
}

var Config = ConfigObj{
	Speed:      500,
	Padding:    20,
	HeaderRect: image.Rect(0, 0, 500, 90),
	FieldRect:  image.Rect(0, 0, 500, 500),
	BoardRect:  image.Rect(0, 0, 500, 590),
	Tile:       image.Point{X: 100, Y: 100},
	FieldSpace: 0.04,
	FieldSize:  image.Point{X: 4, Y: 4},
}
