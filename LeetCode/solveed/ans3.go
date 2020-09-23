package solveed

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/listNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
)

func IsPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	min := 1
	max := num / 2
	for min <= max {
		mid := (min + max) / 2
		if mid*mid < num {
			min = mid + 1
		} else if mid*mid > num {
			max = mid - 1
		} else {
			return true
		}
	}
	return false
}

// greed
func FindContentChildren(g []int, s []int) int {
	child := sort.MergeSort(g)
	cookies := sort.MergeSort(s)
	count := 0
	i := 0
	j := 0
	for i < len(child) && j < len(cookies) {
		if child[i] <= cookies[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}
	return count
}

// greed
func LemonadeChange(bills []int) bool {
	fiveDollarCount := 0
	tenDollarCount := 0
	twentyDollarCount := 0
	for i := range bills {
		if bills[i] == 5 {
			fiveDollarCount++
			continue
		} else if bills[i] == 10 {
			tenDollarCount++
			if fiveDollarCount >= 1 {
				fiveDollarCount--
			} else {
				return false
			}
		} else if bills[i] == 20 {
			twentyDollarCount++
			if fiveDollarCount >= 1 && tenDollarCount >= 1 {
				fiveDollarCount--
				tenDollarCount--
			} else if fiveDollarCount >= 3 {
				fiveDollarCount -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func ReverseBits(num uint32) uint32 {
	var res uint32 = 0

	for i := 0; i < 32; i++ {
		res = (res << 1) ^ (num & 1)
		num >>= 1
	}
	return res
}

func ValidIPAddress(IP string) string {
	res := "Neither"
	dot := strings.Count(IP, ".")
	dash := strings.Count(IP, ":")
	isIPV4 := dot > 0
	isIPV6 := dash > 0
	if isIPV4 {
		sets := strings.Split(IP, ".")
		isValidIP := true
		if len(sets) == 4 {
			for i := range sets {
				isValidIP = ipv4(sets[i]) && isValidIP
			}
		} else {
			isValidIP = false
		}
		if isValidIP {
			res = "IPv4"
		}
	}
	if isIPV6 {
		sets := strings.Split(IP, ":")
		isValidIP := true
		if len(sets) == 8 {
			for i := range sets {
				isValidIP = ipv6(sets[i]) && isValidIP
			}
		} else {
			isValidIP = false
		}
		if isValidIP {
			res = "IPv6"
		}
	}
	return res
}

func ipv4(s string) bool {
	if len(s) == 1 {
		return true
	} else {
		temp, err := strconv.Atoi(s)
		if temp > 255 || err != nil {
			return false
		}
		if s[0] == '0' {
			return false
		}
	}
	return true
}

func ipv6(s string) bool {
	if len(s) > 4 || len(s) < 1 {
		return false
	} else {
		for i := 0; i < len(s); i++ {
			_, err := strconv.ParseUint(string(s[i]), 16, 64)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func IsPalindromeLinkList(head *listNode.ListNode) bool {
	sets := []int{}
	for head != nil {
		sets = append(sets, head.Val)
		head = head.Next
	}
	if len(sets) == 0 {
		return true
	}
	frontIndex := 0
	backIndex := len(sets) - 1
	front := sets[frontIndex]
	back := sets[backIndex]
	for frontIndex < backIndex {
		if front != back {
			return false
		}
		frontIndex++
		backIndex--
		front = sets[frontIndex]
		back = sets[backIndex]
	}
	return true
}

func MaxArea(height []int) int {
	max := 0
	size := len(height)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			temp := models.Min(height[i], height[j]) * (j - i)
			max = models.Max(max, temp)
		}
	}
	return max
}

func GenerateParenthesis(n int) []string {
	res := []string{}
	generateParenthesisBacktrack(&res, "", 0, 0, n)
	return res
}

func generateParenthesisBacktrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		(*ans) = append((*ans), cur)
		return
	}
	if open < max {
		generateParenthesisBacktrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		generateParenthesisBacktrack(ans, cur+")", open, close+1, max)
	}
}

// Search in Rotated Sorted Array
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	res := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		midValue := nums[mid]
		if midValue == target {
			res = mid
			break
		}
		if nums[low] <= midValue {
			if target >= nums[low] && target <= midValue {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > midValue && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return res
}

func ReverseListNode(head *listNode.ListNode) *listNode.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var previous *listNode.ListNode
	first := head
	current := head
	preceding := head.Next
	for preceding != nil {
		current.Next = previous
		previous = current
		current = preceding
		preceding = preceding.Next
	}
	current.Next = previous
	first = current
	return first
}

func CountBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		numStr := fmt.Sprintf("%032b", i)
		count := strings.Count(numStr, "1")
		res = append(res, count)
	}
	return res
}
