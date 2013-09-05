package main

import (
	"fmt"
	"math"
)

const delta = 1e-10

/* Find the sqrt of a number x by repeating Newton's method until the
   difference is less than a very small delta */
func Sqrt(x float64) float64 {
	z := x
	for {
		n := z - (z*z-x)/(2*z)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range vals {
		fmt.Printf("%f\t%f\n", Sqrt(float64(v)), math.Sqrt(float64(v)))
	}
}
