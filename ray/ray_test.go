package ray_test

import (
	"testing"

	"nicolashollmann.de/raytracer-challange/matrix"
	"nicolashollmann.de/raytracer-challange/ray"
	"nicolashollmann.de/raytracer-challange/tuple"
)

func TestRayConstructor(t *testing.T) {
	origin := tuple.Point(1, 2, 3)
	direction := tuple.Vector(4, 5, 6)

	r := ray.New(origin, direction)

	if !r.Origin.Equal(origin) {
		t.Errorf("ray constructor defect origin")
	}
	if !r.Direction.Equal(direction) {
		t.Errorf("ray constructor defect direction")
	}
}

func TestRayPosition(t *testing.T) {
	r := ray.New(tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0))

	if !r.Position(0).Equal(tuple.Point(2, 3, 4)) {
		t.Errorf("ray position 0 defect")
	}
	if !r.Position(1).Equal(tuple.Point(3, 3, 4)) {
		t.Errorf("ray position 1 defect")
	}
	if !r.Position(-1).Equal(tuple.Point(1, 3, 4)) {
		t.Errorf("ray position -1 defect")
	}
	if !r.Position(2.5).Equal(tuple.Point(4.5, 3, 4)) {
		t.Errorf("ray position 2.5 defect")
	}
}

func TestRayTranslation(t *testing.T) {
	r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := matrix.Translation(3, 4, 5)

	r2 := r.Transform(m)

	if !r2.Origin.Equal(tuple.Point(4, 6, 8)) {
		t.Errorf("ray translate wrong origin")
	}
	if !r2.Direction.Equal(tuple.Vector(0, 1, 0)) {
		t.Errorf("ray translate wrong direction")
	}
}

func TestRayScale(t *testing.T) {
	r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := matrix.Scaling(2, 3, 4)

	r2 := r.Transform(m)

	if !r2.Origin.Equal(tuple.Point(2, 6, 12)) {
		t.Errorf("ray scale wrong origin")
	}
	if !r2.Direction.Equal(tuple.Vector(0, 3, 0)) {
		t.Errorf("ray scale wrong direction")
	}
}
