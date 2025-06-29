package drawing

import (
	"image"
	"image/color"
	"math"
	"neoblade/internal/colors"
)

func NewLineImage(size, x0, y0, x1, y1 int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	DrawLineInImg(img, x0, y0, x1, y1)
	return img
}

func PosToRad(pos int, maxPos int) float64 {
	return 2 * math.Pi * float64(pos) / float64(maxPos)
}

func FindEndpoints(size int, angle float64) (int, int, int, int) {
	size -= 1
	radius := float64(size) / 2
	converted := math.Pi/2 - angle
	x0 := radius * (math.Cos(converted) + 1)
	y0 := radius * (math.Sin(converted) + 1)
	x1 := float64(size) - x0
	y1 := float64(size) - y0
	return int(math.Round(x0)), int(math.Round(y0)), int(math.Round(x1)), int(math.Round(y1))
}

func DrawLineInImg(img *image.RGBA, x0, y0, x1, y1 int) {
	dx := Abs(x1 - x0)
	dy := -Abs(y1 - y0)
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
		img.Set(x0, y0, colors.GREEN)
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

func ExtractLinePixels(img image.Image, x0, y0, x1, y1 int) []color.RGBA {
	var result []color.RGBA

	dx := Abs(x1 - x0)
	dy := -Abs(y1 - y0)
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy

	for {
		r, g, b, a := img.At(x0, y0).RGBA()
		result = append(result, color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		})

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

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
