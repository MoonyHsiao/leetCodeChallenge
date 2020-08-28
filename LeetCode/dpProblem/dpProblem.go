package dpProblem

import (
	"fmt"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
)

// https://www.geeksforgeeks.org/dynamic-programming/

// ============Ugly Numbers============
func maxDivide(a, b int) int {
	for a%b == 0 {
		a = a / b
	}
	return a
}

func isUgly(no int) bool {
	no = maxDivide(no, 2)
	no = maxDivide(no, 3)
	no = maxDivide(no, 5)
	return (no == 1)
}

func GetNthUglyNo(n int) int {
	count := 1
	ith := 1
	for n > ith {
		count++
		if isUgly(count) {
			ith++
		}
	}
	return count
}

// (1) 1×2, 2×2, 3×2, 4×2, 5×2, …
// (2) 1×3, 2×3, 3×3, 4×3, 5×3, …
// (3) 1×5, 2×5, 3×5, 4×5, 5×5, …

func GetNthUglyNoDP(n int) uint16 {
	ugly := []uint16{}
	var i2, i3, i5 uint16 = 0, 0, 0
	var nextMultipleOf2, nextMultipleOf3, nextMultipleOf5 uint16 = 2, 3, 5
	var nextUglyNo uint16 = 1
	ugly = append(ugly, 1)

	for i := 1; i < n; i++ {
		nextUglyNo = models.MinUint16(nextMultipleOf2, models.MinUint16(nextMultipleOf3, nextMultipleOf5))
		ugly = append(ugly, nextUglyNo)
		if nextUglyNo == nextMultipleOf2 {
			i2++
			nextMultipleOf2 = ugly[i2] * 2
		}
		if nextUglyNo == nextMultipleOf3 {
			i3++
			nextMultipleOf3 = ugly[i3] * 3
		}
		if nextUglyNo == nextMultipleOf5 {
			i5++
			nextMultipleOf5 = ugly[i5] * 5
		}

	}
	return nextUglyNo
}

// ============nth Catalan Number============

func Catalan(n int) uint32 {
	if n <= 1 {
		return 1
	}
	var res uint32 = 0
	for i := 0; i < n; i++ {
		res += Catalan(i) * Catalan(n-i-1)
	}
	return res
}

func Binomial(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return Binomial(n-1, k-1) + Binomial(n-1, k)
}

var dp [100][3]int

func RGCBlock_ButtonUP(n, m int) int {

	if n == 1 {
		return 1
	}
	if dp[n][m] != 0 {
		return dp[n][m]
	}
	// red
	if m == 0 {
		dp[n][m] = RGCBlock_ButtonUP(n-1, 0) + RGCBlock_ButtonUP(n-1, 1) + RGCBlock_ButtonUP(n-1, 2)
	}
	//green
	if m == 1 {
		dp[n][m] = RGCBlock_ButtonUP(n-1, 0) + RGCBlock_ButtonUP(n-1, 1)
	}
	//blue
	if m == 2 {
		dp[n][m] = RGCBlock_ButtonUP(n-1, 0) + RGCBlock_ButtonUP(n-1, 2)
	}
	return dp[n][m]
}

func RGCBlock_TopDown(n, m int) int {

	prevRed := 1
	prevGreen := 1
	prevBlue := 1
	for i := 1; i < n; i++ {
		tempRed := prevRed
		tempGreen := prevGreen
		tempBlue := prevBlue
		prevRed = tempRed + tempGreen + tempBlue
		prevGreen = tempRed + tempBlue
		prevBlue = tempRed + tempGreen
	}
	res := 0
	if m == 0 {
		res = prevRed
	} else if m == 1 {
		res = prevGreen
	} else if m == 2 {
		res = prevBlue
	}
	return res
}

func Hanoi(n int, a, b, c string) {
	if n == 1 {
		res := fmt.Sprintf("%v->%v", a, c)
		fmt.Printf("move:%v\n", res)
	} else {

		Hanoi(n-1, a, c, b)
		Hanoi(1, a, b, c)
		Hanoi(n-1, b, a, c)
	}
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
