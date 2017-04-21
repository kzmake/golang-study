package main

import (
	"fmt"
	"math"
)

// constにできないのでしかたなくvar
var Epsilon = math.Nextafter(1, 2) - 1

func Sqrt(x float64) float64 {
	z := 1.0

	println(Epsilon)
	for {
		preZ := z
		z = z - 0.5*(z*z-x)/z

		if math.Abs(z-preZ) <= Epsilon {
			return z
		}
	}
}

func main() {
	msqrt := math.Sqrt(2)
	sqrt := Sqrt(2)

	fmt.Println("math.Sqrt(2) :\t", msqrt)
	fmt.Println("Sqrt(2)      :\t", sqrt)
	fmt.Println("  -> Error (Sqrt(2) - math.Sqrt(2)) :", sqrt-msqrt)
}
