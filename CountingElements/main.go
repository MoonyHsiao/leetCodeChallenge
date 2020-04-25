package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 1, 2}
	res := countElements(arr)
	fmt.Printf("res:%v", res)
}

func countElements(arr []int) int {
	res := 0
	m := make(map[int]int)
	keys := make([]int, 0, len(m))
	for i := range arr {
		_, ok := m[arr[i]]
		if ok {
			m[arr[i]] += 1
		} else {
			temp := arr[i]
			m[temp] = 1
			keys = append(keys, temp)
		}
	}

	sort.Ints(keys)
	for i := 0; i < len(keys)-1; i++ {
		if keys[i]+1 == keys[i+1] {
			res += m[keys[i]]
		}
	}

	return res
}
