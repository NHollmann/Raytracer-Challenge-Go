package pattern

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Pattern interface {
	SetTransform(mat matrix.Mat44)
	PatternAt(point tuple.Tuple) color.Color
	PatternAtTransform(invTransform matrix.Mat44, worldPoint tuple.Tuple) color.Color
	subPat(objectPoint tuple.Tuple) color.Color
}

type BasePattern struct {
	Pattern
	invTransform matrix.Mat44
}

func newBasePattern() *BasePattern {
	return &BasePattern{
		invTransform: matrix.Identity44(),
	}
}

func (p *BasePattern) SetTransform(mat matrix.Mat44) {
	p.invTransform = mat.Inverse()
}

func (p *BasePattern) PatternAtTransform(invTransform matrix.Mat44, worldPoint tuple.Tuple) color.Color {
	objectPoint := invTransform.MulTuple(worldPoint)
	patternPoint := p.invTransform.MulTuple(objectPoint)
	return p.PatternAt(patternPoint)
}

func (p *BasePattern) subPat(objectPoint tuple.Tuple) color.Color {
	patternPoint := p.invTransform.MulTuple(objectPoint)
	return p.PatternAt(patternPoint)
}
