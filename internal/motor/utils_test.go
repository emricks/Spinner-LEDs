package motor

import (
	"reflect"
	"testing"
)

func TestCalculatePossibleRads(t *testing.T) {
	result := CalculatePossibleRads(12, 4)
	expected := []float64{0, 2.0943951023931953, 4.1887902047863905, 6.283185307179586}
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("have %d; want %d", result, expected)
	}
}
