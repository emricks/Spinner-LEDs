package main

import (
	"neoblade/internal/display"
	"neoblade/internal/drawing"
	"neoblade/internal/errorlight"
	"neoblade/internal/motor"
)

var eLight = errorlight.NewErrorLight()
var myDisplay = display.NewDisplay()
var encoder *motor.Encoder
var myMotor = motor.NewMotor()

const imgSize = 64

func main() {
	encoder, _ = motor.NewEncoder(4)
	//lastPos := uint(0)
	for {
		select {
		case pos := <-encoder.PositionChannel:

			rad := drawing.PosToRad(int(pos), int(encoder.StepsPerRevolution))
			x0, y0, x1, y1 := drawing.FindEndpoints(imgSize-1, rad)
			//err := myDisplay.WriteText(fmt.Sprintf("%d,%d,%d,%d", x0, y0, x1, y1))

			myDisplay.ShowImage(drawing.NewLineImage(imgSize, x0, y0, x1, y1))

		default:
			// nothing
		}
	}
}
