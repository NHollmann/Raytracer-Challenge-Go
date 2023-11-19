package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type RingPattern struct {
	*BasePattern
	A Pattern
	B Pattern
}

func NewRingPattern(a, b Pattern) *RingPattern {
	basePat := newBasePattern()
	gradient := &RingPattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = gradient
	return gradient
}

func NewRingPatternColor(a, b color.Color) *RingPattern {
	return NewRingPattern(NewSolidPattern(a), NewSolidPattern(b))
}

func (p *RingPattern) PatternAt(point tuple.Tuple) color.Color {
	hypo := math.Sqrt(math.Pow(point.X(), 2) + math.Pow(point.Z(), 2))
	if int64(math.Floor(hypo))%2 == 0 {
		return p.A.subPat(point)
	}
	return p.B.subPat(point)
}
