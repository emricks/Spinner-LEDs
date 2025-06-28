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

func main() {
	encoder, _ = motor.NewEncoder()
	img := drawing.NewLineImage()
	myDisplay.ShowImage(img)

	for {
		select {
		case pos := <-encoder.PositionChannel:
			if pos == 0 {
				eLight.Toggle()
			}
		default:
			// nothing
		}
	}
}
