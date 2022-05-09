package fluky

import "github.com/Pencroff/go-toolkit/math"

// Weighted choose random value base on weights
func (f *Fluky) Weighted(values []interface{}, weights []uint) interface{} {
	sum := uint(0)
	mxIdx := math.Min(len(values), len(weights)) - 1
	for idx, weight := range weights {
		if idx > mxIdx {
			break
		}
		sum += weight
	}
	rnd := f.rng.Float64() * float64(sum)
	for idx, weight := range weights {
		rnd -= float64(weight)
		if rnd <= 0 {
			return values[idx]
		}
	}
	return nil
}

// PickOne choose random value from slice and return index and value
func (f Fluky) PickOne(values []interface{}) (idx int, value interface{}) {
	l := uint64(len(values))
	if l == 0 {
		return -1, nil
	}
	r := f.rng.Uint64()
	idx = int(r % l)
	value = values[idx]
	return
}
