package intersection_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

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
	if !w.Objects[0].Material.Color.Equal(color.New(0.8, 1.0, 0.6)) {
		t.Errorf("object 0 wrong color")
	}
	if !flt.Equal(w.Objects[0].Material.Diffuse, 0.7) {
		t.Errorf("object 0 wrong diffuse")
	}
	if !flt.Equal(w.Objects[0].Material.Specular, 0.2) {
		t.Errorf("object 0 wrong specular")
	}
	if !w.Objects[1].Transform.Equal(matrix.Scaling(0.5, 0.5, 0.5)) {
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
	i := intersection.NewIntersection(4, &shape)
	comps := i.PrepareComputations(r)

	c := w.ShadeHit(comps)
	if !c.Equal(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wrong color")
	}
}

func TestWorldShadeInside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	w.Lights[0] = light.NewPoint(tuple.Point(0, 0.25, 0), color.New(1, 1, 1))
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	shape := w.Objects[1]
	i := intersection.NewIntersection(0.5, &shape)
	comps := i.PrepareComputations(r)

	c := w.ShadeHit(comps)
	if !c.Equal(color.New(0.90498, 0.90498, 0.90498)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtMiss(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 1, 0))

	c := w.ColorAt(r)

	if !c.Equal(color.New(0, 0, 0)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtOutside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

	c := w.ColorAt(r)

	if !c.Equal(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("wrong color")
	}
}

func TestWorldColorAtInside(t *testing.T) {
	w := intersection.NewDefaultWorld()
	w.Objects[0].Material.Ambient = 1
	w.Objects[1].Material.Ambient = 1
	inner := w.Objects[1]
	r := ray.New(tuple.Point(0, 0, 0.75), tuple.Vector(0, 0, -1))

	c := w.ColorAt(r)

	if !c.Equal(inner.Material.Color) {
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
	i := intersection.NewIntersection(4, &s2)

	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	if !c.Equal(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("wrong color")
	}
}
