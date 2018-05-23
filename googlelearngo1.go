package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch {
	case op == "+":
		return a + b, nil
	case op == "-":
		return a - b, nil
	case op == "*":
		return a * b, nil
	case op == "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with arguments: [%d,%d]\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func sumArgs(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	if result, err := eval(2, 3, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	//fmt.Println(eval(2, 3, "+"))
	fmt.Println(div(13, 3))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sumArgs(1, 2, 3, 555, 3))
}
