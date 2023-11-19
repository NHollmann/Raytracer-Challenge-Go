package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type StripePattern struct {
	*BasePattern
	A color.Color
	B color.Color
}

func NewStripePattern(a, b color.Color) *StripePattern {
	basePat := newBasePattern()
	stripe := &StripePattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = stripe
	return stripe
}

func (p *StripePattern) PatternAt(point tuple.Tuple) color.Color {
	if int64(math.Floor(point.X()))%2 == 0 {
		return p.A
	}
	return p.B
}
