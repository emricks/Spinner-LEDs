package images

import (
	"image/color"
)

type ImageLines struct {
	Size int
	Data map[float64][]color.RGBA
}
