package light

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
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
