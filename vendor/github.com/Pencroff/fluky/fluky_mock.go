package fluky

import "github.com/stretchr/testify/mock"

type RngMock struct {
	mock.Mock
}

func (r *RngMock) Seed(v int64) {
	r.Called(v)
}

func (r *RngMock) Uint64() uint64 {
	return r.Called().Get(0).(uint64)
}

func (r *RngMock) Int63() int64 {
	return r.Called().Get(0).(int64)
}

func (r *RngMock) Float64() float64 {
	return r.Called().Get(0).(float64)
}
