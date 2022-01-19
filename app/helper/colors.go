package helper

import (
	"github.com/pkg/errors"
	"image/color"
)

var errInvalidColorFormat = errors.New("invalid color format, expected #RGB / #RGBA")

func ColorStrToColor(s string) (color.Color, error) {
	var err error
	rgba := color.RGBA{A: 0xff}

	if s[0] != '#' {
		return rgba, errInvalidColorFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidColorFormat
		return 0
	}

	switch len(s) {
	case 9:
		rgba.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		rgba.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		rgba.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		rgba.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		rgba.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		rgba.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		rgba.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 5:
		rgba.R = hexToByte(s[1]) * 17
		rgba.G = hexToByte(s[2]) * 17
		rgba.B = hexToByte(s[3]) * 17
		rgba.A = hexToByte(s[4]) * 17
	case 4:
		rgba.R = hexToByte(s[1]) * 17
		rgba.G = hexToByte(s[2]) * 17
		rgba.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidColorFormat
	}
	return rgba, err
}
