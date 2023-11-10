package matrix_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
)

func TestMat33Index(t *testing.T) {
	m := matrix.Mat33{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	if !flt.Equal(m.Index(0, 0), 1) {
		t.Errorf("Mat33 [0,0] doesn't equal 1")
	}

	if !flt.Equal(m.Index(1, 1), 5) {
		t.Errorf("Mat33 [1,1] doesn't equal 5")
	}

	if !flt.Equal(m.Index(2, 2), 9) {
		t.Errorf("Mat33 [2,2] doesn't equal 9")
	}

	if !flt.Equal(m.Index(1, 2), 6) {
		t.Errorf("Mat33 [1,2] doesn't equal 6")
	}
}

func TestMat33Equal(t *testing.T) {
	a := matrix.Mat33{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	b := matrix.Mat33{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	c := matrix.Mat33{
		1, 2, 3,
		4, 5, 9,
		7, 8, 9,
	}

	if !a.Equal(b) {
		t.Errorf("Mat33 a and b aren't equal")
	}
	if a.Equal(c) {
		t.Errorf("Mat33 a and c are equal")
	}
}

func TestMat33Print(t *testing.T) {
	a := matrix.Mat33{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	expected := "|1.000000|2.000000|3.000000|\n"
	expected += "|4.000000|5.000000|6.000000|\n"
	expected += "|7.000000|8.000000|9.000000|\n"

	if a.String() != expected {
		t.Errorf("Mat33 string is wrong, got=\n%s", a.String())
	}
}

func TestMat33Submatrix(t *testing.T) {
	a := matrix.Mat33{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	expected := matrix.Mat22{
		4, 6,
		7, 9,
	}

	if !a.Submatrix(0, 1).Equal(expected) {
		t.Errorf("Mat33 submatrix is wrong, got=\n%v", a.Submatrix(0, 1))
	}
}

func TestMat33Minor(t *testing.T) {
	a := matrix.Mat33{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	}
	b := a.Submatrix(1, 0)

	if !flt.Equal(b.Determinant(), 25) {
		t.Errorf("Mat33 submatrix determinant is wrong, got=%v", b.Determinant())
	}
	if !flt.Equal(a.Minor(1, 0), 25) {
		t.Errorf("Mat33 minor is wrong, got=%v", a.Minor(1, 0))
	}
}

func TestMat33Cofactor(t *testing.T) {
	a := matrix.Mat33{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	}
	if !flt.Equal(a.Minor(0, 0), -12) {
		t.Errorf("Mat33 minor 0,0 is wrong, got=%v", a.Minor(0, 0))
	}
	if !flt.Equal(a.Cofactor(0, 0), -12) {
		t.Errorf("Mat33 cofactor 0,0 is wrong, got=%v", a.Cofactor(0, 0))
	}
	if !flt.Equal(a.Minor(1, 0), 25) {
		t.Errorf("Mat33 minor 1,0 is wrong, got=%v", a.Minor(1, 0))
	}
	if !flt.Equal(a.Cofactor(1, 0), -25) {
		t.Errorf("Mat33 cofactor 1,0 is wrong, got=%v", a.Cofactor(1, 0))
	}
}

func TestMat33Determinant(t *testing.T) {
	a := matrix.Mat33{
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	}
	if !flt.Equal(a.Cofactor(0, 0), 56) {
		t.Errorf("Mat33 cofactor 0,0 is wrong, got=%v", a.Cofactor(0, 0))
	}
	if !flt.Equal(a.Cofactor(0, 1), 12) {
		t.Errorf("Mat33 cofactor 0,1 is wrong, got=%v", a.Cofactor(0, 1))
	}
	if !flt.Equal(a.Cofactor(0, 2), -46) {
		t.Errorf("Mat33 cofactor 0,2 is wrong, got=%v", a.Cofactor(0, 2))
	}
	if !flt.Equal(a.Determinant(), -196) {
		t.Errorf("Mat33 determinant is wrong, got=%v", a.Determinant())
	}
}
