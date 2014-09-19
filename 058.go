package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %.f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0.0 {
        return 0.0, ErrNegativeSqrt(x)
    }
    z := 1.0
    prev := 0.0
    for {
        z = z - (z*z - x) / (x * 2)
        if math.Abs(z - prev) < 0.00000000001 {
		    return z, nil
        }
        prev = z
    }
}


func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
