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

func RemoveIndex(s []int, index int) []int {
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

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func MinTriple(a, b, c int) int {
	return Min(Min(a, b), Min(a, c))
}

func TransVerticalBoard(board [][]byte) [][]byte {
	verticalBoard := [][]byte{}
	for c := range board[0] {
		temp := []byte{}
		for r := range board {
			// fmt.Printf("[%d,%d] %5v  ", c, r, board[r][c])
			temp = append(temp, board[r][c])
		}
		verticalBoard = append(verticalBoard, temp)
	}
	// fmt.Printf("origin:%v\n", board)
	// fmt.Printf("trans:%v\n", verticalBoard)
	return verticalBoard
}

func GenEmptyZeroSizeGrid(boundaryWidth, boundaryHigh int) [][]int {
	sets := make([][]int, boundaryHigh)
	for i := 0; i < boundaryHigh; i++ {
		temp := []int{}
		for j := 0; j < boundaryWidth; j++ {
			temp = append(temp, 0)
		}
		sets[i] = temp
	}
	return sets
}

func LearnBits() {
	// i >>= 1
	// <==>
	// tmp := i >> 1
	// i = tmp
	// Bitwise Operators
	// https://michaelchen.tech/golang-programming/operator/
	// https://stackoverflow.com/questions/32933333/what-does-the-operator-do-in-golang
	var res uint32 = 0
	var num uint32 = 10
	lens := 4
	formatStr := fmt.Sprintf("%v%v%v", "%0", lens, "b")
	resStr := fmt.Sprintf(formatStr, res)
	numStr := fmt.Sprintf(formatStr, num)
	fmt.Printf("=====in for res:%v ,num:%v=====\n", res, num)
	for i := 0; i < lens; i++ {
		resStr = fmt.Sprintf(formatStr, res)
		numStr = fmt.Sprintf(formatStr, num)
		fmt.Printf("ith:%v before res:%b, num in base 2:%b\n", i, res, num)
		fmt.Printf("ith:%v before res:%v, num in use fmt base 2:%v\n", i, resStr, numStr)
		res = (res << 1) ^ (num & 1)
		num >>= 1
		resStr = fmt.Sprintf(formatStr, res)
		numStr = fmt.Sprintf(formatStr, num)
		fmt.Printf("ith:%v after res:%b, num in base 2:%b\n", i, res, num)
		fmt.Printf("ith:%v after res:%v, num in use fmt base 2:%v\n", i, resStr, numStr)
	}
	resStr = fmt.Sprintf(formatStr, res)
	numStr = fmt.Sprintf(formatStr, num)
	fmt.Printf("=====out for res:%v ,num:%v=====\n", res, num)
}
