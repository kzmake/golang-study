package main

import (
	"fmt"
	"math"
)

const (
	CONVERGENCE_VALUE = 1e-10
)

type ErrNegativeSqrt float64

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
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
