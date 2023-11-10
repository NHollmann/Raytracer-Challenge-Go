package camera_test

import (
	"math"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/camera"
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

func TestCameraConstructor(t *testing.T) {
	c := camera.New(160, 120, math.Pi/2.0)
	if c.Hsize != 160 {
		t.Errorf("camera hsize incorrect")
	}
	if c.Vsize != 120 {
		t.Errorf("camera vsize incorrect")
	}
	if !flt.Equal(c.FieldOfView, math.Pi/2.0) {
		t.Errorf("camera fov incorrect")
	}
	if !c.Transform.Equal(matrix.Identity44()) {
		t.Errorf("camera transform incorrect")
	}
}

func TestPixelSizeHorizontal(t *testing.T) {
	c := camera.New(200, 125, math.Pi/2.0)
	if !flt.Equal(c.PixelSize, 0.01) {
		t.Errorf("pixel size incorrect")
	}
}

func TestPixelSizeVertical(t *testing.T) {
	c := camera.New(125, 200, math.Pi/2.0)
	if !flt.Equal(c.PixelSize, 0.01) {
		t.Errorf("pixel size incorrect")
	}
}

func TestCenterRay(t *testing.T) {
	c := camera.New(201, 101, math.Pi/2.0)
	r := c.RayForPixel(100, 50)
	if !r.Origin.Equal(tuple.Point(0, 0, 0)) {
		t.Errorf("origin incorrect")
	}
	if !r.Direction.Equal(tuple.Vector(0, 0, -1)) {
		t.Errorf("direction incorrect")
	}
}

func TestCornerRay(t *testing.T) {
	c := camera.New(201, 101, math.Pi/2.0)
	r := c.RayForPixel(0, 0)
	if !r.Origin.Equal(tuple.Point(0, 0, 0)) {
		t.Errorf("origin incorrect")
	}
	if !r.Direction.Equal(tuple.Vector(0.66519, 0.33259, -0.66851)) {
		t.Errorf("direction incorrect")
	}
}

func TestCenterTransformedRay(t *testing.T) {
	c := camera.New(201, 101, math.Pi/2.0)
	c.Transform = matrix.RotationY(math.Pi / 4.0).Mul(matrix.Translation(0, -2, 5))
	r := c.RayForPixel(100, 50)
	if !r.Origin.Equal(tuple.Point(0, 2, -5)) {
		t.Errorf("origin incorrect")
	}
	if !r.Direction.Equal(tuple.Vector(math.Sqrt(2)/2.0, 0, -math.Sqrt(2)/2.0)) {
		t.Errorf("direction incorrect")
	}
}

func TestRender(t *testing.T) {
	w := intersection.NewDefaultWorld()
	c := camera.New(11, 11, math.Pi/2.0)
	from := tuple.Point(0, 0, -5)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)
	c.Transform = matrix.ViewTransform(from, to, up)
	image := c.Render(w)
	if !image.PixelAt(5, 5).Equal(color.New(0.38066, 0.47583, 0.2855)) {
		t.Errorf("color incorrect")
	}
}
