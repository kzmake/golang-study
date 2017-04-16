package main

import (
	"fmt"
	"math"
)

const (
	CONVERGENCE_VALUE = 1e-10
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// %gは%vと違い、暗黙的にfloat64型へキャストされてるっぽいので無限ループにならない
	// ただ、%g -> %v の変更の可能性と気持ち悪いので、float64(e)と明示的にキャスト
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1.0)
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
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))

	// Error Tests
	fmt.Println(Sqrt(-2))
}
