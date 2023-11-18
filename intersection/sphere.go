package intersection

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Sphere struct {
	*BaseShape
}

func NewSphere() *Sphere {
	shape := NewShape()
	sphere := &Sphere{
		shape,
	}
	shape.Shape = sphere
	return sphere
}

func (s *Sphere) localIntersect(r ray.Ray) Intersections {
	sphereToRay := r.Origin.Sub(tuple.Point(0, 0, 0))

	a := r.Direction.Dot(r.Direction)
	b := 2.0 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := (b * b) - 4.0*a*c
	if discriminant < 0.0 {
		return Intersections{}
	}

	sqrDiscriminant := math.Sqrt(discriminant)
	a2 := 2 * a

	t1 := (-b - sqrDiscriminant) / a2
	t2 := (-b + sqrDiscriminant) / a2

	return Intersections{
		NewIntersection(t1, s),
		NewIntersection(t2, s),
	}
}

func (s *Sphere) localNormalAt(p tuple.Tuple) tuple.Tuple {
	return p.Sub(tuple.Point(0, 0, 0))
}
