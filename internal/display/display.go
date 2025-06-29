package display

import (
	"image"
	"image/color"
	"machine"
	"neoblade/internal/colors"
	"neoblade/internal/drawing"
	"time"
	"tinygo.org/x/drivers/ssd1331"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

type Display struct {
	screen ssd1331.Device
}

func NewDisplay() Display {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 16000000,
	})
	screen := ssd1331.New(machine.SPI1, machine.D8, machine.D9, machine.D10)
	screen.Configure(ssd1331.Config{})

	display := Display{screen}
	display.Clear()
	return display
}

func (d *Display) Clear() {
	//d.screen.FillScreen(colors.BLACK)
	d.screen.Tx([]byte{
		0x26, 0x01, 0x22,
		0x00, 0x00, // x0, y0
		0x5F, 0x3F, // x1, y1
		0x00, 0x00, 0x00, // outline color = black
		0x00, 0x00, 0x00, // fill color = black
	}, true)
	time.Sleep(1 * time.Millisecond)
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

func (d *Display) DrawAngledLine(x0, y0, x1, y1 int, colors []color.RGBA) {
	dx := drawing.Abs(x1 - x0)
	dy := -drawing.Abs(y1 - y0)
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy

	i := 0
	for {
		if i >= len(colors) {
			break // Prevent overflow
		}
		d.screen.SetPixel(int16(x0), int16(y0), colors[i])
		i++

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}
