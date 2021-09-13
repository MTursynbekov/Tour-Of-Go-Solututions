package main

import (
	"fmt";
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		temp := z
		z -= (z*z - x) / (2 * z)
		if temp == z {
			break
		}
		fmt.Print(z, " ")
	}
	fmt.Println()
	return z
}

func main() {
	fmt.Println(Sqrt(111))
	fmt.Println(math.Sqrt(111))
}