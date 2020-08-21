package solveed

import (
	"fmt"
	"math"
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

func TrailingZeroes(n int) int {
	res := 0
	if n == 0 {
		return res
	}

	var i float64
	i = 1
	for n >= int(math.Pow(5, i)) {
		var temp float64
		temp = float64(n)
		res += int(temp / (math.Pow(5, i)))
		i++
	}
	return res

}

func FindTwoSumIndex(s []int, target int) (a, b int) {
	maxIndex := len(s) - 1
	first := -1
	sec := -1
	for i := range s {
		if target < s[i] {
			maxIndex = i - 1
			break
		}
	}

	isFind := false
	for i := 0; i < maxIndex; i++ {
		temp := target - s[i]

		for j := maxIndex; j > 0; j-- {
			if temp == s[j] {
				first = i
				sec = j
				isFind = true
				break
			}
		}
		if isFind {
			break
		}
	}
	return first, sec
}

func Rob(nums []int) int {
	prevGetMax := 0
	prevNotGetMax := 0
	for i := range nums {
		temp := prevGetMax
		prevGetMax = nums[i] + prevNotGetMax
		if temp > prevNotGetMax {
			prevNotGetMax = temp
		}
	}
	return models.Max(prevNotGetMax, prevGetMax)
}

func FindTwoSumIndexV2(numbers []int, target int) []int {
	saved := make(map[int]int)
	first := -1
	sec := -1
	for i := range numbers {
		_, ok := saved[numbers[i]]
		if !ok {
			temp := target - numbers[i]
			saved[temp] = i
		} else {
			temp := numbers[i]
			first = saved[temp]
			sec = i
			break
		}
	}
	res := []int{first + 1, sec + 1}
	return res
}

func HammingWeight(num uint32) int {

	trans := fmt.Sprintf("%b", num)
	count := 0
	for i := range trans {
		if string(trans[i]) == "1" {
			count++
		}
	}
	return count
}

func JudgeCircle(moves string) bool {
	U := strings.Count(moves, "U")
	D := strings.Count(moves, "D")
	L := strings.Count(moves, "L")
	R := strings.Count(moves, "R")
	return U == D && L == R
}

func SortArrayByParity(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return MergeArrayByParity(SortArrayByParity(slice[:mid]), SortArrayByParity(slice[mid:]))
}

func MergeArrayByParity(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i]%2 == 0 {
			slice[k] = left[i]
			i++
		} else if right[j]%2 == 0 {
			slice[k] = right[j]
			j++
		} else if right[j]%2 != 0 && left[i]%2 != 0 {
			k1 := k
			for i1 := i; i1 < len(left); i1++ {
				slice[k1] = left[i1]
				k1++
			}
			for l1 := j; l1 < len(right); l1++ {
				slice[k1] = right[l1]
				k1++
			}
			break
		}
	}
	return slice
}

func SortArrayByParityV2(A []int) []int {
	odd := []int{}
	even := []int{}
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}
	return append(even, odd...)
}

func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int)
	res := -1
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]] += 1
			if m[nums[i]] > len(nums)/2 {
				res = nums[i]
				break
			}
		} else {
			temp := nums[i]
			m[temp] = 1
		}
	}
	return res
}

func CountPrimes(n int) int {
	primeArray := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997, 1009, 1013, 1019, 1021, 1031, 1033, 1039, 1049, 1051, 1061, 1063, 1069, 1087, 1091, 1093, 1097, 1103, 1109, 1117, 1123, 1129, 1151, 1153, 1163, 1171, 1181, 1187, 1193, 1201, 1213, 1217, 1223, 1229, 1231, 1237, 1249, 1259, 1277, 1279, 1283, 1289, 1291, 1297, 1301, 1303, 1307, 1319, 1321, 1327, 1361, 1367, 1373, 1381, 1399, 1409, 1423, 1427, 1429, 1433, 1439, 1447, 1451, 1453, 1459, 1471, 1481, 1483, 1487, 1489, 1493, 1499, 1511, 1523, 1531, 1543, 1549, 1553, 1559, 1567, 1571, 1579, 1583, 1597, 1601, 1607, 1609, 1613, 1619, 1621, 1627, 1637, 1657, 1663, 1667, 1669, 1693, 1697, 1699, 1709, 1721, 1723, 1733, 1741, 1747, 1753, 1759, 1777, 1783, 1787, 1789, 1801, 1811, 1823, 1831, 1847, 1861, 1867, 1871, 1873, 1877, 1879, 1889, 1901, 1907, 1913, 1931, 1933, 1949, 1951, 1973, 1979, 1987, 1993, 1997, 1999, 2003, 2011, 2017, 2027, 2029, 2039, 2053, 2063, 2069, 2081, 2083, 2087, 2089, 2099, 2111, 2113, 2129, 2131, 2137, 2141, 2143, 2153, 2161, 2179, 2203, 2207, 2213, 2221, 2237, 2239, 2243, 2251, 2267, 2269, 2273, 2281, 2287, 2293, 2297, 2309, 2311, 2333, 2339, 2341, 2347, 2351, 2357, 2371, 2377, 2381, 2383, 2389, 2393, 2399, 2411, 2417, 2423, 2437, 2441, 2447, 2459, 2467, 2473, 2477, 2503, 2521, 2531, 2539, 2543, 2549, 2551, 2557, 2579, 2591, 2593, 2609, 2617, 2621, 2633, 2647, 2657, 2659, 2663, 2671, 2677, 2683, 2687, 2689, 2693, 2699, 2707, 2711, 2713, 2719, 2729, 2731, 2741, 2749, 2753, 2767, 2777, 2789, 2791, 2797, 2801, 2803, 2819, 2833, 2837, 2843, 2851, 2857, 2861, 2879, 2887, 2897, 2903, 2909, 2917, 2927, 2939, 2953, 2957, 2963, 2969, 2971, 2999, 3001, 3011, 3019, 3023, 3037, 3041, 3049, 3061, 3067, 3079, 3083, 3089, 3109, 3119, 3121, 3137, 3163, 3167, 3169, 3181, 3187, 3191, 3203, 3209, 3217, 3221, 3229, 3251, 3253, 3257, 3259, 3271, 3299, 3301, 3307, 3313, 3319, 3323, 3329, 3331, 3343, 3347, 3359, 3361, 3371, 3373, 3389, 3391, 3407, 3413, 3433, 3449, 3457, 3461, 3463, 3467, 3469, 3491, 3499, 3511, 3517, 3527, 3529, 3533, 3539, 3541, 3547, 3557, 3559, 3571, 3581, 3583, 3593, 3607, 3613, 3617, 3623, 3631, 3637, 3643, 3659, 3671, 3673, 3677, 3691, 3697, 3701, 3709, 3719, 3727, 3733, 3739, 3761, 3767, 3769, 3779, 3793, 3797, 3803, 3821, 3823, 3833, 3847, 3851, 3853, 3863, 3877, 3881, 3889, 3907, 3911, 3917, 3919, 3923, 3929, 3931, 3943, 3947, 3967, 3989}
	count := 0
	isInArray := false
	primeArrayMax := primeArray[len(primeArray)-1]
	if n < primeArrayMax {
		isInArray = true
	}
	//先檢查這個值有沒有在array裡面 沒有就生成到有
	if !isInArray {
		for i := primeArrayMax + 1; i < n; i++ {
			ok := isPrimes(i, &primeArray)
			if ok {
				primeArray = append(primeArray, i)
			}
		}

	}

	//最後檢查
	for i := range primeArray {
		if n > primeArray[i] {
			count++

		} else {
			break
		}
	}
	return count
}

func isPrimes(n int, primeArray *[]int) bool {
	isPrime := true
	temp := math.Sqrt(float64(n))
	for i := 0; i < int(temp); i++ {
		if n%(*primeArray)[i] == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}

func GenPrimes(n int) []string {
	primeArray := []int{2}
	primeArrayS := []string{"2"}
	for i := 2; i < n; i++ {
		isPrime := true
		for j := range primeArray {
			if i%primeArray[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeArray = append(primeArray, i)
			primeArrayS = append(primeArrayS, strconv.Itoa(i))
		}
	}
	return primeArrayS
}

func IsIsomorphic(s string, t string) bool {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		// fmt.Printf("s[i]:%v,t[i]:%v\n", s[i], t[i])
		// fmt.Printf("sMap[s[i]] :%v,tMap[t[i]]:%v\n", sMap[s[i]], tMap[t[i]])
		if sMap[s[i]] == tMap[t[i]] {
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		} else {
			return false
		}
	}
	return true
}

func ReverseList(head *listNode.ListNode) *listNode.ListNode {
	res := &listNode.ListNode{}
	resHead := res
	arr := []int{}
	if head == nil {
		return nil
	}

	for head != nil {
		val := head.Val
		arr = append(arr, val)
		head = head.Next
	}

	for i := len(arr) - 1; i >= 0; i-- {
		res.Val = arr[i]
		temp := &listNode.ListNode{}
		if i == 0 {
			break
		}
		res.Next = temp
		res = res.Next
	}
	return resHead
}

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			return true
		} else {
			temp := nums[i]
			m[temp] = 1
		}
	}
	return false
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	size, i, j := len(nums1), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i >= m && j <= n {
			slice[k] = nums2[j]
			j++
		} else if j >= n && i <= m {
			slice[k] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			slice[k] = nums1[i]
			i++
		} else {
			slice[k] = nums2[j]
			j++
		}
	}
	for i := range nums1 {
		nums1[i] = slice[i]
	}
}

func InvertTree(root *treeNode.TreeNode) *treeNode.TreeNode {
	if root == nil {
		return nil
	}
	temp := InvertTree(root.Right)
	root.Right = InvertTree(root.Left)
	root.Left = temp
	return root
}

// 感覺是設計上的問題 還不習慣這樣設計del
func DeleteNode(node *listNode.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[string]string)
	for i := range s {
		temp := string(s[i])
		_, exist := set[temp]
		if !exist {
			set[temp] = temp
		}
	}

	for k := range set { // Loop
		sample := k
		sCount := strings.Count(s, sample)
		tCount := strings.Count(t, sample)
		if sCount != tCount {
			return false
		}
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsAnagramV2(s string, t string) bool {
	s1 := SortString(s)
	t1 := SortString(t)
	return s1 == t1
}
