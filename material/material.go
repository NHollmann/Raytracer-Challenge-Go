package material

import (
	"math"

	"nicolashollmann.de/raytracer-challange/color"
	"nicolashollmann.de/raytracer-challange/light"
	"nicolashollmann.de/raytracer-challange/tuple"
)

type Material struct {
	Color     color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func New() Material {
	return Material{
		Color:     color.New(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m *Material) Lighting(l light.PointLight, point, eyev, normalv tuple.Tuple, inShadow bool) color.Color {
	effectiveColor := m.Color.MulColor(l.Intensity)
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
