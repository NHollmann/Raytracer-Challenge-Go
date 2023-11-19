package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestRingPatternConstructor(t *testing.T) {
	p := pattern.NewRingPattern(WHITE, BLACK)

	if !p.A.Equal(WHITE) {
		t.Errorf("pattern A color wrong")
	}
	if !p.B.Equal(BLACK) {
		t.Errorf("pattern A color wrong")
	}
}

func TestRingPattern(t *testing.T) {
	p := pattern.NewRingPattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 1)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.708, 0, 0.708)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
}
