package matrix

import (
	"fmt"

	"nicolashollmann.de/raytracer-challange/flt"
	"nicolashollmann.de/raytracer-challange/tuple"
)

type Mat44 [16]float64

func Identity44() Mat44 {
	return Mat44{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m Mat44) Index(r, c int) float64 {
	return m[r*4+c]
}

func (a Mat44) Equal(b Mat44) bool {
	for i := 0; i < 16; i++ {
		if !flt.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func (m Mat44) String() string {
	res := ""

	for r := 0; r < 4; r++ {
		res += "|"
		for c := 0; c < 4; c++ {
			res += fmt.Sprintf("%f|", m.Index(r, c))
		}
		res += "\n"
	}

	return res
}

func (a Mat44) Mul(b Mat44) Mat44 {
	res := Mat44{}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res[r*4+c] = a.Index(r, 0)*b.Index(0, c) +
				a.Index(r, 1)*b.Index(1, c) +
				a.Index(r, 2)*b.Index(2, c) +
				a.Index(r, 3)*b.Index(3, c)
		}
	}

	return res
}

func (a Mat44) MulTuple(b tuple.Tuple) tuple.Tuple {
	res := tuple.Tuple{}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res[r] = a.Index(r, 0)*b.X() +
				a.Index(r, 1)*b.Y() +
				a.Index(r, 2)*b.Z() +
				a.Index(r, 3)*b.W()
		}
	}

	return res
}

func (a Mat44) Transpose() Mat44 {
	res := Mat44{}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res[r*4+c] = a.Index(c, r)
		}
	}

	return res
}

func (m Mat44) Submatrix(rx, cx int) Mat33 {
	res := Mat33{}

	r4 := 0
	c4 := 0
	for r3 := 0; r3 < 3; r3++ {
		if r4 == rx {
			r4++
		}
		for c3 := 0; c3 < 3; c3++ {
			if c4 == cx {
				c4++
			}

			res[r3*3+c3] = m.Index(r4, c4)
			c4++
		}
		c4 = 0
		r4++
	}

	return res
}

func (m Mat44) Minor(rx, cx int) float64 {
	return m.Submatrix(rx, cx).Determinant()
}

func (m Mat44) Cofactor(rx, cx int) float64 {
	sign := 1.0
	if (rx+cx)%2 == 1 {
		sign = -1.0
	}
	return sign * m.Minor(rx, cx)
}

func (a Mat44) Determinant() float64 {
	det := 0.0
	for c := 0; c < 4; c++ {
		det += a.Index(0, c) * a.Cofactor(0, c)
	}
	return det
}

func (a Mat44) IsInvertible() bool {
	return !flt.Equal(a.Determinant(), 0.0)
}

func (a Mat44) Inverse() Mat44 {
	if !a.IsInvertible() {
		panic("matrix is not invertible")
	}

	res := Mat44{}
	det := a.Determinant()

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			cofactor := a.Cofactor(r, c)

			// Column and row are swapped to implement a transpose operation
			res[c*4+r] = cofactor / det
		}
	}

	return res
}
