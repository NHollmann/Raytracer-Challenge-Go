package material

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/pattern"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Material struct {
	Pattern    pattern.Pattern
	Color      color.Color
	Ambient    float64
	Diffuse    float64
	Specular   float64
	Shininess  float64
	Reflective float64
}

func New() Material {
	return Material{
		Pattern:    nil,
		Color:      color.New(1, 1, 1),
		Ambient:    0.1,
		Diffuse:    0.9,
		Specular:   0.9,
		Shininess:  200.0,
		Reflective: 0.0,
	}
}

func (m *Material) Lighting(l light.PointLight, invTransform matrix.Mat44, point, eyev, normalv tuple.Tuple, inShadow bool) color.Color {
	matColor := m.Color
	if m.Pattern != nil {
		matColor = m.Pattern.PatternAtTransform(invTransform, point)
	}

	effectiveColor := matColor.MulColor(l.Intensity)
	lightv := l.Position.Sub(point).Normalize()
	ambient := effectiveColor.MulScalar(m.Ambient)
	diffuse := color.New(0, 0, 0)
	specular := color.New(0, 0, 0)

	if inShadow {
		return ambient
	}

	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal >= 0.0 {
		diffuse = effectiveColor.MulScalar(m.Diffuse).MulScalar(lightDotNormal)

		reflectv := lightv.Neg().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)
		if reflectDotEye > 0.0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.MulScalar(m.Specular).MulScalar(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
