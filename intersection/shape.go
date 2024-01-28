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
	GetTransform() *matrix.Mat44
	GetInvTransform() *matrix.Mat44
}

type BaseShape struct {
	Shape
	transform    matrix.Mat44
	invTransform matrix.Mat44
	Material     material.Material
}

func NewShape() *BaseShape {
	return &BaseShape{
		transform:    matrix.Identity44(),
		invTransform: matrix.Identity44(),
		Material:     material.New(),
	}
}

func (s *BaseShape) GetMaterial() *material.Material {
	return &s.Material
}

func (s *BaseShape) SetTransform(transform matrix.Mat44) {
	s.transform = transform
	s.invTransform = transform.Inverse()
}

func (s *BaseShape) GetTransform() *matrix.Mat44 {
	return &s.transform
}

func (s *BaseShape) GetInvTransform() *matrix.Mat44 {
	return &s.invTransform
}

func (s *BaseShape) Intersect(r ray.Ray) Intersections {
	r = r.Transform(s.invTransform)
	return s.localIntersect(r)
}

func (s *BaseShape) NormalAt(p tuple.Tuple) tuple.Tuple {
	objectPoint := s.invTransform.MulTuple(p)
	objectNormal := s.localNormalAt(objectPoint)
	worldNormal := s.invTransform.Transpose().MulTuple(objectNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}
