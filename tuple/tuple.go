package tuple

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
)

type Tuple [4]float64

func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) IsPoint() bool {
	return flt.Equal(t[3], 1.0)
}

func (t Tuple) IsVector() bool {
	return flt.Equal(t[3], 0.0)
}

func (t *Tuple) X() float64 {
	return t[0]
}

func (t *Tuple) Y() float64 {
	return t[1]
}

func (t *Tuple) Z() float64 {
	return t[2]
}

func (t *Tuple) W() float64 {
	return t[3]
}

func (a Tuple) Equal(b Tuple) bool {
	return flt.Equal(a[0], b[0]) &&
		flt.Equal(a[1], b[1]) &&
		flt.Equal(a[2], b[2]) &&
		flt.Equal(a[3], b[3])
}

func (a Tuple) Add(b Tuple) Tuple {
	return Tuple{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
		a[3] + b[3],
	}
}

func (a Tuple) Sub(b Tuple) Tuple {
	return Tuple{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
		a[3] - b[3],
	}
}

func (a Tuple) Neg() Tuple {
	return Tuple{
		-a[0],
		-a[1],
		-a[2],
		-a[3],
	}
}

func (a Tuple) Mul(s float64) Tuple {
	return Tuple{
		a[0] * s,
		a[1] * s,
		a[2] * s,
		a[3] * s,
	}
}

func (a Tuple) Div(s float64) Tuple {
	return Tuple{
		a[0] / s,
		a[1] / s,
		a[2] / s,
		a[3] / s,
	}
}

func (a Tuple) Magnitude() float64 {
	res := a[0] * a[0]
	res += a[1] * a[1]
	res += a[2] * a[2]
	res += a[3] * a[3]
	return math.Sqrt(res)
}

func (a Tuple) Normalize() Tuple {
	mag := a.Magnitude()
	return Tuple{
		a[0] / mag,
		a[1] / mag,
		a[2] / mag,
		a[3] / mag,
	}
}

func (a Tuple) Dot(b Tuple) float64 {
	res := a[0] * b[0]
	res += a[1] * b[1]
	res += a[2] * b[2]
	res += a[3] * b[3]
	return res
}

func (a Tuple) Cross(b Tuple) Tuple {
	return Tuple{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
		0.0,
	}
}

func (a Tuple) Reflect(normal Tuple) Tuple {
	scale := 2 * a.Dot(normal)
	return a.Sub(normal.Mul(scale))
}
