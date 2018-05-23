package main

import (
	"fmt"
)

func main() {
	strs:=[]string{}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs)<2{
		if len(strs)<1{
			return ""
		}
		return strs[0]
	}

	for k,v:=range strs{

	}
}
