package models

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
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

func RemoveIndex(s []int, index int) []int {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetFirstNum(num int) int {
	res := num
	for num%10 > 0 {
		res = num % 10
		num = num / 10
	}
	return res
}

type intSlice []int

func NewIntSlice(a []int) intSlice {
	b := intSlice{}
	for _, v := range a {
		b = append(b, v)
	}
	return b
}
func (s intSlice) Len() int {
	return len(s)
}

// 就是a，b兩個數，比較函數直接返回 a+b<b+a的值
func (s intSlice) Less(i, j int) bool {
	stri := strconv.Itoa(s[i])
	strj := strconv.Itoa(s[j])
	s1, _ := strconv.Atoi(stri + strj)
	s2, _ := strconv.Atoi(strj + stri)
	return s1 > s2
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
