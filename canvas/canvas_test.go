package canvas_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/canvas"
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
)

func TestCanvasConstructor(t *testing.T) {
	canv := canvas.New(10, 20)
	if canv.Width() != 10 {
		t.Errorf("canvas width incorrect")
	}
	if canv.Height() != 20 {
		t.Errorf("canvas height incorrect")
	}
	for y := uint32(0); y < 20; y++ {
		for x := uint32(0); x < 10; x++ {
			if !canv.PixelAt(x, y).Equal(color.New(0, 0, 0)) {
				t.Errorf("canvas not initialized black")
				t.FailNow()
			}
		}
	}
}

func TestCanvasWrite(t *testing.T) {
	canv := canvas.New(10, 20)
	red := color.New(1, 0, 0)
	canv.SetPixel(2, 3, red)
	if !canv.PixelAt(2, 3).Equal(red) {
		t.Errorf("canvas color not set correctly")
	}
}
