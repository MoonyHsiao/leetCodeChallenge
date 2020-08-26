package unsolveed

import (
	"fmt"
	"strconv"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
)

// res:[1 10 100 101 102 103 104 105 106 107 108 109 11 110 111 112 113 114 115 116 117 118 119 12 120 121 122 123 124 125 126 127 128 129 13 130 131 132 133 134 135 136 137 138 139 14 140 141 142 143 144 145 146 147 148 149 15 150 151 152 153 154 155 156 157 158 159 16 160 161 162 163 164 165 166 167 168 169 17 170 171 172 173 174 175 176 177 178 179 18 180 181 182 183 184 185 186 187 188 189 19 190 191 192 193 194 195 196 197 198 199 2 20 200 21 22 23 24 25 26 27 28 29 3 30 31 32 33 34 35 36 37 38 39 4 40 41 42 43 44 45 46 47 48 49 5 50 51 52 53 54 55 56 57 58 59 6 60 61 62 63 64 65 66 67 68 69 7 70 71 72 73 74 75 76 77 78 79 8 80 81 82 83 84 85 86 87 88 89 9 90 91 92 93 94 95 96 97 98 99]
// res:[1 10 100 101 102 103 104 105 106 107 108 109 11 110 111 112 113 114 115 116 117 118 119 12 120 121 122 123 124 125 126 127 128 129 13 130 131 132 133 134 135 136 137 138 139 14 140 141 142 143 144 145 146 147 148 149 15 150 151 152 153 154 155 156 157 158 159 16 160 161 162 163 164 165 166 167 168 169 17 170 171 172 173 174 175 176 177 178 179 18 180 181 182 183 184 185 186 187 188 189 19 190 191 192 193 194 195 196 197 198 199 2 20 200 201 202 203 204 205 206 207 208 209 21 22 23 24 25 26 27 28 29 3 30 31 32 33 34 35 36 37 38 39 4 40 41 42 43 44 45 46 47 48 49 5 50 51 52 53 54 55 56 57 58 59 6 60 61 62 63 64 65 66 67 68 69 7 70 71 72 73 74 75 76 77 78 79 8 80 81 82 83 84 85 86 87 88 89 9 90]
func LexicalOrder(n int) []int {
	res := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			// 要++
			cur++
			for cur%10 == 0 {
				cur = cur / 10
			}
		}

	}
	return res

}
func LexicalOrderAns(n int) []int {
	res := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur >= n {
				cur = cur / 10
			}
			cur++
			for cur%10 == 0 {
				cur = cur / 10
			}

		}
	}
	return res

}

// https://blog.csdn.net/weixin_40546602/article/details/88546583
// https://leetcode.com/problems/number-of-islands/

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
