package intersection

import (
	"sort"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Intersection struct {
	T      float64
	Object *Sphere // TODO Allgemeines Objekt (interface)
}

type Intersections []Intersection

type PreparedComps struct {
	T         float64
	Object    *Sphere // TODO Allgemeines Objekt (interface)
	Point     tuple.Tuple
	OverPoint tuple.Tuple
	EyeV      tuple.Tuple
	NormalV   tuple.Tuple
	Inside    bool
}

func NewIntersection(T float64, Object *Sphere) Intersection {
	return Intersection{
		T:      T,
		Object: Object,
	}
}

func (x1 Intersection) Equal(x2 Intersection) bool {
	return flt.Equal(x1.T, x2.T) && x1.Object == x2.Object
}

func (xs Intersections) Sort() {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
}

func (xs Intersections) Hit() *Intersection {
	for i := range xs {
		if xs[i].T > 0.0 {
			return &xs[i]
		}
	}
	return nil
}

func (x Intersection) PrepareComputations(r ray.Ray) PreparedComps {
	point := r.Position(x.T)
	result := PreparedComps{
		T:      x.T,
		Object: x.Object,

		Point:   point,
		EyeV:    r.Direction.Neg(),
		NormalV: x.Object.NormalAt(point),
		Inside:  false,
	}

	if result.NormalV.Dot(result.EyeV) < 0 {
		result.Inside = true
		result.NormalV = result.NormalV.Neg()
	}

	result.OverPoint = result.Point.Add(result.NormalV.Mul(flt.Epsilon))

	return result
}
