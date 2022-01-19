package helper

import (
	"encoding/binary"
	"fmt"
	"github.com/OneOfOne/xxhash"
	"github.com/fogleman/gg"
	"github.com/spaolacci/murmur3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_murmur_hash(t *testing.T) {
	tbl := []struct {
		in  string
		out uint32
	}{
		{"", 0x0},
		{"hello", 0x248BFA47},
		{"hello world", 0x5E928F0F},
		{"1234567890", 0x3204634D},
	}

	h32 := murmur3.New32()

	for n, elem := range tbl {
		h32.Write([]byte(elem.in))
		assert.Equal(t, elem.out, h32.Sum32(), "check #%d - %v", n, elem.in)
		h32.Reset()
	}
}

func Test_zero_input(t *testing.T) {
	m64 := murmur3.New64()
	x64 := xxhash.New64()
	buff := make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, 0)
	fmt.Println(buff)
	m64.Write(buff)
	x64.Write(buff)
	assert.Equal(t, uint64(0x28df63b7cc57c3cb), m64.Sum64(), "murmur incorrect")
	assert.Equal(t, uint64(0x34c96acdcadb1bbb), x64.Sum64(), "xxhash incorrect")
	fmt.Printf("%x\n", m64.Sum64())
	fmt.Printf("%x\n", x64.Sum64())
}

func Test_randomness(t *testing.T) {
	size := 1024
	n := uint64(size * size)
	hasher64 := murmur3.New64()
	xxhash64 := xxhash.New64()
	dc := gg.NewContext(size*2+2, size)
	var i uint64
	for i = 0; i <= n; i += 1 {
		buff := make([]byte, 8)
		binary.LittleEndian.PutUint64(buff, i)
		hasher64.Write(buff)
		xxhash64.Write(buff)
		v64 := hasher64.Sum64()
		x1, y1, x2, y2, r1, g1, b1, r2, g2, b2 := extract(v64)

		dc.SetRGB255(int(r1), int(g1), int(b1))
		dc.SetPixel(int(x1), int(y1))
		dc.SetRGB255(int(r2), int(g2), int(b2))
		dc.SetPixel(int(x2), int(y2))

		v64 = xxhash64.Sum64()
		x1, y1, x2, y2, r1, g1, b1, r2, g2, b2 = extract(v64)

		dc.SetRGB255(int(r1), int(g1), int(b1))
		dc.SetPixel(int(x1)+size+2, int(y1))
		dc.SetRGB255(int(r2), int(g2), int(b2))
		dc.SetPixel(int(x2)+size+2, int(y2))
		hasher64.Reset()
		xxhash64.Reset()
	}
	dc.SavePNG("out.png")

	// The 32 bit version
	//x := helper.Extract(uint64(v32), 0, 10)
	//y := helper.Extract(uint64(v32), 10, 10)
	//r := helper.Extract(uint64(v32), 20, 4)
	//g := helper.Extract(uint64(v32), 24, 4)
	//b := helper.Extract(uint64(v32), 28, 4)
	//dc.SetRGB255(int(r)*16, int(g)*16, int(b)*16)
}

func extract(v uint64) (x1, y1, x2, y2, r1, g1, b1, r2, g2, b2 uint64) {
	x1 = Extract(v, 0, 10)
	y1 = Extract(v, 10, 10)
	x2 = Extract(v, 20, 10)
	y2 = Extract(v, 30, 10)
	r1 = Extract(v, 40, 4) * 8
	g1 = Extract(v, 44, 4) * 8
	b1 = Extract(v, 48, 4) * 8
	r2 = Extract(v, 52, 4) * 8
	g2 = Extract(v, 56, 4) * 8
	b2 = Extract(v, 60, 4) * 8
	return
}
