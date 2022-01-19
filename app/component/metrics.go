package component

import (
	"time"
)

type Metrics struct {
	Ticker    *time.Ticker
	Gpu       string
	Tps       float64
	Fps       float64
	Component int
	Entities  int
	Systems   int
}
