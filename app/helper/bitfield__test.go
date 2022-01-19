package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extract(t *testing.T) {
	tbl := []struct {
		in    uint64
		start int
		len   int
		out   uint64
	}{
		{
			in:    0b11010010011100011100000001111111,
			start: 0,
			len:   10,
			out:   0b0001111111,
		}, {
			in:    0b11010010011100011100000001111111,
			start: 10,
			len:   10,
			out:   0b0001110000,
		}, {
			in:    0b11010010011100011100000001111111,
			start: 20,
			len:   4,
			out:   0b0111,
		}, {
			in:    0b11010010011100011100000001111111,
			start: 24,
			len:   4,
			out:   0b0010,
		}, {
			in:    0b11010010011100011100000001111111,
			start: 28,
			len:   4,
			out:   0b1101,
		}, {
			in:    0b11010010011100011100000001111111,
			start: 2,
			len:   20,
			out:   0b11000111000000011111,
		},
	}

	for _, el := range tbl {
		res := Extract(el.in, el.start, el.len)
		assert.Equal(t, el.out, res, "incorrect extract value: %b", res)
	}
}
