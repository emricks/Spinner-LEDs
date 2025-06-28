package errorlight

import "machine"

type ErrorLight struct {
	led machine.Pin
}

func NewErrorLight() *ErrorLight {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &ErrorLight{led: led}
}

func (e *ErrorLight) Activate() {
	e.led.High()
}

func (e *ErrorLight) Deactivate() {
	e.led.Low()
}

func (e *ErrorLight) Toggle() {
	if e.led.Get() {
		e.led.Low()
	} else {
		e.led.High()
	}
}
