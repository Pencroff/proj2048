package rng

type RandomGenerator interface {
	Seed(v int64)
	Uint64() uint64
	Float64() float64
}
