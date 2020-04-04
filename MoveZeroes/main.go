package main

import "fmt"

func main() {
	// data := []int{2, 0, 5, 7, 0, 13}
	data := []int{0, 1, 0, 3, 12}
	moveZeroes(data[:])
	fmt.Println(data)
}

func moveZeroes(nums []int) {
	for i := range nums {
		if nums[i] == 0 {
			slice := nums[i:]
			for j := range slice {
				if slice[j] != 0 {
					value := slice[j]
					nums[i] = value
					slice[j] = 0
					break
				}
			}
		}
	}
}
