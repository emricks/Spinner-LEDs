package display

import (
	"image"
	"image/color"
	"machine"
	"neoblade/internal/colors"
	"tinygo.org/x/drivers/ssd1331"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

type Display struct {
	screen ssd1331.Device
}

func NewDisplay() Display {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
	})
	screen := ssd1331.New(machine.SPI1, machine.D8, machine.D9, machine.D10)
	screen.Configure(ssd1331.Config{})

	display := Display{screen}
	display.Clear()
	return display
}

func (d *Display) Clear() {
	d.screen.FillScreen(colors.BLACK)
}

func (d *Display) WriteText(s string) error {
	d.Clear()
	tinyfont.WriteLine(&d.screen, &freemono.Regular9pt7b, 0, 10, s, colors.WHITE)
	err := d.screen.Display()
	if err != nil {
		return err
	}
	return nil
}

func (d *Display) ShowImage(img *image.RGBA) {
	//d.Clear()
	maxPoint := img.Rect.Size()
	imgColors := ImageToColorRGBA(img)

	_ = d.screen.FillRectangleWithBuffer(0, 0, int16(maxPoint.X), int16(maxPoint.Y), imgColors)
	return
}

func ImageToColorRGBA(img *image.RGBA) []color.RGBA {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	out := make([]color.RGBA, w*h)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := y*w + x
			r, g, b, a := img.At(x+img.Bounds().Min.X, y+img.Bounds().Min.Y).RGBA()
			out[i] = color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
		}
	}

	return out
}
