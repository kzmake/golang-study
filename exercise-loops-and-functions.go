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

func Sqrt(x float64) float64 {
	var z float64 = 1.0

	for {
		_z := z
		z = z - 0.5*(z*z-x)/z

		if isChecked(z, _z) {
			return z
		}
	}
}

func main() {
	fmt.Println("math.Sqrt(2) :\t", math.Sqrt(2))
	fmt.Println("Sqrt(2)      :\t", Sqrt(2))
	fmt.Println("  -> Error (Sqrt(2) - math.Sqrt(2)) :", Sqrt(2)-math.Sqrt(2))
}
