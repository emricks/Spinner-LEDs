package motor

import (
	"machine"
	"time"
)

type Motor struct {
	pwmPin machine.Pin
	freqHz float32
}

func NewMotor() *Motor {
	m := &Motor{}
	m.pwmPin = machine.D23
	m.pwmPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.freqHz = 1000

	return m
}

func (m *Motor) Run(dutyCycle float32) {
	go func() {
		period := time.Second / time.Duration(m.freqHz)
		onTime := time.Duration(float32(period) * dutyCycle)
		offTime := period - onTime

		for {
			m.pwmPin.High()
			time.Sleep(onTime)

			m.pwmPin.Low()
			time.Sleep(offTime)
		}
	}()
}
