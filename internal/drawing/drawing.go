package drawing

import (
	"math"
	"tinygo.org/x/drivers/pixel"
)

func NewLineImage() pixel.Image[pixel.Monochrome] {
	img := pixel.NewImage[pixel.Monochrome](128, 64)
	///img.FillSolidColor(pixel.NewMonochrome(255, 255, 255))
	DrawLine(&img, 0, 0, 32, 32)
	return img
}

func add(a int, b int) int {
	return a + b
}

func PosToRad(pos uint, maxpos uint) float64 {
	return 2 * math.Pi * float64(pos) / float64(maxpos)
}

func FindEndpoints(size float64, angle float64) (uint, uint, uint, uint) {
	radius := size / 2
	converted := math.Pi/2 - angle
	x0 := radius * (math.Cos(converted) + 1)
	y0 := radius * (math.Sin(converted) + 1)
	x1 := size - x0
	y1 := size - y0
	return uint(math.Round(x0)), uint(math.Round(y0)), uint(math.Round(x1)), uint(math.Round(y1))
}

func DrawLine(img *pixel.Image[pixel.Monochrome], x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	dy := -abs(y1 - y0)
	sx := -1
	sy := -1
	if x0 < x1 {
		sx = 1
	}
	if y0 < y1 {
		sy = 1
	}
	err := dx + dy

	for {
		img.Set(x0, y0, true)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
