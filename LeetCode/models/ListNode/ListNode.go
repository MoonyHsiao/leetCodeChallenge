package listNode

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

type LineList struct {
	first *ListNode
}

// Append LineList
func (l *LineList) Append(x int) {
	s := &ListNode{
		Val: x,
	}
	if l.first == nil {
		l.first = s
	} else {
		currentNode := l.first
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = s
	}
}

// ListLength LineList
func ListLength(l *LineList) int {
	len := 0
	listHead := l.first
	if listHead != nil {
		len++
		for listHead.Next != nil {
			listHead = listHead.Next
			len++
		}
	}
	return len
}

// GetElem LineList
func (l *LineList) GetElem(index int) *ListNode {
	res := &ListNode{}
	listHead := l.first
	if index <= ListLength(l) {
		for i := 0; i < index; i++ {
			res = listHead
			listHead = listHead.Next
		}
	}
	return res
}

// Display LineList
func (l *LineList) Display() {
	listHead := l.first
	for listHead != nil {
		fmt.Printf("%+v ->", listHead.Val)
		listHead = listHead.Next
	}
	fmt.Println()
}

// Insert LineList
func (l *LineList) Insert(index, value int) {
	s := &ListNode{
		Val: value,
	}

	len := ListLength(l)

	if index > len || index < 0 {
		return
	}

	if l.first == nil {
		l.first = s
	} else {
		currentNode := l.first
		for i := 0; i < index-1; i++ {
			if currentNode.Next != nil {
				currentNode = currentNode.Next
			}
		}
		// temp := currentNode.Next
		// currentNode.Next = s
		// s.Next = temp
		//跟上面等價 看喜歡用哪個 現在只是用不習慣
		s.Next = currentNode.Next
		currentNode.Next = s

	}
}

// Remove LineList
func (l *LineList) Remove(index int) {
	len := ListLength(l)

	if l.first == nil || index > len || index < 0 {
		return
	}
	if index == 0 {
		currentNode := l.first
		if currentNode.Next != nil {
			currentNode = currentNode.Next
			l.first = currentNode
		}

	} else {
		currentNode := l.first
		prevNode := l.first
		for i := 0; i < index-1; i++ {
			if currentNode.Next != nil {
				prevNode = currentNode
				currentNode = currentNode.Next
			}
		}
		temp := currentNode.Next
		prevNode.Next = temp
	}
}

//InsertAfter 這跟Append沒啥兩樣吧
func (l *LineList) InsertAfter(value int) {

	// s := &ListNode{
	// 	Val: value,
	// }
	// if l.head == nil {
	// 	l.head = s
	// 	return
	// }
	// currentNode := l.head
	// for currentNode.Next != nil {
	// 	currentNode = currentNode.Next
	// }
	// currentNode.Next = s
	l.Append(value)
}

//InsertFront aaa
func (l *LineList) InsertFront(value int) {
	s := &ListNode{
		Val: value,
	}
	if l.first == nil {
		l.first = s
		return
	}
	listHead := l.first
	l.first = s
	s.Next = listHead
}

//Clear aaa
func (l *LineList) Clear() {
	listHead := l.first
	for listHead != nil {
		curr := listHead
		listHead = listHead.Next
		curr.Next = nil
	}
	l.first = nil
}
