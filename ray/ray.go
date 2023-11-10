package ray

import (
	"nicolashollmann.de/raytracer-challange/matrix"
	"nicolashollmann.de/raytracer-challange/tuple"
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
