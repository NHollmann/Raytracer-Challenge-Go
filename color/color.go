package color

import "nicolashollmann.de/raytracer-challange/flt"

type Color [3]float64

func New(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c *Color) R() float64 {
	return c[0]
}

func (c *Color) G() float64 {
	return c[1]
}

func (c *Color) B() float64 {
	return c[2]
}

func (a Color) Equal(b Color) bool {
	return flt.Equal(a[0], b[0]) &&
		flt.Equal(a[1], b[1]) &&
		flt.Equal(a[2], b[2])
}

func (a Color) Add(b Color) Color {
	return Color{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func (a Color) Sub(b Color) Color {
	return Color{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func (a Color) MulScalar(s float64) Color {
	return Color{
		a[0] * s,
		a[1] * s,
		a[2] * s,
	}
}

func (a Color) MulColor(b Color) Color {
	return Color{
		a[0] * b[0],
		a[1] * b[1],
		a[2] * b[2],
	}
}
