package solveed

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lexhsiao135/ds-go/LeetCode/models/ListNode"
	"github.com/lexhsiao135/ds-go/LeetCode/models/TreeNode"
	"github.com/lexhsiao135/ds-go/leetcode/models"
)

func mergeTwoLists(l1 *models.ListNode, l2 *models.ListNode) *models.ListNode {

	arr := []int{}
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				arr = append(arr, l1.Val)
				l1 = l1.Next
			} else if l2.Val < l1.Val {
				arr = append(arr, l2.Val)
				l2 = l2.Next
			}
		} else if l1 == nil && l2 != nil {
			for l2 != nil {
				arr = append(arr, l2.Val)
				l2 = l2.Next
			}
		} else if l1 != nil && l2 == nil {
			for l1 != nil {
				arr = append(arr, l1.Val)
				l1 = l1.Next
			}
		}
	}
	res := models.GenData(arr)
	return res
}

func mergeTwoListsV3(l1 *models.ListNode, l2 *models.ListNode) *models.ListNode {
	temp := &models.ListNode{}
	head := temp

	for l1 != nil && l2 != nil {

		if l1.Val <= l2.Val {
			s := &models.ListNode{
				Val: l1.Val,
			}
			temp.Next = s
			temp = temp.Next
			l1 = l1.Next
		} else if l2.Val < l1.Val {
			s := &models.ListNode{
				Val: l2.Val,
			}
			temp.Next = s
			temp = temp.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		temp.Next = l2
	} else if l2 == nil {
		temp.Next = l1
	}
	return head.Next
}

// 用內部的排序來解

func LargestNumber(nums []int) string {
	s := models.NewIntSlice(nums)
	sort.Sort(s)
	res := ""
	for _, v := range s {
		res += strconv.Itoa(v)
		if res == "0" {
			res = ""
		}
	}
	if res == "" {
		return "0"
	}
	return res
}

func Find132pattern(nums []int) bool {

	for len(nums) > 0 {
		s3 := nums[len(nums)-1]
		s2 := 0
		isFindS2 := false
		isFindS1 := false
		index := 0
		for i := len(nums) - 1; i >= 0; i-- {
			temp := nums[i]
			if temp > s3 {
				s2 = temp
				isFindS2 = true
				index = i
				break
			}
		}

		for i := index; i >= 0; i-- {
			temp := nums[i]
			if temp < s2 && s3 > temp {
				isFindS1 = true
				break
			}
		}
		if isFindS1 && isFindS2 {
			return isFindS1 && isFindS2
		}
		nums = models.RemoveIndex(nums, len(nums)-1)
	}
	return false

}

func IncreasingTriplet(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		m1 := nums[i]
		m2 := models.MaxInt
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > m1 && m2 > nums[j] {
				m2 = nums[j]
			}
			if nums[j] > m2 && m2 > m1 {
				return true
			}
		}
	}
	return false
}

func RemoveDuplicates(nums []int) int {
	count := 0
	temp := models.MaxInt
	ans := []int{}
	if len(nums) <= 0 {
		return count
	}
	for i := range nums {
		if temp == nums[i] {
			continue
		}
		temp = nums[i]
		count += 1
		ans = append(ans, temp)
	}
	fmt.Println(ans[:])
	return count
}

func RemoveDuplicatesV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[:j]) //nums[:j]为新的去重后的数组
	return j
}

func HasCycle(head *models.ListNode) bool {
	m := make(map[*models.ListNode]bool)
	if head == nil {
		return false
	}
	for head.Next != nil {
		_, ok := m[head]
		if ok {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return false
}

func SearchInsert(nums []int, target int) int {

	if target > nums[len(nums)-1] {
		return len(nums)
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			index = i
			break
		}
	}
	return index
}

func LengthOfLastWord(s string) int {
	count := 0
	max := 0
	for i := range s {
		if s[i] != 32 {
			count += 1
			max = count
		}
		if s[i] == 32 {
			count = 0
		}
	}
	return max
}

func StrStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}
	nlen := len(needle)
	for i := 0; i <= len(haystack)-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

func IsPalindrome(s string) bool {
	pat := "[^0-9A-Za-z]"
	re, _ := regexp.Compile(pat)

	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	if len(s) <= 1 {
		return true
	}
	frontIndex := 0
	backIndex := len(s) - 1
	front := s[frontIndex]
	back := s[backIndex]

	for frontIndex < backIndex {
		if front != back {
			return false
		}
		frontIndex++
		backIndex--
		front = s[frontIndex]
		back = s[backIndex]
	}
	return true
}

func MySqrt(x int) int {
	if x == 0 {
		return x
	}
	res := models.GetIntSize(x)
	if res%2 == 0 {
		res = res / 2
	} else {
		res = res/2 + 1
	}
	low := 1
	high := 9
	for i := 1; i < res; i++ {
		low = low * 10
		high = high*10 + 9
	}

	for low < high-1 {
		mid := (low + high) / 2
		if mid*mid < x {
			low = mid
		} else if mid*mid > x {
			high = mid
		} else {
			return mid
		}
	}
	return low
}

func PlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func AddBinary(a string, b string) string {

	count := 0
	if len(a) >= len(b) {
		count = len(a)
	} else {
		count = len(b)
	}

	formatStr := fmt.Sprintf("%v%v%v", "%0", count, "s")
	fmta := fmt.Sprintf(formatStr, a)
	fmtb := fmt.Sprintf(formatStr, b)

	a_arr := make([]int, count)
	b_arr := make([]int, count)
	res_arr := make([]int, count)

	for i := 0; i < count; i++ {
		a, _ := strconv.Atoi(string(byte(fmta[count-1-i])))
		a_arr[i] = a
		b, _ := strconv.Atoi(string(byte(fmtb[count-1-i])))
		b_arr[i] = b
	}

	a_arr = models.ReverseInts(a_arr)

	b_arr = models.ReverseInts(b_arr)

	isCarry := 0

	for i := count - 1; i >= 0; i-- {
		xor := a_arr[i] ^ b_arr[i] ^ isCarry
		res_arr[i] = xor
		// fmt.Printf("step%v,a_arr[i]:%v,b_arr[i]:%v,isCarry:%v,xor:%v,res_arr:%v\n", i, a_arr[i], b_arr[i], isCarry, xor, res_arr[i])
		if (a_arr[i] + b_arr[i] + isCarry) > 1 {
			isCarry = 1
		} else {
			isCarry = 0
		}

	}

	if isCarry == 1 {
		res_arr = append([]int{1}, res_arr...)
	}

	res := models.ArrayToString(res_arr, "")
	return res

}

func DeleteDuplicates(head *ListNode.ListNode) *ListNode.ListNode {
	if head == nil {
		return head
	}
	first := head
	for head.Next != nil {
		if head.Val != head.Next.Val {
			head = head.Next

		} else {
			head.Next = head.Next.Next
		}
	}
	return first
}

//用PreOrderTraverse來解
func IsSameTree(p *TreeNode.TreeNode, q *TreeNode.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func MaxDepth(root *TreeNode.TreeNode) int {
	return countDepth(root, 1)
}

func countDepth(root *TreeNode.TreeNode, count int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return countDepth(root.Right, count+1)
	}

	if root.Right == nil && root.Left != nil {
		return countDepth(root.Left, count+1)
	}

	if root.Right != nil && root.Left != nil {
		return models.Max(countDepth(root.Left, count+1), countDepth(root.Right, count+1))
	}
	return count
}
