package main

import (
	"fmt"
)

func main() {
	var a=[...]int{3,3}
	b:=a[:]
	fmt.Println(twoSum2(b,6))
}

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for k2, v2 := range nums[k+1:] {
			if v+v2 == target {
				return []int{k, k+k2+1}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for k2,v2:=range nums {
		comp := target - v2
		if k, ok := m[comp]; ok {
			return []int {k, k2}
		}
		m[v2] = k2
	}
	return nil
}
