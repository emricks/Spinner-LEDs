package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/jpeg"
	"neoblade/internal/drawing"
	"neoblade/internal/motor"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type TemplateData struct {
	Package string
	VarName string
	Data    map[float64][]color.RGBA
	Size    int
}

const convertStepsPerRevolution = 816

func main() {
	inputPath := flag.String("input", "", "Path to JPEG image")
	flag.Parse()

	if *inputPath == "" {
		_, err := fmt.Fprintln(os.Stderr, "Usage: go run image-convert.go --input image.jpg")
		if err != nil {
			panic(err)
		}
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

	inputImgSize := img.Bounds().Size().Y
	divisor := motor.CalculateDivisor(convertStepsPerRevolution, inputImgSize)
	images := map[float64][]color.RGBA{}
	fmt.Printf("input size: %d, steps per rev: %d, divisor: %d\n", inputImgSize, convertStepsPerRevolution, divisor)
	rads := motor.CalculatePossibleRads(convertStepsPerRevolution, divisor)
	fmt.Println(len(rads))
	for _, rad := range rads {
		x0, y0, x1, y1 := drawing.FindEndpoints(inputImgSize, rad)
		images[rad] = drawing.ExtractLinePixels(img, x0, y0, x1, y1)
	}

	// create template
	varName := filenameToVarName(*inputPath)
	tmplData := TemplateData{
		Package: "images",
		VarName: varName,
		Data:    images,
		Size:    inputImgSize,
	}

	tmpl, err := template.ParseFiles("rgba_map.tmpl")
	if err != nil {
		panic(err)
	}

	out, err := os.Create(fmt.Sprintf("internal/images/%s.go", varName))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, tmplData); err != nil {
		panic(err)
	}
}

func filenameToVarName(filename string) string {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// Split on any non-alphanumeric character
	segments := strings.FieldsFunc(name, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Capitalize each segment
	var varName string
	for _, seg := range segments {
		if len(seg) == 0 {
			continue
		}
		varName += strings.ToUpper(seg[:1]) + seg[1:]
	}

	return fmt.Sprintf("%sLines", varName)
}
