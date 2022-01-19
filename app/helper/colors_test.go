package helper

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func Test_transform_to_color(t *testing.T) {
	tbl := []struct {
		inp string
		out color.Color
	}{
		{"#FF0000", color.RGBA{R: 0xFF, G: 0, B: 0, A: 0}},
		{"#00FF00", color.RGBA{R: 0, G: 0xFF, B: 0, A: 0}},
		{"#0000FF", color.RGBA{R: 0, G: 0, B: 0xFF, A: 0}},
		{"#000000FF", color.RGBA{R: 0, G: 0, B: 0, A: 0xFF}},
		{"#ff0000", color.RGBA{R: 0xFF, G: 0, B: 0, A: 0}},
		{"#00ff00", color.RGBA{R: 0, G: 0xFF, B: 0, A: 0}},
		{"#0000ff", color.RGBA{R: 0, G: 0, B: 0xFF, A: 0}},
		{"#abcdef12", color.RGBA{R: 0xab, G: 0xcd, B: 0xef, A: 0x12}},
	}
	for n, tt := range tbl {
		c, err := ColorStrToColor(tt.inp)
		assert.Equal(t, tt.out, c, "check #%d", n)
		assert.Nil(t, err)
	}
}

func Test_transform_to_color_err(t *testing.T) {
	tbl := []string{
		"#FF0000zz",
		"#00FF0",
		"#0000zz",
		"#0000ooFF",
		"00ff00",
		"#g000ff",
		"#abcdefss",
	}
	for n, inp := range tbl {
		_, err := ColorStrToColor(inp)
		assert.Errorf(t, err, "invalid color format, expected #RGB / #RGBA", "check #%d", n)
	}
}
