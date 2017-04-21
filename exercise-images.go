package main

import (
	"image"
	"image/color"

	"code.google.com/p/go-tour/pic"
)

const (
	RED = iota
	GREEN
	BLUE
	ALPHA
)

type Image struct {
	pix    []uint8
	width  int
	height int
}

func NewImage(width, height int) *Image {
	pix := make([]uint8, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			offset := y*width*4 + x*4
			pix[offset+RED] = uint8(x ^ y)
			pix[offset+GREEN] = uint8((x + y) / 2)
			pix[offset+BLUE] = 255
			pix[offset+ALPHA] = 255
		}
	}
	return &Image{pix, width, height}
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	offset := y*img.width*4 + x*4
	r := img.pix[offset+RED]
	g := img.pix[offset+GREEN]
	b := img.pix[offset+BLUE]
	a := img.pix[offset+ALPHA]
	return color.RGBA{r, g, b, a}
}

func main() {
	m := NewImage(256, 64)
	pic.ShowImage(m)
}
