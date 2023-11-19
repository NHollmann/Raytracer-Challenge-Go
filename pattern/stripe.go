package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type StripePattern struct {
	*BasePattern
	A Pattern
	B Pattern
}

func NewStripePattern(a, b Pattern) *StripePattern {
	basePat := newBasePattern()
	stripe := &StripePattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = stripe
	return stripe
}

func NewStripePatternColor(a, b color.Color) *StripePattern {
	return NewStripePattern(NewSolidPattern(a), NewSolidPattern(b))
}

func (p *StripePattern) PatternAt(point tuple.Tuple) color.Color {
	if int64(math.Floor(point.X()))%2 == 0 {
		return p.A.subPat(point)
	}
	return p.B.subPat(point)
}
