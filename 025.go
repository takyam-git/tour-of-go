package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    prev := 0.0
    for {
        z = z - (z*z - x) / (x * 2)
        if math.Abs(z - prev) < 0.0000000000000001 {
		    return z
        }
        prev = z
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
