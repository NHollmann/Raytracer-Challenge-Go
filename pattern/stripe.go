package pattern

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type StripePattern struct {
	A            color.Color
	B            color.Color
	invTransform matrix.Mat44
}

func NewStripePattern(a, b color.Color) *StripePattern {
	return &StripePattern{
		A:            a,
		B:            b,
		invTransform: matrix.Identity44(),
	}
}

func (p *StripePattern) SetTransform(mat matrix.Mat44) {
	p.invTransform = mat.Inverse()
}

func (p *StripePattern) PatternAt(point tuple.Tuple) color.Color {
	if int64(math.Floor(point.X()))%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *StripePattern) PatternAtTransform(invTransform matrix.Mat44, worldPoint tuple.Tuple) color.Color {
	objectPoint := invTransform.MulTuple(worldPoint)
	patternPoint := p.invTransform.MulTuple(objectPoint)
	return p.PatternAt(patternPoint)
}
