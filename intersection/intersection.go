package intersection

import (
	"slices"
	"sort"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Intersection struct {
	T      float64
	Object Shape
}

type Intersections []Intersection

type PreparedComps struct {
	T          float64
	Object     Shape
	Point      tuple.Tuple
	OverPoint  tuple.Tuple
	UnderPoint tuple.Tuple
	EyeV       tuple.Tuple
	NormalV    tuple.Tuple
	ReflectV   tuple.Tuple
	Inside     bool
	N1         float64 // Refractive index of exiting material
	N2         float64 // Refractive index of entering material
}

func NewIntersection(T float64, Object Shape) Intersection {
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

func (x Intersection) PrepareComputations(r ray.Ray, xs Intersections) PreparedComps {
	if xs == nil {
		xs = Intersections{x}
	}

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
	result.UnderPoint = result.Point.Sub(result.NormalV.Mul(flt.Epsilon))
	result.ReflectV = r.Direction.Reflect(result.NormalV)

	// Find refractive indizies of entering and exiting material
	containers := make([]Shape, 0, len(xs))
	for _, i := range xs {
		if i == x {
			obj, found := last(containers)
			if found {
				result.N1 = obj.GetMaterial().RefractiveIndex
			} else {
				result.N1 = 1.0
			}
		}
		idx := slices.Index(containers, i.Object)
		if idx >= 0 {
			containers = slices.Delete(containers, idx, idx+1)
		} else {
			containers = append(containers, i.Object)
		}
		if i == x {
			obj, found := last(containers)
			if found {
				result.N2 = obj.GetMaterial().RefractiveIndex
			} else {
				result.N2 = 1.0
			}
			break
		}
	}

	return result
}

func last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
