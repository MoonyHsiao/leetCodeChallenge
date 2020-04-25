package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// arr := []int{1, 2, 3, 4, 5}

	head := new(ListNode)
	head.Val = 1
	temp := new(ListNode)
	temp.Val = 2
	head.Next = temp
	temp2 := new(ListNode)
	temp2.Val = 3
	temp.Next = temp2

	res := middleNode(head)
	fmt.Printf("res:%v", res)

}

func middleNode(head *ListNode) *ListNode {
	count := 1
	res := head
	for {
		if head.Next == nil {
			break
		} else {
			count += 1
			head = head.Next
		}
	}
	upper := count / 2

	for i := 0; i < upper; i++ {
		res = res.Next
	}

	return res
}
