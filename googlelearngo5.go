package main

import "fmt"

func div2(s string) (l int) {
	arr1 := []byte(s)

	maxlenth := 0

	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			smaxstring := arr1[i:j]
			mark := false
			for _, v := range smaxstring {
				if arr1[j] == v {
					mark = true
				}
			}
			if !mark {
				if maxlenth < j-i+1 {
					maxlenth = j - i + 1
				}
				continue
			} else {
				break
			}
		}
	}
	return maxlenth
}

func main() {
	str1 := "abasdfafasasdfdfioasdfjwasdfasdfsdfasdfasdfasdfasdf"
	l := div2(str1)
	fmt.Println(l)
}
