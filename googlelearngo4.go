package main

import "fmt"

func div1(str string) (l int, result string) {
	arr1 := []byte(str)

	maxlenth := 0
	maxstring := arr1[0:0]

	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			smaxstring := arr1[i:j]
			if !exist(smaxstring, arr1[j]) {
				if maxlenth < j-i+1 {
					maxlenth = j - i + 1
					maxstring = arr1[i : j+1]
				}
				continue
			} else {
				break
			}
		}
	}
	return maxlenth, string(maxstring)
}

func exist(arr []byte, value byte) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func main() {
	str1 := "abasdfafasasdfdfioasdfjwasdfasdfsdfasdfasdfasdfasdf"
	s, l := div1(str1)
	fmt.Println(s, l)
}
