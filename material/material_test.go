package material_test

import (
	"math"
	"testing"

	"nicolashollmann.de/raytracer-challange/color"
	"nicolashollmann.de/raytracer-challange/flt"
	"nicolashollmann.de/raytracer-challange/light"
	"nicolashollmann.de/raytracer-challange/material"
	"nicolashollmann.de/raytracer-challange/tuple"
)

// L = Light
// E = Eye
// S = Surface
// O = Offset

func TestMaterialConstructor(t *testing.T) {
	m := material.New()

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
