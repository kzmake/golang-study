package main

import (
	"fmt"
	"math"
)

const (
	CONVERGENCE_VALUE = 1e-10
)

func Sqrt(x float64) float64 {
	z := 1.0
	is_checked := func(a, b float64) bool {
		if math.Abs(a-b) < CONVERGENCE_VALUE {
			return true
		}
		return false
	}

	for {
		pre_z := z
		z = z - 0.5*(z*z-x)/z

		if is_checked(z, pre_z) {
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
