package main

import "fmt"

func main() {
	data := []int{2, 2, 1}
	res := singleNumber(data)
	fmt.Println(res)
}

func singleNumber(nums []int) int {

	numsMap := make(map[int]int)
	for _, e := range nums {
		_, isExist := numsMap[e]
		if isExist {
			numsMap[e] += 1
		} else {
			numsMap[e] = 1
		}
	}
	result := 0
	for k, v := range numsMap {
		if v == 1 {
			result = k
			break
		}
	}
	return result
}
