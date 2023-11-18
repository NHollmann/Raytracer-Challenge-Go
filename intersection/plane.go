package intersection

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Plane struct {
	*BaseShape
}

func NewPlane() *Plane {
	shape := NewShape()
	plane := &Plane{
		shape,
	}
	shape.Shape = plane
	return plane
}

func (s *Plane) localIntersect(r ray.Ray) Intersections {
	if math.Abs(r.Direction.Y()) < flt.Epsilon {
		return Intersections{}
	}

	t := -r.Origin.Y() / r.Direction.Y()
	return Intersections{
		NewIntersection(t, s),
	}
}

func (s *Plane) localNormalAt(p tuple.Tuple) tuple.Tuple {
	return tuple.Vector(0, 1, 0)
}
