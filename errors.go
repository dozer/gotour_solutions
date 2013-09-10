package main

import (
	"fmt"
	"math"
)

const delta = 1e-10

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

/* Find the sqrt of a number x by repeating Newton's method until the
   difference is less than a very small delta */
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := f
	for {
		n := z - (z*z-f)/(2*z)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
