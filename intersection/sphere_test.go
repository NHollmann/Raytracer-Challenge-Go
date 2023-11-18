package intersection_test

import (
	"math"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/material"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestSphereDefaultTransform(t *testing.T) {
	s := intersection.NewSphere()

	if !matrix.Identity44().Equal(s.Transform) {
		t.Errorf("sphere has wrong default transform")
	}
}

func TestSphereDefaultMaterial(t *testing.T) {
	s := intersection.NewSphere()

	if s.Material != material.New() {
		t.Errorf("sphere has wrong default material")
	}
}

func TestSphereIntersection2Points(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("sphere intersect count not 2")
	}
	if !flt.Equal(xs[0].T, 4.0) {
		t.Errorf("first sphere intersection not at 4.0")
	}
	if xs[0].Object != s {
		t.Errorf("first sphere intersection has wrong object")
	}
	if !flt.Equal(xs[1].T, 6.0) {
		t.Errorf("second sphere intersection not at 6.0")
	}
	if xs[1].Object != s {
		t.Errorf("second sphere intersection has wrong object")
	}
}

func TestSphereIntersection1Point(t *testing.T) {
	r := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()

	xs := s.Intersect(r)

	// Intersections at a tangent also return two results for simplicity
	if len(xs) != 2 {
		t.Fatalf("sphere intersect count not 2")
	}
	if !flt.Equal(xs[0].T, 5.0) {
		t.Errorf("first sphere intersection not at 5.0")
	}
	if !flt.Equal(xs[1].T, 5.0) {
		t.Errorf("second sphere intersection not at 5.0")
	}
}

func TestSphereIntersectionMiss(t *testing.T) {
	r := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("sphere intersect count not 0")
	}
}

func TestSphereIntersectionFromInside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("sphere intersect count not 2")
	}
	if !flt.Equal(xs[0].T, -1.0) {
		t.Errorf("first sphere intersection not at -1.0")
	}
	if !flt.Equal(xs[1].T, 1.0) {
		t.Errorf("second sphere intersection not at 1.0")
	}
}

func TestSphereIntersectionFromBehind(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("sphere intersect count not 2")
	}
	if !flt.Equal(xs[0].T, -6.0) {
		t.Errorf("first sphere intersection not at -6.0")
	}
	if !flt.Equal(xs[1].T, -4.0) {
		t.Errorf("second sphere intersection not at -4.0")
	}
}

func TestSphereIntersectionScaled(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()
	s.Transform = matrix.Scaling(2, 2, 2)

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Fatalf("sphere intersect count not 2")
	}
	if !flt.Equal(xs[0].T, 3.0) {
		t.Errorf("first sphere intersection not at 3.0")
	}
	if !flt.Equal(xs[1].T, 7.0) {
		t.Errorf("second sphere intersection not at 7.0")
	}
}

func TestSphereIntersectionTranslated(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := intersection.NewSphere()
	s.Transform = matrix.Translation(5, 0, 0)

	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("sphere intersect count not 0")
	}
}

func TestSphereNormalAtX(t *testing.T) {
	s := intersection.NewSphere()
	n := s.NormalAt(tuple.Point(1, 0, 0))

	if !tuple.Vector(1, 0, 0).Equal(n) {
		t.Fatalf("normal wrong")
	}
}

func TestSphereNormalAtY(t *testing.T) {
	s := intersection.NewSphere()
	n := s.NormalAt(tuple.Point(0, 1, 0))

	if !tuple.Vector(0, 1, 0).Equal(n) {
		t.Fatalf("normal wrong")
	}
}

func TestSphereNormalAtZ(t *testing.T) {
	s := intersection.NewSphere()
	n := s.NormalAt(tuple.Point(0, 0, 1))

	if !tuple.Vector(0, 0, 1).Equal(n) {
		t.Fatalf("normal wrong")
	}
}

func TestSphereNormalAtDiagonal(t *testing.T) {
	s := intersection.NewSphere()
	x := math.Sqrt(3) / 3.0
	n := s.NormalAt(tuple.Point(x, x, x))

	if !tuple.Vector(x, x, x).Equal(n) {
		t.Fatalf("normal wrong")
	}
}

func TestSphereNormalAtIsNormalized(t *testing.T) {
	s := intersection.NewSphere()
	x := math.Sqrt(3) / 3.0
	n := s.NormalAt(tuple.Point(x, x, x))

	if !n.Normalize().Equal(n) {
		t.Fatalf("normal is not normalized")
	}
}

func TestSphereNormalAtTranslated(t *testing.T) {
	s := intersection.NewSphere()
	s.Transform = matrix.Translation(0, 1, 0)
	n := s.NormalAt(tuple.Point(0, 1.70711, -0.70711))

	if !tuple.Vector(0, 0.70711, -0.70711).Equal(n) {
		t.Fatalf("normal wrong")
	}
}

func TestSphereNormalAtTransformed(t *testing.T) {
	s := intersection.NewSphere()
	s.Transform = matrix.Scaling(1, 0.5, 1).Mul(matrix.RotationZ(math.Pi / 5.0))
	n := s.NormalAt(tuple.Point(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0))

	if !tuple.Vector(0, 0.97014, -0.24254).Equal(n) {
		t.Fatalf("normal wrong")
	}
}
