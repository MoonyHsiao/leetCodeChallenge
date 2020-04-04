package main

import "fmt"

func main() {
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(data)
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	max := 0
	currentMax := 0
	isAllNega := isAllNega(nums)
	if isAllNega {
		max = nums[0]
		for i := range nums {
			if max < nums[i] {
				max = nums[i]
			}
		}

	} else {
		for i := range nums {
			currentMax += nums[i]
			if currentMax < 0 {
				currentMax = 0
			}
			if max < currentMax {
				max = currentMax
			}
		}
	}
	return max
}

func isAllNega(nums []int) bool {
	isNega := true
	for i := range nums {
		if nums[i] > 0 {
			isNega = false
			break
		}
	}
	return isNega
}
