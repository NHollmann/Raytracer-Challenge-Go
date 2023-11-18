package intersection

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/material"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Shape interface {
	Intersect(r ray.Ray) Intersections
	NormalAt(p tuple.Tuple) tuple.Tuple

	localIntersect(r ray.Ray) Intersections
	localNormalAt(p tuple.Tuple) tuple.Tuple

	GetMaterial() *material.Material
}

type BaseShape struct {
	Shape
	Transform matrix.Mat44
	Material  material.Material
}

func NewShape() *BaseShape {
	return &BaseShape{
		Transform: matrix.Identity44(),
		Material:  material.New(),
	}
}

func (s *BaseShape) GetMaterial() *material.Material {
	return &s.Material
}

func (s *BaseShape) Intersect(r ray.Ray) Intersections {
	r = r.Transform(s.Transform.Inverse())
	return s.localIntersect(r)
}

func (s *BaseShape) NormalAt(p tuple.Tuple) tuple.Tuple {
	objectPoint := s.Transform.Inverse().MulTuple(p)
	objectNormal := s.localNormalAt(objectPoint)
	worldNormal := s.Transform.Inverse().Transpose().MulTuple(objectNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}
