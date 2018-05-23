package main

import "fmt"

func main() {
	//a := &ListNode{7, nil}
	//b := &ListNode{5, nil}
	//c := &ListNode{6, a}
	//b := &ListNode{5, nil}
	//c := &ListNode{5, nil}
	b := &ListNode{1, nil}
	a := &ListNode{9, nil}
	c := &ListNode{9, a}
	fmt.Println(addTwoNumbers(b, c), addTwoNumbers(b, c).Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	result.Val = l1.Val + l2.Val

	if l1.Next == nil && l2.Next == nil {
		if result.Val > 9 {
			result.Next = &ListNode{1, nil}
			result.Val %= 10
		}
		return result
	} else if l1.Next == nil {
		if result.Val > 9 {
			l1.Next = &ListNode{1, nil}
			result.Val %= 10
		} else {
			l1.Next = &ListNode{0, nil}
		}
	} else if l2.Next == nil {
		if result.Val > 9 {
			l2.Next = &ListNode{1, nil}
			result.Val %= 10
		} else {
			l2.Next = &ListNode{0, nil}
		}
	} else {
		if result.Val > 9 {
			l1.Next.Val += 1
			result.Val %= 10
		}
	}

	result.Next = addTwoNumbers(l1.Next, l2.Next)

	return result
}
