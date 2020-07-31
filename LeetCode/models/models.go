package models

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MinUint uint = 0

	MaxUint = ^MinUint

	MaxInt = int(MaxUint >> 1)

	MinInt = ^MaxInt
)

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

// RemoveIndex(nums, 0) 會有bugS
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RemoveIndexV2(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func GetIntSize(num int) int {
	count := 1
	for num/10 > 0 {
		count++
		num = num / 10
	}
	return count
}

func ReverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(ReverseInts(input[1:]), input[0])
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
