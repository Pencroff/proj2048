package helper

import "image"

func ManhattanDistance(p1, p2 image.Point) int {
	p := p1.Sub(p2)
	dist := AbsInt(p.X) + AbsInt(p.Y)
	return dist
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
