package matrix_test

import (
	"testing"

	"nicolashollmann.de/raytracer-challange/flt"
	"nicolashollmann.de/raytracer-challange/matrix"
)

func TestMat22Index(t *testing.T) {
	m := matrix.Mat22{
		1, 2,
		3, 4,
	}

	if !flt.Equal(m.Index(0, 0), 1) {
		t.Errorf("Mat22 [0,0] doesn't equal 1")
	}

	if !flt.Equal(m.Index(0, 1), 2) {
		t.Errorf("Mat22 [0,1] doesn't equal 2")
	}

	if !flt.Equal(m.Index(1, 0), 3) {
		t.Errorf("Mat22 [1,0] doesn't equal 3")
	}

	if !flt.Equal(m.Index(1, 1), 4) {
		t.Errorf("Mat22 [1,1] doesn't equal 4")
	}
}

func TestMat22Equal(t *testing.T) {
	a := matrix.Mat22{
		1, 2,
		3, 4,
	}
	b := matrix.Mat22{
		1, 2,
		3, 4,
	}
	c := matrix.Mat22{
		1, 2,
		4, 4,
	}

	if !a.Equal(b) {
		t.Errorf("Mat22 a and b aren't equal")
	}
	if a.Equal(c) {
		t.Errorf("Mat22 a and c are equal")
	}
}

func TestMat22Print(t *testing.T) {
	a := matrix.Mat22{
		1, 2,
		3, 4,
	}
	expected := "|1.000000|2.000000|\n|3.000000|4.000000|\n"

	if a.String() != expected {
		t.Errorf("Mat22 string is wrong, got=\n%s", a.String())
	}
}

func TestMat22Determinant(t *testing.T) {
	m := matrix.Mat22{
		1, 5,
		-3, 2,
	}

	if !flt.Equal(m.Determinant(), 17) {
		t.Errorf("Mat22 determinant is not equal 17, got=%f", m.Determinant())
	}
}
