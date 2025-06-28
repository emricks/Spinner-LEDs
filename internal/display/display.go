package display

import (
	"machine"
	"neoblade/internal/colors"
	"tinygo.org/x/drivers/i2csoft"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

type Display struct {
	screen ssd1306.Device
}

func NewDisplay() Display {
	i2c := i2csoft.New(machine.D21, machine.D20)
	i2c.Configure(i2csoft.I2CConfig{})

	screen := ssd1306.NewI2C(i2c)
	screen.Configure(ssd1306.Config{Address: 0x3C})
	screen.ClearDisplay()
	//screen.Configure(ssd1306.Config{Width: 64, Height: 16, Address: 0x3C})

	display := Display{screen}
	return display
}

func (d *Display) WriteText(s string) error {
	d.screen.ClearDisplay()
	tinyfont.WriteLine(&d.screen, &freemono.Regular9pt7b, 0, 10, s, colors.WHITE)
	err := d.screen.Display()
	if err != nil {
		return err
	}
	return nil
}

func (d *Display) ShowImage(img pixel.Image[pixel.Monochrome]) {
	d.screen.ClearDisplay()
	d.screen.DrawBitmap(0, 0, img)
	d.screen.Display()
}
