package matrix

import (
	"fmt"

	"nicolashollmann.de/raytracer-challange/flt"
)

type Mat33 [9]float64

func (m Mat33) Index(r, c int) float64 {
	return m[r*3+c]
}

func (a Mat33) Equal(b Mat33) bool {
	for i := 0; i < 9; i++ {
		if !flt.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func (m Mat33) String() string {
	res := ""

	for r := 0; r < 3; r++ {
		res += "|"
		for c := 0; c < 3; c++ {
			res += fmt.Sprintf("%f|", m.Index(r, c))
		}
		res += "\n"
	}

	return res
}

func (m Mat33) Submatrix(rx, cx int) Mat22 {
	res := Mat22{}

	r3 := 0
	c3 := 0
	for r2 := 0; r2 < 2; r2++ {
		if r3 == rx {
			r3++
		}
		for c2 := 0; c2 < 2; c2++ {
			if c3 == cx {
				c3++
			}

			res[r2*2+c2] = m.Index(r3, c3)
			c3++
		}
		c3 = 0
		r3++
	}

	return res
}

func (m Mat33) Minor(rx, cx int) float64 {
	return m.Submatrix(rx, cx).Determinant()
}

func (m Mat33) Cofactor(rx, cx int) float64 {
	sign := 1.0
	if (rx+cx)%2 == 1 {
		sign = -1.0
	}
	return sign * m.Minor(rx, cx)
}

func (a Mat33) Determinant() float64 {
	det := 0.0
	for c := 0; c < 3; c++ {
		det += a.Index(0, c) * a.Cofactor(0, c)
	}
	return det
}
