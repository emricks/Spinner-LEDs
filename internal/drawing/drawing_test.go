package drawing

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddWithTable(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -7, -12},
		{123, 456, 579},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestPosToRad(t *testing.T) {
	result := PosToRad(204, 816)
	expected := math.Pi / 2
	if result != expected {
		t.Errorf("PosToRad(204, 816) = %d; want %d", result, expected)
	}
}

func TestFindEndpoints(t *testing.T) {

	x0, y0, x1, y1 := FindEndpoints(100, math.Pi/3)
	if x0 != int(math.Round(50*(math.Sqrt(3)/2+1))) {
		t.Errorf("x0 = %d; want 93", x0)
	}
	if y0 != 75 {
		t.Errorf("y0 = %d; want 75", y0)
	}
	if x1 != int(math.Round(100-50*(math.Sqrt(3)/2+1))) {
		t.Errorf("x1 = %d; want 7", x1)
	}
	if y1 != 25 {
		t.Errorf("y1 = %d; want 25", y1)
	}
}
