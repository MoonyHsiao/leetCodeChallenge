package main

import (
	"fmt"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/solveed"
)

func main() {
	// data := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	// target := 5
	// data := 100
	a := "1010"

	b := "1011"
	a = "0"
	b = "1"
	a = "1111"
	b = "1111"

	res := solveed.AddBinary(a, b)
	fmt.Printf("res:%v\n", res)

}
