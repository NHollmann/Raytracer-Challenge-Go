package color_test

import (
	"testing"

	"github.com/NHollmann/Raytracer-Challenge-Go/color"
	"github.com/NHollmann/Raytracer-Challenge-Go/flt"
)

func TestColorConstructor(t *testing.T) {
	c := color.New(-0.5, 0.4, 1.7)
	if !flt.Equal(c.R(), -0.5) {
		t.Errorf("r is not equal -0.5")
	}
	if !flt.Equal(c.G(), 0.4) {
		t.Errorf("g is not equal 0.4")
	}
	if !flt.Equal(c.B(), 1.7) {
		t.Errorf("b is not equal 1.7")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := color.New(0.9, 0.6, 0.75)
	c2 := color.New(0.7, 0.1, 0.25)
	if !c1.Add(c2).Equal(color.New(1.6, 0.7, 1.0)) {
		t.Errorf("color add wrong")
	}
}

func TestColorSub(t *testing.T) {
	c1 := color.New(0.9, 0.6, 0.75)
	c2 := color.New(0.7, 0.1, 0.25)
	if !c1.Sub(c2).Equal(color.New(0.2, 0.5, 0.5)) {
		t.Errorf("color sub wrong")
	}
}

func TestColorMulScalar(t *testing.T) {
	c := color.New(0.2, 0.3, 0.4)
	if !c.MulScalar(2).Equal(color.New(0.4, 0.6, 0.8)) {
		t.Errorf("color scalar multiply wrong")
	}
}

func TestColorMulColor(t *testing.T) {
	c1 := color.New(1.0, 0.2, 0.4)
	c2 := color.New(0.9, 1.0, 0.1)
	if !c1.MulColor(c2).Equal(color.New(0.9, 0.2, 0.04)) {
		t.Errorf("color-color multiply wrong")
	}
}
