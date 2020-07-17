package solveed

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

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
