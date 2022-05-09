package rng

import (
	"github.com/Pencroff/go-toolkit/bits"
)

// https://burtleburtle.net/bob/rand/smallprng.html
//
//typedef unsigned long long u8;
//typedef struct ranctx { u8 a; u8 b; u8 c; u8 d; } ranctx;
//
//#define rot(x,k) (((x)<<(k))|((x)>>(64-(k))))
//u8 ranval( ranctx *x ) {
//    u8 e = x->a - rot(x->b, 7);
//    x->a = x->b ^ rot(x->c, 13);
//    x->b = x->c + rot(x->d, 37);
//    x->c = x->d + e;
//    x->d = e + x->a;
//    return x->d;
//}
//
//void raninit( ranctx *x, u8 seed ) {
//    u8 i;
//    x->a = 0xf1ea5eed, x->b = x->c = x->d = seed;
//    for (i=0; i<20; ++i) {
//        (void)ranval(x);
//    }
//}

type SmallPrng struct {
	a        uint64
	b        uint64
	c        uint64
	d        uint64
	floatMul float64
}

func NewSmallPrng() *SmallPrng {
	seed := uint64(11111)
	return &SmallPrng{
		a:        0xf1ea5eed,
		b:        seed,
		c:        seed,
		d:        seed,
		floatMul: 1 / float64(1<<64-1),
	}
}

func (s *SmallPrng) Seed(v int64) {
	s.a = 0xf1ea5eed
	s.b = uint64(v)
	s.c = uint64(v)
	s.d = uint64(v)
}

func (s *SmallPrng) Uint64() uint64 {
	e := s.a - bits.RotateL64(s.b, 7)
	s.a = s.b ^ bits.RotateL64(s.c, 13)
	s.b = s.c + bits.RotateL64(s.d, 37)
	s.c = s.d + e
	s.d = e + s.a
	return s.d
}

func (s *SmallPrng) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *SmallPrng) Float64() float64 {
	rnd := s.Uint64() >> (uint64Bits - precisionBits)
	var res float64
	if rnd == maxDoublePrecision {
		rnd -= 1
	}
	res = float64(rnd) / maxDoublePrecision
	return res
}
