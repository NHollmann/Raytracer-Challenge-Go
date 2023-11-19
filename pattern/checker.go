package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type CheckerPattern struct {
	*BasePattern
	A Pattern
	B Pattern
}

func NewCheckerPattern(a, b Pattern) *CheckerPattern {
	basePat := newBasePattern()
	gradient := &CheckerPattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = gradient
	return gradient
}

func NewCheckerPatternColor(a, b color.Color) *CheckerPattern {
	return NewCheckerPattern(NewSolidPattern(a), NewSolidPattern(b))
}

func (p *CheckerPattern) PatternAt(point tuple.Tuple) color.Color {
	sum := math.Floor(point.X()) + math.Floor(point.Y()) + math.Floor(point.Z())
	if int64(math.Floor(sum))%2 == 0 {
		return p.A.subPat(point)
	}
	return p.B.subPat(point)
}
