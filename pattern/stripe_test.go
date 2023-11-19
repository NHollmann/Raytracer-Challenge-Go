package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

var BLACK = color.New(0, 0, 0)
var WHITE = color.New(1, 1, 1)

func TestStripePatternConstructor(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)

	if !p.A.Equal(WHITE) {
		t.Errorf("pattern A color wrong")
	}
	if !p.B.Equal(BLACK) {
		t.Errorf("pattern A color wrong")
	}
}

func TestStripePatternConstantY(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 1, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 2, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}

func TestStripePatternConstantZ(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 1)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 2)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}

func TestStripePatternAlternateX(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.9, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(-0.1, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(-1, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(-1.1, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}