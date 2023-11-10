package matrix

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func Translation(x, y, z float64) Mat44 {
	return Mat44{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

func Scaling(x, y, z float64) Mat44 {
	return Mat44{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

func RotationX(phi float64) Mat44 {
	cos := math.Cos(phi)
	sin := math.Sin(phi)

	return Mat44{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1,
	}
}

func RotationY(phi float64) Mat44 {
	cos := math.Cos(phi)
	sin := math.Sin(phi)

	return Mat44{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1,
	}
}

func RotationZ(phi float64) Mat44 {
	cos := math.Cos(phi)
	sin := math.Sin(phi)

	return Mat44{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Shear(xy, xz, yx, yz, zx, zy float64) Mat44 {
	return Mat44{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1,
	}
}

func ViewTransform(from, to, up tuple.Tuple) Mat44 {
	forward := to.Sub(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := Mat44{
		left.X(), left.Y(), left.Z(), 0,
		trueUp.X(), trueUp.Y(), trueUp.Z(), 0,
		-forward.X(), -forward.Y(), -forward.Z(), 0,
		0, 0, 0, 1,
	}

	return orientation.Mul(Translation(-from.X(), -from.Y(), -from.Z()))
}
