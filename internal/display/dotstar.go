package display

import (
	"image/color"
	"machine"
	"tinygo.org/x/drivers/apa102"
)

type DotStar struct {
	device *apa102.Device
	length int
}

func NewDotStar() DotStar {
	spi := machine.SPI1
	spi.Configure(machine.SPIConfig{
		Frequency: 16000000,
	})
	dot := DotStar{apa102.New(spi), 72}
	return dot
}

func (dot DotStar) Display(colors []color.RGBA) {
	for len(colors) != dot.length {
		colors = append(colors, color.RGBA{})
	}
	dot.device.WriteColors(colors)
}
