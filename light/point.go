package light

import (
	"nicolashollmann.de/raytracer-challange/color"
	"nicolashollmann.de/raytracer-challange/tuple"
)

type PointLight struct {
	Intensity color.Color
	Position  tuple.Tuple
}

func NewPoint(Position tuple.Tuple, Intensity color.Color) PointLight {
	return PointLight{
		Intensity,
		Position,
	}
}
