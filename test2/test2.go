package main

import "fmt"

func main() {
	s:="PAYPALISHIRING"
	fmt.Printf(convert(s,4))
}

func convert(s string, numRows int) string {
	var col int
	switch  {
	case numRows<=0:
		return ""
	case numRows==1:
		col=13
	case numRows==2||numRows==3||numRows==4:
		col=7
	case numRows==5||numRows==6:
		col=6
	case numRows>6:
		col=15-numRows
	}

	arr:=make([][]string,14)
	for k,v:=range s{
		arr[k]=make([]string,numRows)
		for k<=numRows{
		}
	}
}
