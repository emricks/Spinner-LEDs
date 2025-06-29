package main

import (
	"neoblade/internal/display"
	"neoblade/internal/drawing"
	"neoblade/internal/errorlight"
	"neoblade/internal/images"
	"neoblade/internal/motor"
)

var eLight = errorlight.NewErrorLight()
var myDisplay = display.NewDisplay()
var encoder *motor.Encoder
var myMotor = motor.NewMotor()

const stepsPerRevolution = 816

func main() {
	renderImage := images.Enky64Lines
	divisor := motor.CalculateDivisor(stepsPerRevolution, renderImage.Size)
	encoder, _ = motor.NewEncoder(stepsPerRevolution, divisor)
	//lastPos := uint(0)
	for {
		select {
		case pos := <-encoder.PositionChannel:

			rad := drawing.PosToRad(pos, encoder.StepsPerRevolution)
			x0, y0, x1, y1 := drawing.FindEndpoints(renderImage.Size, rad)
			//err := myDisplay.WriteText(fmt.Sprintf("%d,%d,%d,%d", x0, y0, x1, y1))
			myDisplay.Clear()
			myDisplay.DrawAngledLine(x0, y0, x1, y1, renderImage.Data[rad])
		default:
			// nothing
		}
	}
}
