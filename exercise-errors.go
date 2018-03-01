package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	v float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.v)
}


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrNegativeSqrt{x}
	}
	z := x / 2
	for i:=0; i<10; i++ {
		pre_z := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(pre_z - z) < 0.00000000001 {
			break
		}
	}
	return z, nil
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
