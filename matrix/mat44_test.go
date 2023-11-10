package matrix_test

import (
	"testing"

	"nicolashollmann.de/raytracer-challange/flt"
	"nicolashollmann.de/raytracer-challange/matrix"
	"nicolashollmann.de/raytracer-challange/tuple"
)

func TestMat44Index(t *testing.T) {
	m := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}

	cs := []struct {
		row      int
		column   int
		expected float64
	}{
		{0, 0, 1.0},
		{1, 1, 6.0},
		{2, 2, 11.0},
		{3, 3, 16.0},
		{3, 0, 13.0},
		{0, 3, 4.0},
		{2, 1, 10.0},
	}

	for _, c := range cs {
		if !flt.Equal(m.Index(c.row, c.column), c.expected) {
			t.Errorf("Mat44 [%d,%d] doesn't equal %f, got=%f", c.row, c.column, c.expected, m.Index(c.row, c.column))
		}
	}
}

func TestMat44Equal(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}
	b := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}
	c := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 3,
	}

	if !a.Equal(b) {
		t.Errorf("Mat44 a and b aren't equal")
	}
	if a.Equal(c) {
		t.Errorf("Mat44 a and c are equal")
	}
}

func TestMat44Print(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	expected := "|1.000000|2.000000|3.000000|4.000000|\n"
	expected += "|5.000000|6.000000|7.000000|8.000000|\n"
	expected += "|9.000000|10.000000|11.000000|12.000000|\n"
	expected += "|13.000000|14.000000|15.000000|16.000000|\n"

	if a.String() != expected {
		t.Errorf("Mat44 string is wrong, got=\n%s", a.String())
	}
}

func TestMat44Mul(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}
	b := matrix.Mat44{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	}
	expected := matrix.Mat44{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	}

	if !a.Mul(b).Equal(expected) {
		t.Errorf("Mat44 multiplication result is wrong, got=\n%v", a.Mul(b))
	}
}

func TestMat44MulTuple(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	}
	b := tuple.New(1, 2, 3, 1)
	expected := tuple.New(18, 24, 33, 1)

	if !a.MulTuple(b).Equal(expected) {
		t.Errorf("Mat44 tuple multiplication result is wrong, got=\n%v", a.MulTuple(b))
	}
}

func TestMat44MulIdentity(t *testing.T) {
	a := matrix.Mat44{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	}
	ident := matrix.Identity44()

	if !a.Mul(ident).Equal(a) {
		t.Errorf("Mat44 identity multiplication result is wrong, got=\n%v", a.Mul(ident))
	}
}

func TestMat44MulTupleIdentity(t *testing.T) {
	a := tuple.New(1, 2, 3, 4)
	ident := matrix.Identity44()

	if !ident.MulTuple(a).Equal(a) {
		t.Errorf("Mat44 tuple multiplication with identity result is wrong, got=\n%v", ident.MulTuple(a))
	}
}

func TestMat44Transpose(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	expected := matrix.Mat44{
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	}

	if !a.Transpose().Equal(expected) {
		t.Errorf("Mat44 transpose result is wrong, got=\n%v", a.Transpose())
	}
}

func TestMat44TransposeIdentity(t *testing.T) {
	ident := matrix.Identity44()

	if !ident.Transpose().Equal(ident) {
		t.Errorf("Mat44 transpose identity result is wrong, got=\n%v", ident.Transpose())
	}
}

func TestMat44Submatrix(t *testing.T) {
	a := matrix.Mat44{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	expected := matrix.Mat33{
		1, 3, 4,
		5, 7, 8,
		13, 15, 16,
	}

	if !a.Submatrix(2, 1).Equal(expected) {
		t.Errorf("Mat44 submatrix is wrong, got=\n%v", a.Submatrix(2, 1))
	}
}

func TestMat44Determinant(t *testing.T) {
	a := matrix.Mat44{
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	}
	if !flt.Equal(a.Cofactor(0, 0), 690) {
		t.Errorf("Mat44 cofactor 0,0 is wrong, got=%v", a.Cofactor(0, 0))
	}
	if !flt.Equal(a.Cofactor(0, 1), 447) {
		t.Errorf("Mat44 cofactor 0,1 is wrong, got=%v", a.Cofactor(0, 1))
	}
	if !flt.Equal(a.Cofactor(0, 2), 210) {
		t.Errorf("Mat44 cofactor 0,2 is wrong, got=%v", a.Cofactor(0, 2))
	}
	if !flt.Equal(a.Cofactor(0, 3), 51) {
		t.Errorf("Mat44 cofactor 0,3 is wrong, got=%v", a.Cofactor(0, 3))
	}
	if !flt.Equal(a.Determinant(), -4071) {
		t.Errorf("Mat44 determinant is wrong, got=%v", a.Determinant())
	}
}

func TestMat44IsInvertible(t *testing.T) {
	a := matrix.Mat44{
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6,
	}

	if !flt.Equal(a.Determinant(), -2120.0) {
		t.Errorf("Mat44 determinant is wrong, got=%v", a.Determinant())
	}

	if !a.IsInvertible() {
		t.Errorf("Mat44 should be invertible, nit IsInvertible returned false")
	}
}

func TestMat44IsInvertibleFalse(t *testing.T) {
	a := matrix.Mat44{
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	}

	if !flt.Equal(a.Determinant(), 0.0) {
		t.Errorf("Mat44 determinant is wrong, got=%v", a.Determinant())
	}

	if a.IsInvertible() {
		t.Errorf("Mat44 should not be invertible, nit IsInvertible returned true")
	}
}

func TestMat44InverseFull(t *testing.T) {
	a := matrix.Mat44{
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	}

	b := a.Inverse()

	if !flt.Equal(a.Determinant(), 532.0) {
		t.Errorf("Mat44 determinant is wrong, got=%v", a.Determinant())
	}

	if !flt.Equal(a.Cofactor(2, 3), -160.0) {
		t.Errorf("Mat44 cofactor at [2,3] is wrong, got=%v", a.Cofactor(2, 3))
	}

	if !flt.Equal(b.Index(3, 2), -160.0/532.0) {
		t.Errorf("Mat44 inverse at [3,2] is wrong, got=%v", a.Index(3, 2))
	}

	if !flt.Equal(a.Cofactor(3, 2), 105.0) {
		t.Errorf("Mat44 cofactor at [3,2] is wrong, got=%v", a.Cofactor(3, 2))
	}

	if !flt.Equal(b.Index(2, 3), 105.0/532.0) {
		t.Errorf("Mat44 inverse at [2,3] is wrong, got=%v", a.Index(2, 3))
	}

	expected := matrix.Mat44{
		0.218045, 0.451128, 0.240602, -0.045113,
		-0.808271, -1.456767, -0.443609, 0.520677,
		-0.078947, -0.223684, -0.052632, 0.197368,
		-0.522556, -0.813910, -0.300752, 0.306391,
	}

	if !b.Equal(expected) {
		t.Errorf("Mat44 inverse is wrong, got=\n%v", b)
	}
}

func TestMat44InverseOne(t *testing.T) {
	a := matrix.Mat44{
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	}

	b := a.Inverse()

	expected := matrix.Mat44{
		-0.153846, -0.153846, -0.282051, -0.538462,
		-0.076923, 0.123077, 0.025641, 0.030769,
		0.358974, 0.358974, 0.435897, 0.923077,
		-0.692308, -0.692308, -0.769231, -1.923077,
	}

	if !b.Equal(expected) {
		t.Errorf("Mat44 inverse is wrong, got=\n%v", b)
	}
}

func TestMat44InverseTwo(t *testing.T) {
	a := matrix.Mat44{
		9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2,
	}

	b := a.Inverse()

	expected := matrix.Mat44{
		-0.040741, -0.077778, 0.144444, -0.222222,
		-0.077778, 0.033333, 0.366667, -0.333333,
		-0.029012, -0.146296, -0.109259, 0.129630,
		0.177778, 0.066667, -0.266667, 0.333333,
	}

	if !b.Equal(expected) {
		t.Errorf("Mat44 inverse is wrong, got=\n%v", b)
	}
}

func TestMat44InverseAsDivision(t *testing.T) {
	a := matrix.Mat44{
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	}

	b := matrix.Mat44{
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	}

	c := a.Mul(b)

	if !c.Mul(b.Inverse()).Equal(a) {
		t.Errorf("Mat44 multiplying a product by its inverse is wrong, got=\n%v", c.Mul(b.Inverse()))
	}
}
