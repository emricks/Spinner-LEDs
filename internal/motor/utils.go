package motor

import "neoblade/internal/drawing"

func calculatePossibleRads(steps, divisor int) []float64 {
	// calculate the possible radian values given encoder "e" has e.StepsPerRevolution and e.stepDivisor
	var arr []float64
	for a := 0; a <= steps; a++ {
		if a%divisor == 0 {
			arr = append(arr, drawing.PosToRad(a, steps))
		}
	}
	return arr
}
