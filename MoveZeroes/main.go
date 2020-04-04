package main

import "fmt"

func main() {
	data := []int{2, 0, 5, 7, 0, 13}
	moveZeroes(data[:])
	fmt.Println(data)
}

func moveZeroes(nums []int) {
	// nums[1] = 100
	for i := range nums {
		nums[i] = nums[i] * 100
	}
}
