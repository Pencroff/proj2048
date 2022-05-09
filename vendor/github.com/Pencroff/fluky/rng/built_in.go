package rng

import "math/rand"

type BuiltIn struct {
	r *rand.Rand
}

func NewBuiltIn() *BuiltIn {
	r := rand.New(rand.NewSource(11111))
	return &BuiltIn{r: r}
}

func (b *BuiltIn) Seed(v int64) {
	b.r.Seed(v)
}

func (b *BuiltIn) Uint64() uint64 {
	return b.r.Uint64()
}

func (b *BuiltIn) Int63() int64 {
	return b.r.Int63()
}

func (b *BuiltIn) Float64() float64 {
	return b.r.Float64()
}
