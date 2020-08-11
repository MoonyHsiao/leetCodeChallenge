package solveed

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/listNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/treeNode"
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

func DeleteDuplicates(head *listNode.ListNode) *listNode.ListNode {
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
func IsSameTree(p *treeNode.TreeNode, q *treeNode.TreeNode) bool {

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

func MaxDepth(root *treeNode.TreeNode) int {
	return countDepthMax(root, 1)
}

func countDepthMax(root *treeNode.TreeNode, count int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return countDepthMax(root.Right, count+1)
	}

	if root.Right == nil && root.Left != nil {
		return countDepthMax(root.Left, count+1)
	}

	if root.Right != nil && root.Left != nil {
		return models.Max(countDepthMax(root.Left, count+1), countDepthMax(root.Right, count+1))
	}
	return count
}

func CheckTwoSum(s []int, target int) bool {
	for i := range s {
		temp := target - s[i]
		tempS := s[i:len(s)]
		check := checkAnother(tempS, temp)
		if check {
			return true
		}
	}
	return false
}

func checkAnother(s []int, t int) bool {
	for i := range s {
		if s[i] == t {
			return true
		}
	}
	return false
}

func HasPathSum(root *treeNode.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	temp := sum - root.Val
	fmt.Printf("root:%v\n", root)
	fmt.Printf("temp:%v\n", temp)
	fmt.Printf("root.Left:%v\n", root.Left)
	fmt.Printf("root.Right:%v\n", root.Right)
	fmt.Printf("=========\n")
	//最終終止條件 這沒問題
	if temp == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return HasPathSum(root.Left, temp) || HasPathSum(root.Right, temp)
}

func LevelOrderBottom(root *treeNode.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := []Rest{}
	max := 0
	isRoot(root, 1, &res, &max)
	ans := make([][]int, max)
	for i := range res {
		temp := res[i]
		ans[temp.Level] = append(ans[temp.Level], temp.Val)
	}
	return reverseInts(ans[1:])
}

func reverseInts(input [][]int) [][]int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
func isRoot(root *treeNode.TreeNode, level int, res *[]Rest, max *int) {

	temp := Rest{}
	temp.Val = root.Val
	temp.Level = level
	*res = append(*res, temp)
	level++
	if *max < level {
		*max = level
	}
	if root.Left != nil {
		isRoot(root.Left, level, res, max)
	}
	if root.Right != nil {
		isRoot(root.Right, level, res, max)
	}

}

type Rest struct {
	Val   int
	Level int
}

func IsSymmetric(root *treeNode.TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}
	res1 := []int{}
	res2 := []int{}
	isSymmetricLeft(root.Left, &res1)
	isSymmetricRight(root.Right, &res2)

	if len(res1) != len(res2) {
		return false
	}
	for i := range res1 {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true

}

func isSymmetricLeft(root *treeNode.TreeNode, s *[]int) {
	if root == nil {
		*s = append(*s, 0)
	} else {
		*s = append(*s, root.Val)
		isSymmetricLeft(root.Left, s)
		isSymmetricLeft(root.Right, s)
	}
}

func isSymmetricRight(root *treeNode.TreeNode, s *[]int) {

	if root == nil {
		*s = append(*s, 0)
	} else {
		*s = append(*s, root.Val)
		isSymmetricRight(root.Right, s)
		isSymmetricRight(root.Left, s)
	}
}

func MaxProfit(prices []int) int {
	p := 0
	for i := range prices {
		temp := prices[i:]
		max := 0
		for j := range temp {
			if temp[j]-prices[i] > max {
				max = temp[j] - prices[i]
			}

		}
		if max > p {
			p = max
		}
	}
	return p
}

func MinDepth(root *treeNode.TreeNode) int {
	return countDepthMin(root, 1)
}

func countDepthMin(root *treeNode.TreeNode, count int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return countDepthMin(root.Right, count+1)
	}

	if root.Right == nil && root.Left != nil {
		return countDepthMin(root.Left, count+1)
	}

	if root.Right != nil && root.Left != nil {
		return models.Min(countDepthMin(root.Left, count+1), countDepthMin(root.Right, count+1))
	}
	return count
}

func Generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		temp := []int{}
		for j := 0; j <= i; j++ {
			val := Binomial(i, j)
			temp = append(temp, val)
		}
		res = append(res, temp)
	}
	return res
}

func Binomial(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return Binomial(n-1, k-1) + Binomial(n-1, k)
}

func TitleToNumber(s string) int {
	temp := []int{}
	for i := range s {
		l1 := int(s[i]) - 64
		temp = append(temp, l1)
	}

	sum := 0
	mult := 1
	for i := len(temp) - 1; i >= 0; i-- {
		mult = mult * 26
		if i == len(temp)-1 {
			mult = 1
		}
		sum += temp[i] * mult
	}
	return sum
}

func GetIntersectionNode(headA, headB *listNode.ListNode) *listNode.ListNode {
	countA := 0
	tempA := headA
	for tempA != nil {
		countA++
		tempA = tempA.Next
	}
	countB := 0
	tempB := headB
	for tempB != nil {
		countB++
		tempB = tempB.Next
	}

	temp := &listNode.ListNode{}
	if countA > countB {
		pass := countA - countB
		for i := 0; i < pass; i++ {
			headA = headA.Next
		}
	} else if countB > countA {
		pass := countB - countA
		for i := 0; i < pass; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			temp = headA
			break
		}
		headA = headA.Next
		headB = headB.Next
	}

	for headA != nil {
		if headA != headB {
			return nil
		}

		headA = headA.Next
		headB = headB.Next
	}
	if temp.Val == 0 {
		return nil
	}
	return temp
}

func CountElements(arr []int) int {
	res := 0
	m := make(map[int]int)
	keys := make([]int, 0, len(m))
	for i := range arr {
		_, ok := m[arr[i]]
		if ok {
			m[arr[i]] += 1
		} else {
			temp := arr[i]
			m[temp] = 1
			keys = append(keys, temp)
		}
	}

	sort.Ints(keys)
	for i := 0; i < len(keys)-1; i++ {
		if keys[i]+1 == keys[i+1] {
			res += m[keys[i]]
		}
	}

	return res
}

func GroupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for i := range strs {
		temp := strings.Split(strs[i], "")
		sort.Strings(temp)
		res := strings.Join(temp, "")
		m_temp := m[res]
		m_temp = append(m_temp, strs[i])
		m[res] = m_temp
	}

	for k, _ := range m {
		key := m[k]
		res = append(res, key)
	}

	return res
}

func IsHappy(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n == 4 || n == 16 || n == 37 || n == 58 || n == 89 || n == 145 || n == 42 || n == 20 {
		return false
	}

	result := sqrt(n)

	return IsHappy(result)
}

func sqrt(n int) int {

	sum := 0
	for n >= 10 {
		t := n % 10
		sum += t * t
		n = n / 10
	}
	sum += n * n

	return sum

}

func MaxSubArray(nums []int) int {
	max := 0
	currentMax := 0
	isAllNega := isAllNega(nums)
	if isAllNega {
		max = nums[0]
		for i := range nums {
			if max < nums[i] {
				max = nums[i]
			}
		}

	} else {
		for i := range nums {
			currentMax += nums[i]
			if currentMax < 0 {
				currentMax = 0
			}
			if max < currentMax {
				max = currentMax
			}
		}
	}
	return max
}

func isAllNega(nums []int) bool {
	isNega := true
	for i := range nums {
		if nums[i] > 0 {
			isNega = false
			break
		}
	}
	return isNega
}

func MoveZeroes(nums []int) {
	for i := range nums {
		if nums[i] == 0 {
			slice := nums[i:]
			for j := range slice {
				if slice[j] != 0 {
					value := slice[j]
					nums[i] = value
					slice[j] = 0
					break
				}
			}
		}
	}
}

func SingleNumber(nums []int) int {

	numsMap := make(map[int]int)
	for _, e := range nums {
		_, isExist := numsMap[e]
		if isExist {
			numsMap[e] += 1
		} else {
			numsMap[e] = 1
		}
	}
	result := 0
	for k, v := range numsMap {
		if v == 1 {
			result = k
			break
		}
	}
	return result
}

func MiddleNode(head *listNode.ListNode) *listNode.ListNode {
	count := 1
	res := head
	for {
		if head.Next == nil {
			break
		} else {
			count += 1
			head = head.Next
		}
	}
	upper := count / 2

	for i := 0; i < upper; i++ {
		res = res.Next
	}

	return res
}

func Rotate(nums []int, k int) {
	end := len(nums) - 1
	for i := 0; i < k; i++ {
		for j := end; j > 0; j-- {
			swapPointArray(&nums, j, j-1)
		}
	}
}

func swapPointArray(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func ConvertToTitle(n int) string {
	strArray := modASCII(n)
	str := []string{}
	for i := range strArray {
		asciiValue := strArray[i] + 64
		temp := rune(asciiValue)
		str = append(str, string(temp))
	}
	res := strings.Join(str, "")
	return res
}

func modASCII(num int) []int {
	res := []int{}
	for num > 0 {
		if num%26 == 0 {
			res = append(res, 26)
			num -= 26
		} else {
			res = append(res, num%26)
			num -= num % 26
		}
		num = num / 26
	}

	return models.ReverseInts(res)
}
