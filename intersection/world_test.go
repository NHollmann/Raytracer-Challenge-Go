package intersection_test

import (
	"math"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

const DEFAULT_RECURSION_MAX = 5

func TestWorldConstructor(t *testing.T) {
	w := intersection.NewWorld()

	if len(w.Objects) != 0 {
		t.Errorf("new world has objects")
	}
	if len(w.Lights) != 0 {
		t.Errorf("new world has lights")
	}
}

func TestDefaultWorldConstructor(t *testing.T) {
	w := intersection.NewDefaultWorld()

	if len(w.Objects) != 2 {
		t.Fatalf("default world does not have two objects but %d", len(w.Lights))
	}
	if len(w.Lights) != 1 {
		t.Fatalf("default world does not have one light but %d", len(w.Lights))
	}
	if !w.Lights[0].Intensity.Equal(color.New(1, 1, 1)) {
		t.Errorf("wrong light intensity")
	}
	if !w.Lights[0].Position.Equal(tuple.Point(-10, 10, -10)) {
		t.Errorf("wrong light position")
	}
	if !w.Objects[0].GetMaterial().Color.Equal(color.New(0.8, 1.0, 0.6)) {
		t.Errorf("object 0 wrong color")
	}
	if !flt.Equal(w.Objects[0].GetMaterial().Diffuse, 0.7) {
		t.Errorf("object 0 wrong diffuse")
	}
	if !flt.Equal(w.Objects[0].GetMaterial().Specular, 0.2) {
		t.Errorf("object 0 wrong specular")
	}
	if !w.Objects[1].GetTransform().Equal(matrix.Scaling(0.5, 0.5, 0.5)) {
		t.Errorf("object 1 wrong transform")
	}
}

func TestWorldAddObject(t *testing.T) {
	w := intersection.NewWorld()

	if len(w.Objects) != 0 {
		t.Fatalf("wrong object count %d (expected 0)", len(w.Objects))
	}

	w.AddObject(intersection.NewSphere())
	if len(w.Objects) != 1 {
		t.Fatalf("wrong object count %d (expected 1)", len(w.Objects))
	}

	w.AddObject(intersection.NewSphere())
	w.AddObject(intersection.NewSphere())
	if len(w.Objects) != 3 {
		t.Fatalf("wrong object count %d (expected 3)", len(w.Objects))
	}
}

func TestWorldAddLight(t *testing.T) {
	w := intersection.NewWorld()

	if len(w.Lights) != 0 {
		t.Fatalf("wrong light count %d (expected 0)", len(w.Lights))
	}

	w.AddLight(light.NewPoint(tuple.Point(0, 0, 0), color.New(1, 1, 1)))
	if len(w.Lights) != 1 {
		t.Fatalf("wrong light count %d (expected 1)", len(w.Lights))
	}

	w.AddLight(light.NewPoint(tuple.Point(0, 0, 0), color.New(1, 1, 1)))
	w.AddLight(light.NewPoint(tuple.Point(0, 0, 0), color.New(1, 1, 1)))
	if len(w.Lights) != 3 {
		t.Fatalf("wrong light count %d (expected 3)", len(w.Lights))
	}
}

func TestWorldIntersection(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

	xs := w.Intersect(r)

	if len(xs) != 4 {
		t.Fatalf("not 4 intersections")
	}
	if !flt.Equal(xs[0].T, 4) {
		t.Errorf("intersection 0 wrong")
	}
	if !flt.Equal(xs[1].T, 4.5) {
		t.Errorf("intersection 1 wrong")
	}
	if !flt.Equal(xs[2].T, 5.5) {
		t.Errorf("intersection 2 wrong")
	}
	if !flt.Equal(xs[3].T, 6) {
		t.Errorf("intersection 3 wrong")
	}
}

func TestWorldShade(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	shape := w.Objects[0]
	i := intersection.NewIntersection(4, shape)
	comps := i.PrepareComputations(r, nil)

	c := w.ShadeHit(comps, DEFAULT_RECURSION_MAX)
	if !c.Equal(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wrong color")
	}
}

func TestWorldShadeInside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	w.Lights[0] = light.NewPoint(tuple.Point(0, 0.25, 0), color.New(1, 1, 1))
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	shape := w.Objects[1]
	i := intersection.NewIntersection(0.5, shape)
	comps := i.PrepareComputations(r, nil)

	c := w.ShadeHit(comps, DEFAULT_RECURSION_MAX)
	if !c.Equal(color.New(0.90498, 0.90498, 0.90498)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtMiss(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 1, 0))

	c := w.ColorAt(r, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0, 0, 0)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtOutside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

	c := w.ColorAt(r, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtInside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	w.Objects[0].GetMaterial().Ambient = 1
	w.Objects[1].GetMaterial().Ambient = 1
	inner := w.Objects[1]
	r := ray.New(tuple.Point(0, 0, 0.75), tuple.Vector(0, 0, -1))

	c := w.ColorAt(r, DEFAULT_RECURSION_MAX)

	if !c.Equal(inner.GetMaterial().Color) {
		t.Errorf("wrong color")
	}
}

func TestShadowCollinear(t *testing.T) {
	w := intersection.NewDefaultWorld()
	p := tuple.Point(0, 10, 0)

	if w.IsShadowed(p, w.Lights[0]) {
		t.Errorf("wrong shadow result")
	}
}

func TestShadowBetween(t *testing.T) {
	w := intersection.NewDefaultWorld()
	p := tuple.Point(10, -10, 10)

	if !w.IsShadowed(p, w.Lights[0]) {
		t.Errorf("wrong shadow result")
	}
}

func TestShadowBehindLight(t *testing.T) {
	w := intersection.NewDefaultWorld()
	p := tuple.Point(-20, 20, -20)

	if w.IsShadowed(p, w.Lights[0]) {
		t.Errorf("wrong shadow result")
	}
}

func TestShadowBehindObject(t *testing.T) {
	w := intersection.NewDefaultWorld()
	p := tuple.Point(-2, 2, -2)

	if w.IsShadowed(p, w.Lights[0]) {
		t.Errorf("wrong shadow result")
	}
}

func TestShadeHitShadow(t *testing.T) {
	w := intersection.NewWorld()
	w.AddLight(light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1)))

	s1 := intersection.NewSphere()
	w.AddObject(s1)

	s2 := intersection.NewSphere()
	s2.Transform = matrix.Translation(0, 0, 10)
	w.AddObject(s2)

	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	i := intersection.NewIntersection(4, s2)

	comps := i.PrepareComputations(r, nil)
	c := w.ShadeHit(comps, DEFAULT_RECURSION_MAX)
	if !c.Equal(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("wrong color")
	}
}

func TestWorldReflectedColorOfNonreflective(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	w.Objects[1].GetMaterial().Ambient = 1
	i := intersection.NewIntersection(1, w.Objects[1])

	comps := i.PrepareComputations(r, nil)
	c := w.ReflectedColor(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.0, 0.0, 0.0)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldReflectedColor(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := intersection.NewPlane()
	shape.GetMaterial().Reflective = 0.5
	shape.Transform = matrix.Translation(0, -1, 0)
	w.AddObject(shape)
	r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -math.Sqrt(2)/2.0, math.Sqrt(2)/2.0))
	i := intersection.NewIntersection(math.Sqrt(2), shape)

	comps := i.PrepareComputations(r, nil)
	c := w.ReflectedColor(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.190332, 0.23791, 0.142749)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldReflectedShadeHit(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := intersection.NewPlane()
	shape.GetMaterial().Reflective = 0.5
	shape.Transform = matrix.Translation(0, -1, 0)
	w.AddObject(shape)
	r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -math.Sqrt(2)/2.0, math.Sqrt(2)/2.0))
	i := intersection.NewIntersection(math.Sqrt(2), shape)

	comps := i.PrepareComputations(r, nil)
	c := w.ShadeHit(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.876757, 0.92434, 0.829174)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldReflectedEndlessRecursion(t *testing.T) {
	w := intersection.NewWorld()
	pLight := light.NewPoint(tuple.Point(0, 0, 0), color.New(1, 1, 1))
	w.AddLight(pLight)

	lower := intersection.NewPlane()
	lower.GetMaterial().Reflective = 1.0
	lower.Transform = matrix.Translation(0, -1, 0)
	w.AddObject(lower)

	upper := intersection.NewPlane()
	upper.GetMaterial().Reflective = 1.0
	upper.Transform = matrix.Translation(0, 1, 0)
	w.AddObject(upper)

	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 1, 0))
	c := w.ColorAt(r, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(11.4, 11.4, 11.4)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldReflectedMaxDepth(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := intersection.NewPlane()
	shape.GetMaterial().Reflective = 0.5
	shape.Transform = matrix.Translation(0, -1, 0)
	w.AddObject(shape)
	r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -math.Sqrt(2)/2.0, math.Sqrt(2)/2.0))
	i := intersection.NewIntersection(math.Sqrt(2), shape)

	comps := i.PrepareComputations(r, nil)
	c := w.ReflectedColor(comps, 0)

	if !c.Equal(color.New(0.0, 0.0, 0.0)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldRefractedColorOfNonrefracted(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := w.Objects[0]
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	i1 := intersection.NewIntersection(4, shape)
	i2 := intersection.NewIntersection(6, shape)
	xs := intersection.Intersections{i1, i2}

	comps := xs[0].PrepareComputations(r, xs)
	c := w.RefractedColor(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.0, 0.0, 0.0)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldRefractedMaxDepth(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := w.Objects[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5

	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	i1 := intersection.NewIntersection(4, shape)
	i2 := intersection.NewIntersection(6, shape)
	xs := intersection.Intersections{i1, i2}

	comps := xs[0].PrepareComputations(r, xs)
	c := w.RefractedColor(comps, 0)

	if !c.Equal(color.New(0.0, 0.0, 0.0)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldRefractedTotalInternalReflection(t *testing.T) {
	w := intersection.NewDefaultWorld()
	shape := w.Objects[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5

	r := ray.New(tuple.Point(0, 0, math.Sqrt(2.0)/2.0), tuple.Vector(0, 1, 0))
	i1 := intersection.NewIntersection(-math.Sqrt(2.0)/2.0, shape)
	i2 := intersection.NewIntersection(math.Sqrt(2.0)/2.0, shape)
	xs := intersection.Intersections{i1, i2}

	comps := xs[1].PrepareComputations(r, xs)
	c := w.RefractedColor(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.0, 0.0, 0.0)) {
		t.Errorf("wrong color %+v", c)
	}
}

func TestWorldRefractedColor(t *testing.T) {
	w := intersection.NewDefaultWorld()
	a := w.Objects[0]
	a.GetMaterial().Ambient = 1.0

	b := w.Objects[1]
	b.GetMaterial().Transparency = 1.0
	b.GetMaterial().RefractiveIndex = 1.5

	r := ray.New(tuple.Point(0, 0, 0.1), tuple.Vector(0, 1, 0))
	i1 := intersection.NewIntersection(-0.9899, a)
	i2 := intersection.NewIntersection(-0.4899, b)
	i3 := intersection.NewIntersection(0.4899, b)
	i4 := intersection.NewIntersection(0.9899, a)
	xs := intersection.Intersections{i1, i2, i3, i4}

	comps := xs[2].PrepareComputations(r, xs)
	c := w.RefractedColor(comps, DEFAULT_RECURSION_MAX)

	if !c.Equal(color.New(0.0, 0.99888, 0.04725)) {
		t.Errorf("wrong color %+v", c)
	}
}
