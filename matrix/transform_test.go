package matrix_test

import (
	"math"
	"testing"

	"nicolashollmann.de/raytracer-challange/matrix"
	"nicolashollmann.de/raytracer-challange/tuple"
)

func TestTranslation(t *testing.T) {
	transform := matrix.Translation(5, -3, 2)
	p := tuple.Point(-3, 4, 5)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(2, 1, 7)) {
		t.Errorf("Translation result is wrong, got=\n%v", res)
	}
}

func TestTranslationReverse(t *testing.T) {
	transform := matrix.Translation(5, -3, 2)
	inv := transform.Inverse()
	p := tuple.Point(-3, 4, 5)
	res := inv.MulTuple(p)

	if !res.Equal(tuple.Point(-8, 7, 3)) {
		t.Errorf("Translation result is wrong, got=\n%v", res)
	}
}

func TestTranslationVector(t *testing.T) {
	transform := matrix.Translation(5, -3, 2)
	v := tuple.Vector(-3, 4, 5)
	res := transform.MulTuple(v)

	if !res.Equal(v) {
		t.Errorf("Translation result is wrong, got=\n%v", res)
	}
}

func TestScaling(t *testing.T) {
	transform := matrix.Scaling(2, 3, 4)
	p := tuple.Point(-4, 6, 8)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(-8, 18, 32)) {
		t.Errorf("Scaling result is wrong, got=\n%v", res)
	}
}

func TestScalingVector(t *testing.T) {
	transform := matrix.Scaling(2, 3, 4)
	v := tuple.Vector(-4, 6, 8)
	res := transform.MulTuple(v)

	if !res.Equal(tuple.Vector(-8, 18, 32)) {
		t.Errorf("Scaling result is wrong, got=\n%v", res)
	}
}

func TestScalingReverse(t *testing.T) {
	transform := matrix.Scaling(2, 3, 4)
	inv := transform.Inverse()
	v := tuple.Vector(-4, 6, 8)
	res := inv.MulTuple(v)

	if !res.Equal(tuple.Vector(-2, 2, 2)) {
		t.Errorf("Scaling result is wrong, got=\n%v", res)
	}
}

func TestScalingReflection(t *testing.T) {
	transform := matrix.Scaling(-1, 1, 1)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(-2, 3, 4)) {
		t.Errorf("Scaling result is wrong, got=\n%v", res)
	}
}

func TestRotationX(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := matrix.RotationX(math.Pi / 4.0)
	fullQuarter := matrix.RotationX(math.Pi / 2.0)

	res := halfQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(0, math.Sqrt(2)/2.0, math.Sqrt(2)/2.0)) {
		t.Errorf("Rotation half quarter result is wrong, got=\n%v", res)
	}

	res = fullQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(0, 0, 1)) {
		t.Errorf("Rotation full quarter result is wrong, got=\n%v", res)
	}
}

func TestRotationXReverse(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := matrix.RotationX(math.Pi / 4.0)
	inv := halfQuarter.Inverse()

	res := inv.MulTuple(p)
	if !res.Equal(tuple.Point(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)) {
		t.Errorf("Rotation half quarter result is wrong, got=\n%v", res)
	}
}

func TestRotationY(t *testing.T) {
	p := tuple.Point(0, 0, 1)
	halfQuarter := matrix.RotationY(math.Pi / 4.0)
	fullQuarter := matrix.RotationY(math.Pi / 2.0)

	res := halfQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(math.Sqrt(2)/2.0, 0, math.Sqrt(2)/2.0)) {
		t.Errorf("Rotation half quarter result is wrong, got=\n%v", res)
	}

	res = fullQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(1, 0, 0)) {
		t.Errorf("Rotation full quarter result is wrong, got=\n%v", res)
	}
}

func TestRotationZ(t *testing.T) {
	p := tuple.Point(0, 1, 0)
	halfQuarter := matrix.RotationZ(math.Pi / 4.0)
	fullQuarter := matrix.RotationZ(math.Pi / 2.0)

	res := halfQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(-math.Sqrt(2)/2.0, math.Sqrt(2)/2.0, 0)) {
		t.Errorf("Rotation half quarter result is wrong, got=\n%v", res)
	}

	res = fullQuarter.MulTuple(p)
	if !res.Equal(tuple.Point(-1, 0, 0)) {
		t.Errorf("Rotation full quarter result is wrong, got=\n%v", res)
	}
}

func TestShearXY(t *testing.T) {
	transform := matrix.Shear(1, 0, 0, 0, 0, 0)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(5, 3, 4)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestShearXZ(t *testing.T) {
	transform := matrix.Shear(0, 1, 0, 0, 0, 0)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(6, 3, 4)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestShearYX(t *testing.T) {
	transform := matrix.Shear(0, 0, 1, 0, 0, 0)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(2, 5, 4)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestShearYZ(t *testing.T) {
	transform := matrix.Shear(0, 0, 0, 1, 0, 0)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(2, 7, 4)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestShearZX(t *testing.T) {
	transform := matrix.Shear(0, 0, 0, 0, 1, 0)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(2, 3, 6)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestShearZY(t *testing.T) {
	transform := matrix.Shear(0, 0, 0, 0, 0, 1)
	p := tuple.Point(2, 3, 4)
	res := transform.MulTuple(p)

	if !res.Equal(tuple.Point(2, 3, 7)) {
		t.Errorf("Shear result is wrong, got=\n%v", res)
	}
}

func TestCombineSequence(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := matrix.RotationX(math.Pi / 2.0)
	B := matrix.Scaling(5, 5, 5)
	C := matrix.Translation(10, 5, 7)

	p2 := A.MulTuple(p)
	if !p2.Equal(tuple.Point(1, -1, 0)) {
		t.Errorf("Combine sequence p2 is wrong, got=\n%v", p2)
	}

	p3 := B.MulTuple(p2)
	if !p3.Equal(tuple.Point(5, -5, 0)) {
		t.Errorf("Combine sequence p3 is wrong, got=\n%v", p3)
	}

	p4 := C.MulTuple(p3)
	if !p4.Equal(tuple.Point(15, 0, 7)) {
		t.Errorf("Combine sequence p4 is wrong, got=\n%v", p4)
	}
}

func TestCombineChain(t *testing.T) {
	p := tuple.Point(1, 0, 1)
	A := matrix.RotationX(math.Pi / 2.0)
	B := matrix.Scaling(5, 5, 5)
	C := matrix.Translation(10, 5, 7)
	T := C.Mul(B.Mul(A))
	res := T.MulTuple(p)

	if !res.Equal(tuple.Point(15, 0, 7)) {
		t.Errorf("Combine chain result is wrong, got=\n%v", res)
	}
}

func TestViewDefault(t *testing.T) {
	from := tuple.Point(0, 0, 0)
	to := tuple.Point(0, 0, -1)
	up := tuple.Vector(0, 1, 0)

	T := matrix.ViewTransform(from, to, up)

	if !matrix.Identity44().Equal(T) {
		t.Errorf("view transform wrong")
	}
}

func TestViewPositiveZ(t *testing.T) {
	from := tuple.Point(0, 0, 0)
	to := tuple.Point(0, 0, 1)
	up := tuple.Vector(0, 1, 0)

	T := matrix.ViewTransform(from, to, up)

	if !matrix.Scaling(-1, 1, -1).Equal(T) {
		t.Errorf("view transform wrong")
	}
}

func TestViewMovesWorld(t *testing.T) {
	from := tuple.Point(0, 0, 8)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)

	T := matrix.ViewTransform(from, to, up)

	if !matrix.Translation(0, 0, -8).Equal(T) {
		t.Errorf("view transform wrong")
	}
}

func TestViewArbitrary(t *testing.T) {
	from := tuple.Point(1, 3, 2)
	to := tuple.Point(4, -2, 8)
	up := tuple.Vector(1, 1, 0)

	T := matrix.ViewTransform(from, to, up)
	expected := matrix.Mat44{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.00000,
		0.00000, 0.00000, 0.00000, 1.00000,
	}
	if !expected.Equal(T) {
		t.Errorf("view transform wrong")
	}
}
