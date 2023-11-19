package pattern_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

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

func TestStripePatternObjectTransform(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)
	mat := matrix.Scaling(2, 2, 2).Inverse()

	if !p.PatternAtTransform(mat, tuple.Point(1.5, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}

func TestStripePatternTransform(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)
	p.SetTransform(matrix.Scaling(2, 2, 2))
	mat := matrix.Identity44().Inverse()

	if !p.PatternAtTransform(mat, tuple.Point(1.5, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}

func TestStripePatternBoothTransform(t *testing.T) {
	p := pattern.NewStripePattern(WHITE, BLACK)
	p.SetTransform(matrix.Translation(0.5, 0, 0))
	mat := matrix.Scaling(2, 2, 2).Inverse()

	if !p.PatternAtTransform(mat, tuple.Point(2.5, 0, 0)).Equal(WHITE) {
		t.Errorf("pattern result color wrong")
	}
}
