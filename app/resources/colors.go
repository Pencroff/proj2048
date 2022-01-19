package resources

import (
	log "github.com/go-pkgz/lgr"
	. "github.com/pencroff/proj2048/app/common"
	. "github.com/pencroff/proj2048/app/helper"
)

var BoardStrColor ColorStrPair
var BoardColor ColorPair

var FieldStrColor ColorStrPair
var FieldColor ColorPair

var ColorStrMap map[int]ColorStrPair
var ColorMap map[int]ColorPair

func InitColors() {
	BoardStrColor = ColorStrPair{Ink: "#776E65", Paper: "#FAF8EF"}
	BoardColor = toColors(BoardStrColor)
	FieldStrColor = ColorStrPair{Ink: "#EEE4DA59", Paper: "#BBADA0"}
	FieldColor = toColors(FieldStrColor)
	ColorStrMap = map[int]ColorStrPair{
		2:     {"#776E64", "#EEE4DA"},
		4:     {"#776E64", "#EDE0C8"},
		8:     {"#F9F6F2", "#F2B179"},
		16:    {"#F9F6F2", "#F59563"},
		32:    {"#F9F6F2", "#F67C5F"},
		64:    {"#F9F6F2", "#F65E3B"},
		128:   {"#F9F6F2", "#EDCF72"},
		256:   {"#F9F6F2", "#EDCC61"},
		512:   {"#F9F6F2", "#EDC850"},
		1024:  {"#F9F6F2", "#EDC53F"},
		2048:  {"#F9F6F2", "#EDC22E"},
		4096:  {"#F9F6F2", "#3C3A32"},
		8192:  {"#F2B179", "#3C3A32"},
		16384: {"#EDCC61", "#3C3A32"},
		32768: {"#776E64", "#3C3A32"},
		65536: {"#F67C5F", "#3C3A32"},
	}
	ColorMap = toColorsMap(ColorStrMap)
}

func toColorsMap(v map[int]ColorStrPair) map[int]ColorPair {
	colorMap := map[int]ColorPair{}
	for key, pair := range v {
		colorMap[key] = toColors(pair)
	}
	return colorMap
}

func toColors(pair ColorStrPair) ColorPair {
	ink, err := ColorStrToColor(pair.Ink)
	if err != nil {
		log.Printf("Transform ink color: %v", err)
	}
	paper, err := ColorStrToColor(pair.Paper)
	if err != nil {
		log.Printf("Transform paper color: %v", err)
	}
	return ColorPair{Ink: ink, Paper: paper}
}
