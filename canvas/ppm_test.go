package canvas_test

import (
	"bytes"
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/canvas"
	"github.com/NHollmann/Raytracer-Challenge-Go/color"
)

func TestCanvasSavePpm(t *testing.T) {
	canv := canvas.New(5, 3)
	canv.SetPixel(0, 0, color.New(1.5, 0, 0))
	canv.SetPixel(2, 1, color.New(0, 0.5, 0))
	canv.SetPixel(4, 2, color.New(-0.5, 0, 1))

	var buf bytes.Buffer
	canv.SavePpm(&buf)
	res := buf.String()

	expected := "P3\n5 3\n255\n" +
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n" +
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 \n" +
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 \n"
	if res != expected {
		t.Errorf("Invalid canvas export, got=\n%s", res)
	}
}
