package unsolveed

import (
	"fmt"
	"strconv"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
)

// https://leetcode.com/problems/magnetic-force-between-two-balls/
// 用binary search
func MaxDistance(position []int, m int) int {
	res := 0
	sorted := sort.MergeSort(position)
	low := 0
	high := len(position) - 1
	max := position[high] - position[low]
	res = max
	fmt.Printf("sorted:%v\n", sorted)
	// 找到最大最小之後 下一個值只能是最大的1/2以下
	for m > 0 {

		m--
	}
	return res
}

func isPossible() bool {
	return false
}

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

func test() {
	res := 0
	fmt.Print(res)
	fmt.Printf("res:%v\n", res)
}
