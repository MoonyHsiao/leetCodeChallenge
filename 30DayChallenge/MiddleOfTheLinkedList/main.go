package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}

	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }

	for i := range arr {
		addTestData(arr[i])
	}

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

func addTestData(x int) {
	s := &ListNode{
		Val: x,
	}
	if head == nil {
		head = s
	} else {
		currentNode := head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = s
	}
}
