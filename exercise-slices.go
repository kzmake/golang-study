package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pix := make([][]uint8, dy)
	for y := range pix {
		pix[y] = make([]uint8, dx)
	}

	for y := range pix {
		for x := range pix[y] {
			pix[y][x] = uint8((x + y) / 2)
		}
	}

	return pix
}

func main() {
	pic.Show(Pic)
}
