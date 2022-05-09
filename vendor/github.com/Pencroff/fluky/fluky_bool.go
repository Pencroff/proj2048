package fluky

type BoolOptionsFn func(b *BoolOptions)

type BoolOptions struct {
	likelihood float64
}

// WithLikelihood makes closure with passed likelihood for options object
func WithLikelihood(v float64) BoolOptionsFn {
	return func(b *BoolOptions) {
		b.likelihood = v
	}
}

// Bool returns random bool value with 0.5 likelihood by default
// Options changing likelihood of returned value
func (f *Fluky) Bool(opts ...BoolOptionsFn) bool {
	b := &BoolOptions{likelihood: 0.5}
	for _, optFn := range opts {
		optFn(b)
	}
	return f.rng.Float64() < b.likelihood
}
