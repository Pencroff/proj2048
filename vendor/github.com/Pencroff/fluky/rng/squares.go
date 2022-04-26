package rng

import "github.com/Pencroff/go-toolkit/bits"

//inline static uint64_t squares64(uint64_t ctr, uint64_t key) {
//uint64_t t, x, y, z;
//y = x = ctr * key; z = y + key;
//x = x*x + y; x = (x>>32) | (x<<32); /* round 1 */
//x = x*x + z; x = (x>>32) | (x<<32); /* round 2 */
//x = x*x + y; x = (x>>32) | (x<<32); /* round 3 */
//t = x = x*x + z; x = (x>>32) | (x<<32); /* round 4 */
//return t ^ ((x*x + y) >> 32); /* round 5 */
//}

type Squares struct {
	t        uint64
	x        uint64
	y        uint64
	z        uint64
	key      uint64
	seed     uint64
	floatMul float64
}

func NewSquares() *Squares {
	return &Squares{
		key:      0x548c9decbce65297,
		seed:     11111,
		floatMul: 1 / float64(1<<64-1),
	}
}

func (s *Squares) Seed(v int64) {
	s.seed = uint64(v)
}

func (s *Squares) Uint64() uint64 {
	v := s.seed * s.key
	s.x = v
	s.y = v
	s.z = s.y + s.key
	s.x = s.x*s.x + s.y
	s.x = bits.RotateL64(s.x, 32)
	s.x = s.x*s.x + s.z
	s.x = bits.RotateL64(s.x, 32)
	s.x = s.x*s.x + s.y
	s.x = bits.RotateL64(s.x, 32)
	s.x = s.x*s.x + s.z
	s.x = bits.RotateL64(s.x, 32)
	s.t = s.x
	s.seed = s.t ^ ((s.x*s.x + s.y) >> 32)
	return s.seed
}

func (s *Squares) Float64() float64 {
	rnd := s.Uint64()
	return float64(rnd) * s.floatMul
}
