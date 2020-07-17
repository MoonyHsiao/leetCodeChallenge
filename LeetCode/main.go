package main

import (
	"fmt"

	"github.com/lexhsiao135/ds-go/LeetCode/solveed"
)

func main() {
	// data := []int{1, 3, 5, 6}
	// target := 5
	data := "Marge, let's \"[went].\" I await {news} telegram."
	res := solveed.IsPalindrome(data)
	fmt.Printf("res:%v\n", res)

}
