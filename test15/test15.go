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
	result := ""
	for startLeft := 0; startLeft < length; {
		singleResult := ""
		startRight := startLeft

		for startLeft < length && startRight == startLeft {
			startRight++
		}

		if singleResult = find(startLeft-1, startRight); len(singleResult) > len(result) {
			result = singleResult
		}
		startLeft = startRight
	}
	return result
}

func main() {
	s := "adddddaadadasdadadadadad"
	fmt.Println(longestPalindrome(s))
}
