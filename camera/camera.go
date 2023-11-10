package camera

import (
	"math"

	"github.com/NHollmann/Raytracer-Challenge-Go/canvas"
	"github.com/NHollmann/Raytracer-Challenge-Go/intersection"
	"github.com/NHollmann/Raytracer-Challenge-Go/matrix"
	"github.com/NHollmann/Raytracer-Challenge-Go/ray"
	"github.com/NHollmann/Raytracer-Challenge-Go/tuple"
)

type Camera struct {
	Hsize       uint32
	Vsize       uint32
	FieldOfView float64
	Transform   matrix.Mat44
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
}

func New(hsize, vsize uint32, fieldOfView float64) Camera {
	halfView := math.Tan(fieldOfView / 2.0)
	aspect := float64(hsize) / float64(vsize)

	var halfWidth float64
	var HalfHeight float64
	if aspect >= 1 {
		halfWidth = halfView
		HalfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		HalfHeight = halfView
	}

	return Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		FieldOfView: fieldOfView,
		Transform:   matrix.Identity44(),
		HalfWidth:   halfWidth,
		HalfHeight:  HalfHeight,
		PixelSize:   (halfWidth * 2) / float64(hsize),
	}
}

func (c *Camera) RayForPixel(px, py float64) ray.Ray {
	xoffset := (px + 0.5) * c.PixelSize
	yoffset := (py + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xoffset
	worldY := c.HalfHeight - yoffset

	pixel := c.Transform.Inverse().MulTuple(tuple.Point(worldX, worldY, -1))
	origin := c.Transform.Inverse().MulTuple(tuple.Point(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()

	return ray.New(origin, direction)
}

func (c *Camera) Render(w intersection.World) canvas.Canvas {
	image := canvas.New(c.Hsize, c.Vsize)

	for y := uint32(0); y < c.Vsize; y++ {
		for x := uint32(0); x < c.Hsize; x++ {
			pixelRay := c.RayForPixel(float64(x), float64(y))
			pixelColor := w.ColorAt(pixelRay)
			image.SetPixel(x, y, pixelColor)
		}
	}

	return image
}
