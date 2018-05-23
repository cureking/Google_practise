package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	for x%10 == 0 {
		x /= 10
	}
	y := 0
	for x/10 != 0 {
		y *= 10
		y += x % 10
		x /= 10
	}
	y *= 10
	y += x % 10
	if y < math.MinInt32 || y > math.MaxInt32 {
		return 0
	}
	return y
}

func main() {
	fmt.Println(reverse(0))
}
