package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestSolidPatternConstructor(t *testing.T) {
	p := pattern.NewSolidPattern(WHITE)

	if !p.A.Equal(WHITE) {
		t.Errorf("pattern A color wrong")
	}
}

func TestSolidPattern(t *testing.T) {
	p := pattern.NewSolidPattern(WHITE)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 1, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 1)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1, 1, 1)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(-1, -1, -1)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}
