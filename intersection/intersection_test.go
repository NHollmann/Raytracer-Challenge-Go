package intersection_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestIntersectionConstructor(t *testing.T) {
	s := intersection.NewSphere()
	i := intersection.NewIntersection(3.5, s)

	if !flt.Equal(i.T, 3.5) {
		t.Errorf("intersections doesn't have a t of 3.5")
	}
	if i.Object != s {
		t.Errorf("intersections has the wrong object")
	}
}

func TestHitPositive(t *testing.T) {
	s := intersection.NewSphere()
	i1 := intersection.NewIntersection(1.0, s)
	i2 := intersection.NewIntersection(2.0, s)
	xs := intersection.Intersections{i1, i2}

	xs.Sort()
	i := xs.Hit()
	if !i1.Equal(*i) {
		t.Errorf("wrong hit result")
	}
}

func TestHitSomeNegative(t *testing.T) {
	s := intersection.NewSphere()
	i1 := intersection.NewIntersection(-1.0, s)
	i2 := intersection.NewIntersection(1.0, s)
	xs := intersection.Intersections{i1, i2}

	xs.Sort()
	i := xs.Hit()
	if !i2.Equal(*i) {
		t.Errorf("wrong hit result")
	}
}

func TestHitAllNegative(t *testing.T) {
	s := intersection.NewSphere()
	i1 := intersection.NewIntersection(-2.0, s)
	i2 := intersection.NewIntersection(-1.0, s)
	xs := intersection.Intersections{i1, i2}

	xs.Sort()
	i := xs.Hit()
	if i != nil {
		t.Errorf("wrong hit result")
	}
}

func TestHitAlwaysLowest(t *testing.T) {
	s := intersection.NewSphere()
	i1 := intersection.NewIntersection(5.0, s)
	i2 := intersection.NewIntersection(7.0, s)
	i3 := intersection.NewIntersection(-3.0, s)
	i4 := intersection.NewIntersection(2.0, s)
	xs := intersection.Intersections{i1, i2, i3, i4}

	xs.Sort()
	i := xs.Hit()
	if !i4.Equal(*i) {
		t.Errorf("wrong hit result")
	}
}

func TestHitEmpty(t *testing.T) {
	xs := intersection.Intersections{}

	xs.Sort()
	i := xs.Hit()
	if i != nil {
		t.Errorf("wrong hit result")
	}
}

func TestPrecomputeIntersection(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	i := intersection.NewIntersection(4, shape)

	comps := i.PrepareComputations(r)
	if !flt.Equal(comps.T, i.T) {
		t.Errorf("wrong t")
	}
	if comps.Object != i.Object {
		t.Errorf("wrong object")
	}
	if !comps.Point.Equal(tuple.Point(0, 0, -1)) {
		t.Errorf("wrong point")
	}
	if !comps.EyeV.Equal(tuple.Vector(0, 0, -1)) {
		t.Errorf("wrong eyeV")
	}
	if !comps.NormalV.Equal(tuple.Vector(0, 0, -1)) {
		t.Errorf("wrong normalV")
	}
}

func TestPrecomputeIntersectionOutside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	i := intersection.NewIntersection(4, shape)

	comps := i.PrepareComputations(r)
	if comps.Inside {
		t.Errorf("should be outside")
	}
}

func TestPrecomputeIntersectionInside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	i := intersection.NewIntersection(1, shape)

	comps := i.PrepareComputations(r)
	if !flt.Equal(comps.T, i.T) {
		t.Errorf("wrong t")
	}
	if comps.Object != i.Object {
		t.Errorf("wrong object")
	}
	if !comps.Point.Equal(tuple.Point(0, 0, 1)) {
		t.Errorf("wrong point")
	}
	if !comps.EyeV.Equal(tuple.Vector(0, 0, -1)) {
		t.Errorf("wrong eyeV")
	}
	if !comps.NormalV.Equal(tuple.Vector(0, 0, -1)) {
		t.Errorf("wrong normalV")
	}
	if !comps.Inside {
		t.Errorf("should be inside")
	}
}

func TestHitOffsetPoint(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	shape.Transform = matrix.Translation(0, 0, 1)
	i := intersection.NewIntersection(5, shape)
	comps := i.PrepareComputations(r)
	if comps.OverPoint.Z() >= flt.Epsilon/2.0 || comps.Point.Z() <= comps.OverPoint.Z() {
		t.Errorf("over point wrong")
	}
}
