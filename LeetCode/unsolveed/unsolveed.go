package unsolveed

import (
	"strconv"
)

// https://leetcode.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	values []int
	head   int
	tail   int
	count  int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	res := MyQueue{}
	return res
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.values = append(this.values, x)
}

/** Removes the element from in front of queue and returns that element. */
// func (this *MyQueue) Pop() int {

// }

/** Get the front element. */
// func (this *MyQueue) Peek() int {

// }

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.values) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// 9 的倍數的判斷式啊
// 原數 mod 9 = 新數 mod 9, 所以必定只差 9 的倍數
// 給定任意正整數，找出位數總和與它相同，而且大於它的最小正整數, 例如 772 => 781 ; 293 => 329;  100 => 1000
// (abcde)
// = a*10000+b*1000+c*100+d*10+e
// = a*(9999+1)+b*(999+1)+c*(99+1)+d*(9+1)+e
// = (9999a+999b+99c+9d)+(a+b+c+d+e)
// = 9*(1111a+111b+11c+d)+(a+b+c+d+e)
func MinimumGreaterThan(n int) int {
	return n
}

func MinimumGreaterThanGenAnswer(n int) int {
	res := -1
	base := checkint(n)

	isEnd := false
	for !isEnd {
		n += 9
		isEnd = (base == checkint(n))
	}
	res = n
	return res
}

func checkint(n int) int {
	s := strconv.Itoa(n)
	nums := trans(s)
	return sum(nums)
}

func sum(nums []int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}

func trans(n string) []int {
	nums := []int{}
	for i := range n {
		temp, _ := strconv.Atoi(string(n[i]))
		nums = append(nums, temp)
	}
	return nums
}
