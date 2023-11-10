package ray

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

func New(origin, direction tuple.Tuple) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r Ray) Position(t float64) tuple.Tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Transform(m matrix.Mat44) Ray {
	return New(
		m.MulTuple(r.Origin),
		m.MulTuple(r.Direction),
	)
}
