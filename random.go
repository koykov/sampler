package sampler

type RNG interface {
	Seed(int64)
	Float64() float64
}

type random struct {
	rate float64
	rng  RNG
}

func NewRandom(rate float64, rng RNG) Interface {
	return &random{
		rate: rate,
		rng:  rng,
	}
}

func (s *random) Sample() bool {
	if s.rate <= 0 || s.rng == nil {
		return true
	}
	return s.rng.Float64() < s.rate
}
