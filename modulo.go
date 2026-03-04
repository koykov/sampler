package sampler

import (
	"math"
	"sync/atomic"
)

type modulo struct {
	base, c uint64
}

func NewModulo(base uint64) Interface {
	return &modulo{
		base: base,
		c:    math.MaxUint64,
	}
}

func (s modulo) Sample() bool {
	if s.base == 0 {
		return true
	}
	return atomic.AddUint64(&s.c, 1)%s.base == 0
}
