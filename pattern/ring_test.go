package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestRingPattern(t *testing.T) {
	p := pattern.NewRingPatternColor(WHITE, BLACK)

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
