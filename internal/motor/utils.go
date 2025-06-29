package motor

import (
	"math"
	"neoblade/internal/drawing"
)

func CalculatePossibleRads(steps, divisor int) []float64 {
	// calculate the possible radian values given encoder "e" has e.StepsPerRevolution and e.stepDivisor
	var arr []float64
	for a := 0; a <= steps; a++ {
		if a%divisor == 0 {
			arr = append(arr, drawing.PosToRad(a, steps))
		}
	}
	return arr
}

func CalculateDivisor(steps, size int) int {
	return steps / int(float64(size)*math.Pi)
}
