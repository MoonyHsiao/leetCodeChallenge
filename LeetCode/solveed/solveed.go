package solveed

import (
	"sort"
	"strconv"

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
