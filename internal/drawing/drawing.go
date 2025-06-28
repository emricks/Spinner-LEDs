package drawing

import (
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
