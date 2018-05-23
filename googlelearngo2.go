package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func sliceprint(s []int) {
	fmt.Printf("%v , length:%d , capacity:%d\n", s, len(s), cap(s))
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{4, 5, 6, 7, 8, 9, 90}
	fmt.Println(arr1, arr2, arr3)
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	for _, i := range arr3 {
		fmt.Println(i)
	}
	s1 := arr3[0:3]
	s2 := s1[2:6]
	fmt.Println(s2, cap(s1), cap(s2))

	var s5 = []int{2, 3, 4, 5, 6}
	s6 := []int{1, 2, 3, 4, 5}
	s7 := make([]int, 10, 32)
	s7[2] = 12
	sliceprint(s7)
	fmt.Println(s5, s6, s7)

	copy(s7, s2)
	sliceprint(s7)
	fmt.Println(s7)

	s7 = append(s7[:3], s7[4:]...)
	sliceprint(s7)
	fmt.Println(s7)

	//pop from front
	front := s7[0]
	s7 = s7[1:]
	sliceprint(s7)
	//pop from back
	tail := s7[len(s7)-1]
	s7 = s7[:len(s7)-1]
	fmt.Println(front, tail)
	sliceprint(s7)
}
