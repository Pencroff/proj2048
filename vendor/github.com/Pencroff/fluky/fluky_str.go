package fluky

import (
	"strings"
)

const (
	lower       = "abcdefghijklmnopqrstuvwxyz"
	upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers     = "0123456789"
	symbols     = "!@#$%^&*+="
	linkSymbols = "_-"
)

type StringOptionsFn func(s *StringOptions)

type StringOptions struct {
	minLen   uint8
	maxLen   uint8
	alphabet string
}

// WithLength makes closure with passed length for options object
func WithLength(v uint8) StringOptionsFn {
	return func(s *StringOptions) {
		s.minLen = v
		s.maxLen = v
	}
}

// WithLengthRange makes closure with passed min and max length for options object
func WithLengthRange(min, max uint8) StringOptionsFn {
	return func(s *StringOptions) {
		s.minLen = min
		s.maxLen = max
	}
}

// WithAlphabet makes closure with passed alphabet for options object
func WithAlphabet(alphabet string) StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = alphabet
	}
}

// AndAlphabet extend configured alphabet for options object
func AndAlphabet(alphabet string) StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += alphabet
	}
}

// WithUrlSafeAlphabet configure safe for url usage alphabet
func WithUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = lower + upper + numbers + linkSymbols
	}
}

// WithHexAlphabet configure hex alphabet
func WithHexAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = numbers + "abcdef"
	}
}

// WithNumericAlphabet configure numeric alphabet
func WithNumericAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = numbers
	}
}

// AndNumericAlphabet extend configured alphabet with numbers
func AndNumericAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += numbers
	}
}

// WithLowerAlphabet configure lower alphabet
func WithLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = lower
	}
}

// AndLowerAlphabet extend configured alphabet with lower
func AndLowerAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += lower
	}
}

// WithUpperAlphabet configure upper alphabet
func WithUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = upper
	}
}

// AndUpperAlphabet extend configured alphabet with upper
func AndUpperAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += upper
	}
}

// WithSymbolsAlphabet configure symbols alphabet
func WithSymbolsAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = symbols + linkSymbols
	}
}

// AndSymbolsAlphabet extend configured alphabet with symbols
func AndSymbolsAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += symbols + linkSymbols
	}
}

// WithSymbolsUrlSafeAlphabet configure safe for url usage alphabet symbols
func WithSymbolsUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet = linkSymbols
	}
}

// AndSymbolsUrlSafeAlphabet extend configured alphabet with symbols
func AndSymbolsUrlSafeAlphabet() StringOptionsFn {
	return func(s *StringOptions) {
		s.alphabet += linkSymbols
	}
}

// String returns random string configured by default options
func (f *Fluky) String(opts ...StringOptionsFn) string {
	b := &StringOptions{minLen: 5, maxLen: 20, alphabet: lower + upper + numbers + symbols + linkSymbols}
	for _, opt := range opts {
		opt(b)
	}
	l := f.Integer(WithIntRange(int(b.minLen), int(b.maxLen)))
	alphabetRunes := []rune(b.alphabet)
	maxIdx := int64(len(alphabetRunes) - 1)
	builder := strings.Builder{}
	for i := 0; i < l; i++ {
		idx := maxIdx
		if maxIdx != 0 {
			idx = f.Int63() % maxIdx
		}
		builder.WriteRune(alphabetRunes[idx])
	}
	return builder.String()
}
