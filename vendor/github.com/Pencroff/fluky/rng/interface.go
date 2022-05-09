package rng

const uint64Bits = 64
const precisionBits = 53
const maxDoublePrecision = 1<<precisionBits - 1

type RandomGenerator interface {
	Seed(v int64)
	Uint64() uint64
	Int63() int64
	Float64() float64
}
