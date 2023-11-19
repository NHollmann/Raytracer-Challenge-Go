package material_test

import (
	"math"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/material"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

// L = Light
// E = Eye
// S = Surface
// O = Offset

func TestMaterialConstructor(t *testing.T) {
	m := material.New()

	if m.Pattern != nil {
		t.Errorf("material default pattern is not nil")
	}
	if !m.Color.Equal(color.New(1, 1, 1)) {
		t.Errorf("material default color wrong")
	}
	if !flt.Equal(m.Ambient, 0.1) {
		t.Errorf("material default ambient wrong")
	}
	if !flt.Equal(m.Diffuse, 0.9) {
		t.Errorf("material default diffuse wrong")
	}
	if !flt.Equal(m.Specular, 0.9) {
		t.Errorf("material default specular wrong")
	}
	if !flt.Equal(m.Shininess, 200.0) {
		t.Errorf("material default shininess wrong")
	}
}

// Eye between light and surface
func TestMaterialShadingLES(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, false)

	if !result.Equal(color.New(1.9, 1.9, 1.9)) {
		t.Errorf("lighting result color wrong")
	}
}

// Eye between light and surface in Shadow
func TestMaterialShadingLESShadow(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, true)

	if !result.Equal(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("lighting result color wrong")
	}
}

// Eye between light and surface, eye offset 45°
func TestMaterialShadingLESO45(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, false)

	if !result.Equal(color.New(1.0, 1.0, 1.0)) {
		t.Errorf("lighting result color wrong")
	}
}

// Eye opposite surface, light offset 45°
func TestMaterialShadingESLO45(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 10, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, false)

	if !result.Equal(color.New(0.7364, 0.7364, 0.7364)) {
		t.Errorf("lighting result color wrong")
	}
}

// Eye in path of the reflection vector
func TestMaterialShadingEO45LO45(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, -math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 10, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, false)

	if !result.Equal(color.New(1.6364, 1.6364, 1.6364)) {
		t.Errorf("lighting result color wrong")
	}
}

// Light behind surface
func TestMaterialShadingESL(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, 10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, false)

	if !result.Equal(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("lighting result color wrong")
	}
}

// Light behind surface in shadow
func TestMaterialShadingESLShadow(t *testing.T) {
	m := material.New()
	pos := tuple.Point(0, 0, 0)
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))
	result := m.Lighting(l, pos, eyev, normalv, true)

	if !result.Equal(color.New(0.1, 0.1, 0.1)) {
		t.Errorf("lighting result color wrong")
	}
}

func TestMaterialPattern(t *testing.T) {
	m := material.New()
	m.Pattern = pattern.NewStripePattern(color.New(1, 1, 1), color.New(0, 0, 0))
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	l := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))

	result1 := m.Lighting(l, tuple.Point(0.9, 0, 0), eyev, normalv, false)
	if !result1.Equal(color.New(1, 1, 1)) {
		t.Errorf("lighting result color wrong")
	}

	result2 := m.Lighting(l, tuple.Point(1.1, 0, 0), eyev, normalv, false)
	if !result2.Equal(color.New(0, 0, 0)) {
		t.Errorf("lighting result color wrong")
	}
}
