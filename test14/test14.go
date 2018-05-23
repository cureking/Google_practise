package main

import "fmt"

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	find := func(left, right int) string {
		for left >= 0 && right < length && s[left] == s[right] {
			left--
			right++
		}
		return s[left+1 : right]
	}
	ret := ""
	for b := 0; b < length; {
		e := b
		for e < length && s[e] == s[b] {
			e++
		}
		if p := find(b-1, e); len(p) > len(ret) {
			ret = p
		}
		b = e
	}
	return ret
}
func main() {
	s := "adddddaadadasdadadadadad"
	//fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome(s))
	//fmt.Println(judge([]byte(s)))
}
