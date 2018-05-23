package main

import (
	"math"
	"math/cmplx"
	"fmt"
)
func main() {
	fmt.Printf("%.3f",
		cmplx.Exp(1i*math.Pi+1))
	//cmplx.Pow(math.E,1i*math.Pi))
}
