package main

import "misc/tour/src/golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pixels := make([][]uint8, dy)
	data := make([]uint8, dx)
	for i := range pixels {
		for j := range data{
			data[j] = uint8(((i ^ 10) / (i + 20)) / 10)
		}
		pixels[i] = data
	}
	return pixels
}

func main() {
	pic.Show(Pic)
}