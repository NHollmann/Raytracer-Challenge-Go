package intersection_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestPlaneNormalAt(t *testing.T) {
	p := intersection.NewPlane()
	n1 := p.NormalAt(tuple.Point(0, 0, 0))
	n2 := p.NormalAt(tuple.Point(10, 0, -10))
	n3 := p.NormalAt(tuple.Point(-5, 0, 150))

	if !tuple.Vector(0, 1, 0).Equal(n1) {
		t.Fatalf("normal wrong")
	}
	if !tuple.Vector(0, 1, 0).Equal(n2) {
		t.Fatalf("normal wrong")
	}
	if !tuple.Vector(0, 1, 0).Equal(n3) {
		t.Fatalf("normal wrong")
	}
}

func TestPlaneIntersectParallel(t *testing.T) {
	p := intersection.NewPlane()
	r := ray.New(tuple.Point(0, 10, 0), tuple.Vector(0, 0, 1))
	xs := p.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("wrong intersection count")
	}
}

func TestPlaneIntersectCoplanar(t *testing.T) {
	p := intersection.NewPlane()
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	xs := p.Intersect(r)

	if len(xs) != 0 {
		t.Fatalf("wrong intersection count")
	}
}

func TestPlaneIntersectAbove(t *testing.T) {
	p := intersection.NewPlane()
	r := ray.New(tuple.Point(0, 1, 0), tuple.Vector(0, -1, 0))
	xs := p.Intersect(r)

	if len(xs) != 1 {
		t.Fatalf("wrong intersection count")
	}
	if !flt.Equal(xs[0].T, 1) {
		t.Fatalf("wrong t")
	}
	if xs[0].Object != p {
		t.Fatalf("wrong object")
	}
}

func TestPlaneIntersectBelow(t *testing.T) {
	p := intersection.NewPlane()
	r := ray.New(tuple.Point(0, -1, 0), tuple.Vector(0, 1, 0))
	xs := p.Intersect(r)

	if len(xs) != 1 {
		t.Fatalf("wrong intersection count")
	}
	if !flt.Equal(xs[0].T, 1) {
		t.Fatalf("wrong t")
	}
	if xs[0].Object != p {
		t.Fatalf("wrong object")
	}
}
