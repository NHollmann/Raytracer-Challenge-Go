package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type GradientPattern struct {
	*BasePattern
	A Pattern
	B Pattern
}

func NewGradientPattern(a, b Pattern) *GradientPattern {
	basePat := newBasePattern()
	gradient := &GradientPattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = gradient
	return gradient
}

func NewGradientPatternColor(a, b color.Color) *GradientPattern {
	return NewGradientPattern(NewSolidPattern(a), NewSolidPattern(b))
}

func (p *GradientPattern) PatternAt(point tuple.Tuple) color.Color {
	colA := p.A.subPat(point)
	distance := p.B.subPat(point).Sub(colA)
	fraction := point.X() - math.Floor(point.X())

	return colA.Add(distance.MulScalar(fraction))
}
