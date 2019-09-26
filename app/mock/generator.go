package mock

import (
	"math/rand"
	"time"
)

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func RandomBool() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func TernaryString(statement bool, a, b string) string {
	if statement {
		return a
	}
	return b
}

func NormalDistributionFactor() float64 {
	res := rand.NormFloat64()/3 + 0.5
	if res > 1.0 || res < 0 {
		res = NormalDistributionFactor()
	}
	return res
}
