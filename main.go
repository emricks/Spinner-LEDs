package main

import (
	"image"
	"math"
	"neoblade/internal/display"
	"neoblade/internal/drawing"
	"neoblade/internal/errorlight"
	"neoblade/internal/motor"
)

var eLight = errorlight.NewErrorLight()
var myDisplay = display.NewDisplay()
var encoder *motor.Encoder
var myMotor = motor.NewMotor()

const imgSize = 32
const stepsPerRevolution = 816

func main() {
	divisor := int(math.Round(stepsPerRevolution / (imgSize * math.Pi)))
	encoder, _ = motor.NewEncoder(stepsPerRevolution, divisor)
	images := map[float64]*image.RGBA{}
	for _, rad := range encoder.PossibleRadians {
		x0, y0, x1, y1 := drawing.FindEndpoints(imgSize, rad)
		images[rad] = drawing.NewLineImage(imgSize, x0, y0, x1, y1)
	}
	//lastPos := uint(0)
	for {
		select {
		case pos := <-encoder.PositionChannel:

			rad := drawing.PosToRad(int(pos), int(encoder.StepsPerRevolution))
			//err := myDisplay.WriteText(fmt.Sprintf("%d,%d,%d,%d", x0, y0, x1, y1))

			myDisplay.ShowImage(images[rad])

		default:
			// nothing
		}
	}
}
