package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	// Parse the input path from the flag
	inputPath := flag.String("input", "", "Path to JPEG image")
	flag.Parse()

	if *inputPath == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run image-convert.go --input image.jpg")
		os.Exit(1)
	}

	// Open the image
	f, err := os.Open(*inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Decode the image
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	fmt.Printf("var imgWidth = %d\n", width)
	fmt.Printf("var imgHeight = %d\n", height)
	fmt.Printf("var imgData = []color.RGBA{\n")

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Printf("    {R: %d, G: %d, B: %d, A: %d},\n", r>>8, g>>8, b>>8, a>>8)
		}
	}

	fmt.Println("}")
}
