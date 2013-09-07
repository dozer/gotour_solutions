package main

import (
	"fmt"
	"math/cmplx"
)

const delta = 1e-10

func Cbrt(x complex128) complex128 {
	z := x
	for {
		n := z - (z*z*z-x)/(3*z*z)
		if cmplx.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2), cmplx.Pow(2, 1.0/3.0))
}
