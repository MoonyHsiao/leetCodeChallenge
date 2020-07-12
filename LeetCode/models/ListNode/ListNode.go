package ListNode

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateTestData(data string) *ListNode {
	if data == "[]" {
		return nil
	}
	data = string([]rune(data)[1 : len(data)-1])
	res := strings.Split(data, ",")
	length := len(res)
	listNode := make([]ListNode, length)
	headVal, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	listNode[0] = ListNode{headVal, nil}
	for i := 1; i < length; i++ {
		headVal, _ = strconv.Atoi(res[i])
		listNode[i] = ListNode{headVal, nil}
		listNode[i-1].Next = &listNode[i]
	}
	return &listNode[0]
}

func Print(listNode *ListNode) {
	if listNode == nil {
		fmt.Println(nil)
	}
	var buffer strings.Builder
	buffer.WriteString("[")
	value := strconv.Itoa(listNode.Val)
	buffer.WriteString(value)
	temp := listNode.Next
	for temp != nil {
		buffer.WriteString(",")
		value = strconv.Itoa(temp.Val)
		buffer.WriteString(value)
		temp = temp.Next
	}
	buffer.WriteString("]")
	fmt.Println(buffer.String())
}

//這邊還要改一下 不過暫時堪用了
func GenData(arr []int) *ListNode {
	temp := &ListNode{}
	head := temp
	for i := range arr {
		s := &ListNode{
			Val: arr[i],
		}
		if temp == nil {
			temp = s
		} else {
			temp.Next = s
		}
		temp = temp.Next
	}
	return head.Next
}
