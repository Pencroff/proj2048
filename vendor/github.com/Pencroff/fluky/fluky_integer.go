package fluky

type IntegerOptionsFn func(i *IntegerOptions)

type IntegerOptions struct {
	min int
	max int
}

// WithIntRange configure min and max values for integer random
func WithIntRange(min, max int) IntegerOptionsFn {
	return func(i *IntegerOptions) {
		i.min = min
		i.max = max
	}
}

// Integer random integer value from range [min, max]
// from −(2^63) to 2^63 − 1 by default
func (f *Fluky) Integer(opts ...IntegerOptionsFn) int {
	o := &IntegerOptions{min: minInt64, max: maxInt64}
	for _, optFn := range opts {
		optFn(o)
	}
	if o.max == o.min {
		return o.min
	}
	if o.min != minInt64 && o.max != maxInt64 {
		return int(f.rng.Uint64())%(o.max-o.min) + o.min
	}
	return int(f.rng.Uint64())
}
