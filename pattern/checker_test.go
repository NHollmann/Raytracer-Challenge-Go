package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestCheckerPatternConstructor(t *testing.T) {
	p := pattern.NewCheckerPattern(WHITE, BLACK)

	if !p.A.Equal(WHITE) {
		t.Errorf("pattern A color wrong")
	}
	if !p.B.Equal(BLACK) {
		t.Errorf("pattern A color wrong")
	}
}

func TestCheckerPatternX(t *testing.T) {
	p := pattern.NewCheckerPattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0.99, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(1.01, 0, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
}

func TestCheckerPatternY(t *testing.T) {
	p := pattern.NewCheckerPattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0.99, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 1.01, 0)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
}

func TestCheckerPatternZ(t *testing.T) {
	p := pattern.NewCheckerPattern(WHITE, BLACK)

	if !p.PatternAt(tuple.Point(0, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 0.99)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
	if !p.PatternAt(tuple.Point(0, 0, 1.01)).Equal(BLACK) {
		t.Errorf("pattern result color wrong")
	}
}
