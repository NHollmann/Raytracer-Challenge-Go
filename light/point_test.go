package light_test

import (
	"testing"

	"nicolashollmann.de/raytracer-challange/color"
	"nicolashollmann.de/raytracer-challange/light"
	"nicolashollmann.de/raytracer-challange/tuple"
)

func TestPointLightConstructor(t *testing.T) {
	intensity := color.New(1, 1, 1)
	position := tuple.Point(0, 0, 0)
	l := light.NewPoint(position, intensity)
	if !l.Position.Equal(position) {
		t.Errorf("light position wrong")
	}
	if !l.Intensity.Equal(intensity) {
		t.Errorf("light intensity wrong")
	}
}
