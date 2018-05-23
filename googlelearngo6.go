package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastoccupied := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastoccupied[ch]; ok && lastI > start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastoccupied[ch] = i
	}
	return maxLength
}

func main() {
	str1 := "abasdfafasasdfdfioasdfjwasdfasdfsdfasdfasdfasdfasdf"
	l := lengthOfNonRepeatingSubStr(str1)
	fmt.Println(l)

}

//lastoccured := make(map[byte]int)
//start := 0
//maxLength := 0
//for i, ch := range []byte(s) {
//if lastI, ok := lastoccured[ch]; ok && lastI >= start {
//start = lastI + 1
//}
//if i-start+1 > maxLength {
//maxLength = i - start + 1
//}
//lastoccured[ch] = i
//}
