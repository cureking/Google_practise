package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	//fmt.Println(1/10)
}

func isPalindrome(x int) bool {
	m, n := x, 0
	for x > 0 {
		n = n*10 + x%10
		x = x / 10
	}
	return m == n
}