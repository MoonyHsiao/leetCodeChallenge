package models

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	s []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MinStack) Pop() {
	index := len(this.s) - 1
	this.s = RemoveIndexV2(this.s, index)
}

func (this *MinStack) Top() int {
	index := len(this.s) - 1
	return this.s[index]
}

func (this *MinStack) GetMin() int {
	min := this.s[0]
	for i := 0; i < len(this.s); i++ {
		if this.s[i] < min {
			min = this.s[i]
		}
	}
	return min
}
