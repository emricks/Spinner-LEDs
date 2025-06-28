package drawing

import "testing"

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
