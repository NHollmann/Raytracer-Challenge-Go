package intersection_test

import (
	"math"
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

	comps := i.PrepareComputations(r, nil)
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

	comps := i.PrepareComputations(r, nil)
	if comps.Inside {
		t.Errorf("should be outside")
	}
}

func TestPrecomputeIntersectionInside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	i := intersection.NewIntersection(1, shape)

	comps := i.PrepareComputations(r, nil)
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

func TestPrecomputeReflectV(t *testing.T) {
	r := ray.New(tuple.Point(0, 1, -1), tuple.Vector(0, -math.Sqrt(2)/2.0, math.Sqrt(2)/2.0))
	shape := intersection.NewPlane()
	i := intersection.NewIntersection(math.Sqrt(2), shape)

	comps := i.PrepareComputations(r, nil)
	if !tuple.Vector(0, math.Sqrt(2)/2.0, math.Sqrt(2)/2.0).Equal(comps.ReflectV) {
		t.Errorf("wrong reflectV")
	}
}

func TestHitOverPoint(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := intersection.NewSphere()
	shape.Transform = matrix.Translation(0, 0, 1)
	i := intersection.NewIntersection(5, shape)
	comps := i.PrepareComputations(r, nil)
	if comps.OverPoint.Z() >= flt.Epsilon/2.0 || comps.Point.Z() <= comps.OverPoint.Z() {
		t.Errorf("over point wrong")
	}
}

func TestHitUnderPoint(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := NewGlassSphere()
	shape.Transform = matrix.Translation(0, 0, 1)
	i := intersection.NewIntersection(5, shape)
	comps := i.PrepareComputations(r, nil)
	if comps.UnderPoint.Z() <= flt.Epsilon/2.0 || comps.Point.Z() >= comps.UnderPoint.Z() {
		t.Errorf("under point wrong")
	}
}

func TestFindN1AndN2(t *testing.T) {
	cs := []struct {
		index int
		n1    float64
		n2    float64
	}{
		{0, 1.0, 1.5},
		{1, 1.5, 2.0},
		{2, 2.0, 2.5},
		{3, 2.5, 2.5},
		{4, 2.5, 1.5},
		{5, 1.5, 1.0},
	}

	a := NewGlassSphere()
	a.Transform = matrix.Scaling(2, 2, 2)
	a.Material.RefractiveIndex = 1.5

	b := NewGlassSphere()
	b.Transform = matrix.Translation(0, 0, -0.25)
	b.Material.RefractiveIndex = 2.0

	c := NewGlassSphere()
	c.Transform = matrix.Translation(0, 0, 0.25)
	c.Material.RefractiveIndex = 2.5

	r := ray.New(tuple.Point(0, 0, -4), tuple.Vector(0, 0, 1))

	i1 := intersection.NewIntersection(2.0, a)
	i2 := intersection.NewIntersection(2.75, b)
	i3 := intersection.NewIntersection(3.25, c)
	i4 := intersection.NewIntersection(4.75, b)
	i5 := intersection.NewIntersection(5.25, c)
	i6 := intersection.NewIntersection(6.0, a)
	xs := intersection.Intersections{i1, i2, i3, i4, i5, i6}

	for _, c := range cs {

		comps := xs[c.index].PrepareComputations(r, xs)
		if !flt.Equal(comps.N1, c.n1) || !flt.Equal(comps.N2, c.n2) {
			t.Errorf("wrong n1 or n2 at index %d", c.index)
		}
	}
}
