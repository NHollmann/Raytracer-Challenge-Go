package matrix

import (
	"fmt"

	"nicolashollmann.de/raytracer-challange/flt"
)

type Mat22 [4]float64

func (m Mat22) Index(r, c int) float64 {
	return m[r*2+c]
}

func (a Mat22) Equal(b Mat22) bool {
	for i := 0; i < 4; i++ {
		if !flt.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func (m Mat22) String() string {
	res := ""

	for r := 0; r < 2; r++ {
		res += "|"
		for c := 0; c < 2; c++ {
			res += fmt.Sprintf("%f|", m.Index(r, c))
		}
		res += "\n"
	}

	return res
}

func (a Mat22) Determinant() float64 {
	return a[0]*a[3] - a[1]*a[2]
}
