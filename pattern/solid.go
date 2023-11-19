package pattern

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type SolidPattern struct {
	*BasePattern
	A color.Color
}

func NewSolidPattern(a color.Color) *SolidPattern {
	basePat := newBasePattern()
	gradient := &SolidPattern{
		BasePattern: basePat,
		A:           a,
	}
	basePat.Pattern = gradient
	return gradient
}

func (p *SolidPattern) PatternAt(_ tuple.Tuple) color.Color {
	return p.A
}
