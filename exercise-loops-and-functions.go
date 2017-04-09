package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    var z float64 = 1.0
    return z
}

func main() {
    fmt.Println("math.Sqrt(2) :", math.Sqrt(2))
    fmt.Println("Sqrt(2) :", Sqrt(2))
    fmt.Println("  -> Error (Sqrt(2) - math.Sqrt(2)) :", (Sqrt(2) - math.Sqrt(2)))
}

