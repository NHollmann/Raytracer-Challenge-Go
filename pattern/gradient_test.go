package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestGradientPatternX(t *testing.T) {
	p := pattern.NewGradientPatternColor(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.25, 0, 0)).Equal(color.New(0.75, 0.75, 0.75)) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.5, 0, 0)).Equal(color.New(0.5, 0.5, 0.5)) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.75, 0, 0)).Equal(color.New(0.25, 0.25, 0.25)) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.99999999, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}
