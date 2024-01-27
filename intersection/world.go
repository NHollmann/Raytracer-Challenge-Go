package intersection

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/light"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type World struct {
	Objects []Shape
	Lights  []light.PointLight
}

func NewWorld() World {
	return World{
		Objects: []Shape{},
		Lights:  []light.PointLight{},
	}
}

func NewDefaultWorld() World {
	s1 := NewSphere()
	s1.Material.Color = color.New(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = matrix.Scaling(0.5, 0.5, 0.5)

	return World{
		Objects: []Shape{s1, s2},
		Lights: []light.PointLight{
			light.NewPoint(tuple.Point(-10, 10, -10), color.New(1, 1, 1)),
		},
	}
}

func (w *World) AddObject(obj Shape) {
	w.Objects = append(w.Objects, obj)
}

func (w *World) AddLight(l light.PointLight) {
	w.Lights = append(w.Lights, l)
}

func (w *World) Intersect(r ray.Ray) Intersections {
	result := make(Intersections, 0, 10)

	for idx := range w.Objects {
		xs := w.Objects[idx].Intersect(r)
		result = append(result, xs...)
	}

	result.Sort()
	return result
}

func (w *World) ShadeHit(comps PreparedComps, remaining int) color.Color {
	result := color.New(0, 0, 0)

	for _, light := range w.Lights {
		shadowed := w.IsShadowed(comps.OverPoint, light)
		surface := comps.Object.GetMaterial().Lighting(
			light,
			comps.Object.GetTransform().Inverse(),
			comps.OverPoint,
			comps.EyeV,
			comps.NormalV,
			shadowed,
		)

		result = result.Add(surface)
	}

	reflected := w.ReflectedColor(comps, remaining)
	refracted := w.RefractedColor(comps, remaining)

	return result.Add(reflected).Add(refracted)
}

func (w *World) ColorAt(r ray.Ray, remaining int) color.Color {
	xs := w.Intersect(r)
	hit := xs.Hit()
	if hit == nil {
		return color.New(0, 0, 0)
	}

	comps := hit.PrepareComputations(r, xs)
	return w.ShadeHit(comps, remaining)
}

func (w *World) ReflectedColor(comps PreparedComps, remaining int) color.Color {
	if remaining < 1 {
		return color.New(0.0, 0.0, 0.0)
	}
	if flt.Equal(comps.Object.GetMaterial().Reflective, 0.0) {
		return color.New(0.0, 0.0, 0.0)
	}
	reflectedRay := ray.New(comps.OverPoint, comps.ReflectV)
	c := w.ColorAt(reflectedRay, remaining-1)
	return c.MulScalar(comps.Object.GetMaterial().Reflective)
}

func (w *World) RefractedColor(comps PreparedComps, remaining int) color.Color {
	if remaining < 1 {
		return color.New(0.0, 0.0, 0.0)
	}
	if flt.Equal(comps.Object.GetMaterial().Transparency, 0.0) {
		return color.New(0.0, 0.0, 0.0)
	}

	nRatio := comps.N1 / comps.N2
	cosI := comps.EyeV.Dot(comps.NormalV)
	sin2t := nRatio * nRatio * (1 - cosI*cosI)
	if sin2t > 1.0 {
		return color.New(0.0, 0.0, 0.0)
	}

	cosT := math.Sqrt(1.0 - sin2t)
	direction := comps.NormalV.Mul(nRatio*cosI - cosT).Sub(comps.EyeV.Mul(nRatio))
	refractedRay := ray.New(comps.UnderPoint, direction)
	result := w.ColorAt(refractedRay, remaining-1)
	result = result.MulScalar(comps.Object.GetMaterial().Transparency)

	return result
}

func (w *World) IsShadowed(point tuple.Tuple, l light.PointLight) bool {
	v := l.Position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := ray.New(point, direction)
	intersections := w.Intersect(r)

	h := intersections.Hit()
	return h != nil && h.T < distance
}
