package tuple_test

import (
	"math"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestTupleTypePoint(t *testing.T) {
	a := tuple.New(4.3, -4.2, 3.1, 1.0)
	if !flt.Equal(a.X(), 4.3) {
		t.Errorf("x is not equal 4.3")
	}
	if !flt.Equal(a.Y(), -4.2) {
		t.Errorf("y is not equal -4.2")
	}
	if !flt.Equal(a.Z(), 3.1) {
		t.Errorf("z is not equal 3.1")
	}
	if !flt.Equal(a.W(), 1.0) {
		t.Errorf("w is not equal 1.0")
	}
	if !a.IsPoint() {
		t.Errorf("tuple is not a point")
	}
	if a.IsVector() {
		t.Errorf("tuple is a vector")
	}
}

func TestTupleTypeVector(t *testing.T) {
	a := tuple.New(4.3, -4.2, 3.1, 0.0)
	if !flt.Equal(a.X(), 4.3) {
		t.Errorf("x is not equal 4.3")
	}
	if !flt.Equal(a.Y(), -4.2) {
		t.Errorf("y is not equal -4.2")
	}
	if !flt.Equal(a.Z(), 3.1) {
		t.Errorf("z is not equal 3.1")
	}
	if !flt.Equal(a.W(), 0.0) {
		t.Errorf("w is not equal 0.0")
	}
	if a.IsPoint() {
		t.Errorf("tuple is a point")
	}
	if !a.IsVector() {
		t.Errorf("tuple is not a vector")
	}
}

func TestTuplePointConstructor(t *testing.T) {
	p := tuple.Point(4, -4, 3)
	if !p.Equal(tuple.New(4, -4, 3, 1)) {
		t.Errorf("point constructor defect")
	}
}

func TestTupleVectorConstructor(t *testing.T) {
	p := tuple.Vector(4, -4, 3)
	if !p.Equal(tuple.New(4, -4, 3, 0)) {
		t.Errorf("vector constructor defect")
	}
}

func TestTupleAdd(t *testing.T) {
	a1 := tuple.New(3, -2, 5, 1)
	a2 := tuple.New(-2, 3, 1, 0)
	if !a1.Add(a2).Equal(tuple.New(1, 1, 6, 1)) {
		t.Errorf("tuple add wrong")
	}
}

func TestTupleSubPoints(t *testing.T) {
	p1 := tuple.Point(3, 2, 1)
	p2 := tuple.Point(5, 6, 7)
	if !p1.Sub(p2).Equal(tuple.Vector(-2, -4, -6)) {
		t.Errorf("point sub wrong")
	}
}

func TestTupleSubPointVector(t *testing.T) {
	p := tuple.Point(3, 2, 1)
	v := tuple.Vector(5, 6, 7)
	if !p.Sub(v).Equal(tuple.Point(-2, -4, -6)) {
		t.Errorf("point - vector sub wrong")
	}
}

func TestTupleSubVectors(t *testing.T) {
	v1 := tuple.Vector(3, 2, 1)
	v2 := tuple.Vector(5, 6, 7)
	if !v1.Sub(v2).Equal(tuple.Vector(-2, -4, -6)) {
		t.Errorf("vector sub wrong")
	}
}

func TestTupleSubZeroVector(t *testing.T) {
	v1 := tuple.Vector(0, 0, 0)
	v2 := tuple.Vector(1, -2, 3)
	if !v1.Sub(v2).Equal(tuple.Vector(-1, 2, -3)) {
		t.Errorf("zero vector sub wrong")
	}
}

func TestTupleNeg(t *testing.T) {
	a := tuple.New(1, -2, 3, -4)
	if !a.Neg().Equal(tuple.New(-1, 2, -3, 4)) {
		t.Errorf("negation wrong")
	}
}

func TestTupleMul(t *testing.T) {
	a := tuple.New(1, -2, 3, -4)
	if !a.Mul(3.5).Equal(tuple.New(3.5, -7, 10.5, -14)) {
		t.Errorf("scalar multiply wrong")
	}
}

func TestTupleMulFrac(t *testing.T) {
	a := tuple.New(1, -2, 3, -4)
	if !a.Mul(0.5).Equal(tuple.New(0.5, -1, 1.5, -2)) {
		t.Errorf("fraction scalar multiply wrong")
	}
}

func TestTupleDiv(t *testing.T) {
	a := tuple.New(1, -2, 3, -4)
	if !a.Div(2).Equal(tuple.New(0.5, -1, 1.5, -2)) {
		t.Errorf("scalar division wrong")
	}
}

func TestTupleMagnitude(t *testing.T) {
	cs := []struct {
		t        tuple.Tuple
		expected float64
	}{
		{tuple.Vector(1, 0, 0), 1},
		{tuple.Vector(0, 1, 0), 1},
		{tuple.Vector(0, 0, 1), 1},
		{tuple.Vector(1, 2, 3), math.Sqrt(14)},
		{tuple.Vector(-1, -2, -3), math.Sqrt(14)},
	}
	for _, c := range cs {
		if !flt.Equal(c.t.Magnitude(), c.expected) {
			t.Errorf("magnitude of %+v wrong, expected=%f got=%f", c.t, c.expected, c.t.Magnitude())
		}
	}
}

func TestTupleNormalize(t *testing.T) {
	cs := []struct {
		t        tuple.Tuple
		expected tuple.Tuple
	}{
		{tuple.Vector(4, 0, 0), tuple.Vector(1, 0, 0)},
		{tuple.Vector(0, 4, 0), tuple.Vector(0, 1, 0)},
		{tuple.Vector(0, 0, 4), tuple.Vector(0, 0, 1)},
		{tuple.Vector(1, 2, 3), tuple.Vector(0.26726, 0.53452, 0.80178)},
	}
	for _, c := range cs {
		if !c.t.Normalize().Equal(c.expected) {
			t.Errorf("normalization of %+v wrong, expected=%+v got=%+v", c.t, c.expected, c.t.Normalize())
		}
	}
}

func TestTupleNormalizeUnit(t *testing.T) {
	a := tuple.Vector(1, 2, 3)
	if !flt.Equal(a.Normalize().Magnitude(), 1.0) {
		t.Errorf("normalized vector is not of length 1")
	}
}

func TestTupleDot(t *testing.T) {
	a := tuple.Vector(1, 2, 3)
	b := tuple.Vector(2, 3, 4)
	if !flt.Equal(a.Dot(b), 20.0) {
		t.Errorf("vector dot product wrong")
	}
}

func TestTupleCross(t *testing.T) {
	a := tuple.Vector(1, 2, 3)
	b := tuple.Vector(2, 3, 4)
	if !a.Cross(b).Equal(tuple.Vector(-1, 2, -1)) {
		t.Errorf("vector cross product wrong")
	}
	if !b.Cross(a).Equal(tuple.Vector(1, -2, 1)) {
		t.Errorf("vector cross product wrong")
	}
}

func TestVectorReflect45(t *testing.T) {
	v := tuple.Vector(1, -1, 0)
	n := tuple.Vector(0, 1, 0)
	r := v.Reflect(n)
	if !r.Equal(tuple.Vector(1, 1, 0)) {
		t.Errorf("vector reflection wrong")
	}
}

func TestVectorReflectSlanted(t *testing.T) {
	v := tuple.Vector(0, -1, 0)
	n := tuple.Vector(math.Sqrt(2)/2.0, math.Sqrt(2)/2.0, 0)
	r := v.Reflect(n)
	if !r.Equal(tuple.Vector(1, 0, 0)) {
		t.Errorf("vector reflection wrong")
	}
}
