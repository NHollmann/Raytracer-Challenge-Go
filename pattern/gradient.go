package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type GradientPattern struct {
	*BasePattern
	A color.Color
	B color.Color
}

func NewGradientPattern(a, b color.Color) *GradientPattern {
	basePat := newBasePattern()
	gradient := &GradientPattern{
		BasePattern: basePat,
		A:           a,
		B:           b,
	}
	basePat.Pattern = gradient
	return gradient
}

func (p *GradientPattern) PatternAt(point tuple.Tuple) color.Color {
	distance := p.B.Sub(p.A)
	fraction := point.X() - math.Floor(point.X())

	return p.A.Add(distance.MulScalar(fraction))
}
