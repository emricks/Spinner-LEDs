package motor

import (
	"reflect"
	"testing"
)

func TestCalculatePossibleRads(t *testing.T) {
	result := calculatePossibleRads(816, 4)
	expected := []float32{3}
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("have %d; want %d", result, expected)
	}
}
