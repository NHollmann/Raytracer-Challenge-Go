package flt_test

import (
	"testing"

	"nicolashollmann.de/raytracer-challange/flt"
)

func TestEqual(t *testing.T) {
	if flt.Equal(0.0, 0.0) == false {
		t.Errorf("0.0 does not equal itself")
	}

	if flt.Equal(1.2345, 1.2345) == false {
		t.Errorf("1.2345 does not equal itself")
	}

	if flt.Equal(0.0, 0.1) == true {
		t.Errorf("0.0 does equal 0.1")
	}

	if flt.Equal(7.1, 23.6) == true {
		t.Errorf("7.1 does equal 23.6")
	}
}
