//go:build teensy41 || teensy40 || mimxrt1062

package motor

import (
	"machine"
)

type Encoder struct {
	pinA               machine.Pin
	pinB               machine.Pin
	position           int
	lastPosition       uint8
	PositionChannel    chan int
	StepsPerRevolution int
	stepDivisor        int
	PossibleRadians    []float64
}

func NewEncoder(stepsPerRevolution, stepDivisor int) (*Encoder, error) {
	a := machine.D18
	b := machine.D19

	a.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	b.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	e := &Encoder{
		pinA:               a,
		pinB:               b,
		position:           0,
		lastPosition:       0,
		PositionChannel:    make(chan int, 512),
		StepsPerRevolution: stepsPerRevolution,
		stepDivisor:        stepDivisor,
		PossibleRadians:    CalculatePossibleRads(stepsPerRevolution, stepDivisor),
	}

	err := a.SetInterrupt(machine.PinToggle, e.handleInterrupt)
	if err != nil {
		return nil, err
	}

	err = b.SetInterrupt(machine.PinToggle, e.handleInterrupt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Encoder) readPosition() uint8 {
	var ab uint8
	if e.pinA.Get() {
		ab |= 0x01
	}
	if e.pinB.Get() {
		ab |= 0x02
	}
	return ab
}

func (e *Encoder) handleInterrupt(pin machine.Pin) {
	ab := e.readPosition()

	transition := (e.lastPosition << 2) | ab

	switch transition {
	case 0b0001, 0b0111, 0b1110, 0b1000:
		e.position++
	case 0b0010, 0b0100, 0b1101, 0b1011:
		e.position--
		// other transitions are invalid/bounce, ignore
	}

	e.lastPosition = ab

	// handle rollover/rollback
	if e.position >= e.StepsPerRevolution {
		e.position = 0
	} else if e.position < 0 {
		e.position = e.StepsPerRevolution - 1
	}

	if e.position%e.stepDivisor == 0 {
		// perform action
		select {
		case e.PositionChannel <- e.position:
		default:
			// drop if buffer full
		}
	}
}
