package fluky

import (
	"github.com/Pencroff/fluky/rng"
)

const (
	maxInt64 = 1<<63 - 1
	minInt64 = -1 << 63
)

// NewFluky create new Fluky instance and return pointer to it
func NewFluky(r rng.RandomGenerator) *Fluky {
	return &Fluky{r}
}

type Fluky struct {
	rng rng.RandomGenerator
}

// Seed internal RNG, reset seed value
func (f *Fluky) Seed(v int64) {
	f.rng.Seed(v)
}

// Uint64 returns random uint64 value
func (f *Fluky) Uint64() uint64 {
	return f.rng.Uint64()
}

// Int63 returns random int64 value in range [0, 2^63)
func (f *Fluky) Int63() int64 {
	return f.rng.Int63()
}

// Float64 returns random float64 value in range [0, 1)
func (f *Fluky) Float64() float64 {
	return f.rng.Float64()
}
