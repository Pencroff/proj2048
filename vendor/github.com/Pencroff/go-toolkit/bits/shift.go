package bits

// RotateR64 cyclic shift to left unint 64 by "bits" number, accepts positive and negative values
func RotateR64(v uint64, bits int8) (res uint64) {
	n := bits & 63
	res = (v >> n) | (v << (64 - n))
	return
}

// RotateL64 cyclic shift to left uint 64 by "bits" number, accepts positive and negative values
func RotateL64(v uint64, bits int8) (res uint64) {
	n := bits & 63
	res = (v << n) | (v >> (64 - n))
	return
}
