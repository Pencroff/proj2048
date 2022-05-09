package fluky

import "math"

const maxPrecision = 20 + 1

type FloatOptionsFn func(f *FloatOptions)

type FloatOptions struct {
	min       float64
	max       float64
	precision int8
}

// WithFloatRange configure min and max values for float random
func WithFloatRange(min, max float64) FloatOptionsFn {
	return func(f *FloatOptions) {
		f.min = min
		f.max = max
	}
}

// WithPrecision configure max precision for float random
func WithPrecision(precision uint8) FloatOptionsFn {
	return func(f *FloatOptions) {
		f.precision = int8(precision<<1>>1) % maxPrecision
	}
}

// Float random float value from range [min, max)
// by default min = 0, max = 1
// max precision = 20
func (f *Fluky) Float(opts ...FloatOptionsFn) float64 {
	o := &FloatOptions{min: 0, max: 1, precision: -1}
	for _, optFn := range opts {
		optFn(o)
	}

	r := f.rng.Float64()
	r = r*(o.max-o.min) + o.min

	if o.precision > -1 {
		p := math.Pow10(int(o.precision))
		r = math.Round(r*p) / p
	}

	return r
}
