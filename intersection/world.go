package intersection

import (
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
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

func (w *World) ShadeHit(comps PreparedComps) color.Color {
	result := color.New(0, 0, 0)

	for _, light := range w.Lights {
		shadowed := w.IsShadowed(comps.OverPoint, light)
		result = result.Add(comps.Object.GetMaterial().Lighting(
			light,
			comps.Object.GetTransform().Inverse(),
			comps.OverPoint,
			comps.EyeV,
			comps.NormalV,
			shadowed,
		))
	}

	return result
}

func (w *World) ColorAt(r ray.Ray) color.Color {
	xs := w.Intersect(r)
	hit := xs.Hit()
	if hit == nil {
		return color.New(0, 0, 0)
	}

	comps := hit.PrepareComputations(r)
	return w.ShadeHit(comps)
}

func (w *World) ReflectedColor(comps PreparedComps) color.Color {
	// TODO
	return color.New(0.0, 0.0, 0.0)
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
