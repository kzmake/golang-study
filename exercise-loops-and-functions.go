package main

import (
	"fmt"
	"math"
)

func isChecked(a float64, b float64) bool {
	if math.Abs(a-b) < 1e-10 {
		return true
	}

	return false
}

func Sqrt(x float64) (float64, int) {
	var z float64 = 1.0

	for i := 0; ; i++ {
		_z := z
		z = z - 0.5*(z*z-x)/z

		if isChecked(z, _z) {
			return z, i + 1
		}
	}
}

func main() {
	msqrt := math.Sqrt(2)
	sqrt, ite := Sqrt(2)

	fmt.Println(ite, "Iterations")
	fmt.Println("math.Sqrt(2) :\t", msqrt)
	fmt.Println("Sqrt(2)      :\t", sqrt)
	fmt.Println("  -> Error (Sqrt(2) - math.Sqrt(2)) :", sqrt-msqrt)
}
