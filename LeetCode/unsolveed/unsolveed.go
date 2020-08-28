package unsolveed

import (
	"fmt"
	"strconv"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
)

// https://leetcode.com/problems/minimum-cost-for-tickets/
func MincostTickets(days []int, costs []int) int {
	if len(days) == 1 {
		return models.MinTriple(costs[0], costs[1], costs[2])
	}
	temp := days[0]
	a, b, c := costs[0], costs[1], costs[2]
	index7Day, index30Day := 0, 0
	is7DayPassEnd, is30DayPassEnd := true, true
	for i := range days {
		index7Day = i
		if temp+7 <= days[i] {
			is7DayPassEnd = false
			break
		}
	}

	for i := range days {
		index30Day = i
		if temp+30 <= days[i] {
			is30DayPassEnd = false
			break
		}
	}
	reduce1DayPassArray, reduce7DayPassArray, reduce30DayPassArray := days[1:], days[index7Day:], days[index30Day:]
	a += MincostTickets(reduce1DayPassArray, costs)
	if !is7DayPassEnd {
		b += MincostTickets(reduce7DayPassArray, costs)
	}
	if !is30DayPassEnd {
		c += MincostTickets(reduce30DayPassArray, costs)
	}
	return models.MinTriple(a, b, c)
}

func MincostTicketsTimeOutVer(days []int, costs []int) int {
	if len(days) == 1 {
		return models.MinTriple(costs[0], costs[1], costs[2])
	}
	temp := days[0]
	a, b, c := costs[0], costs[1], costs[2]
	index7Day, index30Day := 0, 0
	is7DayPassEnd, is30DayPassEnd := true, true
	for i := range days {
		index7Day = i
		if temp+7 <= days[i] {
			is7DayPassEnd = false
			break
		}
	}

	for i := range days {
		index30Day = i
		if temp+30 <= days[i] {
			is30DayPassEnd = false
			break
		}
	}
	reduce1DayPassArray, reduce7DayPassArray, reduce30DayPassArray := days[1:], days[index7Day:], days[index30Day:]
	a += MincostTickets(reduce1DayPassArray, costs)
	if !is7DayPassEnd {
		b += MincostTickets(reduce7DayPassArray, costs)
	}
	if !is30DayPassEnd {
		c += MincostTickets(reduce30DayPassArray, costs)
	}
	return models.MinTriple(a, b, c)
}

func CountIncreaseSubArray(nums []int) int {
	m := make([][]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		for j := 0; j <= i; j++ {
			if nums[j] < temp {
				m[i] = append(m[i], nums[j])
			}
			// else {
			// 	m[i] = append(m[i], -1)
			// }
		}
		// if len(m[i]) == 0 {
		// 	m[i] = append(m[i], temp)
		// }
	}
	// for i := range m {
	// 	res += len(m[i])
	// }
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("m:%v\n", m)
	return res
	// Input : arr[] = {3, 2, 4, 5, 4}
	// Output : 14
	// Sub-sequences are {3}, {2}, {4},
	// count[3]=0
	// count[2]=0
	// count[4]=1+count[2]+1+count[3]=2
	// {3,4},
	// {2,4},
	// count[5]=1+count[4]+1+count[3]+1+count[2]=1+2+1+1=5
	// {3,5},
	// {2,5},
	// {4,5},
	// {2,4,5}
	// {3,4,5},
	// count[4]=1+count[2]+1+count[3]
	// {3,4},
	// {2,4}
}

// https://blog.csdn.net/weixin_40546602/article/details/88546583

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

func MinimumGreaterThanV2(data int) int {
	r := data % 100
	q := data / 100
	step := 0
	if r == 0 {
		b0 := 100
		for q%10 == 0 {
			b0 *= 10
			q /= 10
		}
		s0 := q % 10
		step = (b0-1)*(10-s0)/9 + 1
		q /= 10
		for q%10 == 9 {
			step += s0
			s0 *= 10
			q /= 10
		}
	} else if r > 90 {
		b9X := 10
		s9X := r % 10
		for q%10 == 9 {
			b9X *= 10
			q /= 10
		}
		step = (b9X-1)*s9X/9 + 1
	} else if r%10 == 0 {
		sX0 := r / 10
		step = 11 - sX0

		for q%10 == 9 {
			step += sX0
			sX0 *= 10
			q /= 10
		}
	} else {
		step = 1
	}

	return data + step*9
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
