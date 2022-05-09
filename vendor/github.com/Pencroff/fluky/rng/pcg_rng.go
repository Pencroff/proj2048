package rng

import (
	bt "github.com/Pencroff/go-toolkit/bits"
	et "github.com/Pencroff/go-toolkit/extended_type"
)

type PcgRng struct {
	state             et.Uint128
	inc               et.Uint128
	floatMul          float64
	defaultMultiplier et.Uint128
}

func NewPcgRng() *PcgRng {
	v := &PcgRng{
		state:             et.Uint128{0, 0},
		inc:               et.Uint128{0, 0},
		defaultMultiplier: et.New(0x4385df649fccf645, 0x2360ed051fc65da4),
		floatMul:          1 / float64(1<<64-1),
	}
	v.Seed(11111)
	return v
}

func (r *PcgRng) Seed(v int64) {
	initseq := uint64(54)
	r.state = et.ZeroUint128
	r.inc = et.From64(initseq).Lsh(1).Or64(1)
	r.step()
	r.state = r.state.Add64(uint64(v))
	r.step()
}

func (r *PcgRng) Uint64() uint64 {
	r.step()
	return r.stateToValue()
}

func (r *PcgRng) Int63() int64 {
	return int64(r.Uint64() >> 1)
}

func (r *PcgRng) Float64() float64 {
	rnd := r.Uint64() >> (uint64Bits - precisionBits)
	var res float64
	if rnd == maxDoublePrecision {
		rnd -= 1
	}
	res = float64(rnd) / maxDoublePrecision
	return res
}

func (r *PcgRng) step() {
	r.state = r.state.Mul(r.defaultMultiplier).Add(r.inc)
}

func (r *PcgRng) stateToValue() (v uint64) {
	v = r.state.Lo ^ r.state.Hi
	rot := int8(r.state.Rsh(122).Lo & 0x3F)
	v = bt.RotateR64(v, rot)
	return
}
