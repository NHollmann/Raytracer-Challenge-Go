package intersection

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/material"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Sphere struct {
	Transform matrix.Mat44
	Material  material.Material
}

func NewSphere() Sphere {
	return Sphere{
		Transform: matrix.Identity44(),
		Material:  material.New(),
	}
}

func (s *Sphere) Intersect(r ray.Ray) Intersections {
	r = r.Transform(s.Transform.Inverse())

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

func (s *Sphere) NormalAt(p tuple.Tuple) tuple.Tuple {
	objectPoint := s.Transform.Inverse().MulTuple(p)
	objectNormal := objectPoint.Sub(tuple.Point(0, 0, 0))
	worldNormal := s.Transform.Inverse().Transpose().MulTuple(objectNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}
