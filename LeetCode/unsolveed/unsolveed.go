package unsolveed

import (
	"fmt"
	"strconv"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/treeNode"
)

// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
func NaryLevelOrder(root *treeNode.NaryNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := &Queue{nodes: make([]*treeNode.NaryNode, 1)}
	q.Push(root)
	nums := []int{root.Val}
	level := 0
	res[level] = append(res[level], nums...)
	level++
	for q.GetLen() != 0 {
		curr := q.Pop()
		// fmt.Printf("val:%v\n", curr.Val)
		temp := curr.Children
		nums := []int{}
		for i := range temp {
			nums = append(nums, temp[i].Val)
			q.Push(temp[i])
		}

		if len(res) > level {
			res[level] = append(res[level], nums...)
		}
		level++
	}
	return res
}

type Queue struct {
	nodes []*treeNode.NaryNode
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n *treeNode.NaryNode) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*treeNode.NaryNode, len(q.nodes)*2)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *treeNode.NaryNode {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q *Queue) GetLen() int {
	return q.count
}

// https://leetcode.com/problems/water-and-jug-problem/submissions/
func CanMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	fmt.Printf("x:%v,y:%v,z:%v\n", x, y, z)

	nums := gcd(x, y)
	return CheckTwoSum(nums, z)
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

func gcd(n1, n2 int) []int {
	res := []int{}

	for n1 != 0 && n2 != 0 {
		if n1 >= n2 {
			n1 = n1 % n2
			res = append(res, n1)
		} else if n2 > n1 {
			n2 = n2 % n1
			res = append(res, n2)
		}
	}
	return res
}

// Your input
// [1,4,6,7,8,10]

// [2,7,15]

// stdout
// ('day1  your tickets=', 2)
// [0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0]
// This days you not traveling
// [0, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0]
// This days you not traveling
// [0, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0]
// ('day4 your tickets=', 2)
// [0, 2, 2, 2, 4, 0, 0, 0, 0, 0, 0]
// This days you not traveling
// [0, 2, 2, 2, 4, 4, 0, 0, 0, 0, 0]
// ('day6th  your tickets=', 2)
// [0, 2, 2, 2, 4, 4, 6, 0, 0, 0, 0]
// (' day7th your tickets=', 7)    #2+2+2+2 =£8 > £7
// [0, 2, 2, 2, 4, 4, 6, 7, 0, 0, 0]
// ('day 8th your tickets=', 2)    #2+2+2+2+2=£10 > £9
// [0, 2, 2, 2, 4, 4, 6, 7, 9, 0, 0]
// This days you not traveling
// [0, 2, 2, 2, 4, 4, 6, 7, 9, 9, 0]
// ('day 10th  your tickets=', 9) # Reason
// [0, 2, 2, 2, 4, 4, 6, 7, 9, 9, 9]
// https://leetcode.com/problems/minimum-cost-for-tickets/
func MincostTickets(days []int, costs []int) int {
	size := len(days) - 1
	maxsize := days[size]
	cost := make([]int, maxsize+1)
	dayCount := 0
	for i := 1; i < len(cost); i++ {
		money := cost[i-1]
		d7Count := 1
		d30Count := 1
		if i == days[dayCount] {
			temp := money + models.MinTriple(costs[0], costs[1], costs[2])
			for d7Count*costs[1] < temp && money >= d7Count*costs[1] {
				d7Count++
			}
			for d30Count*costs[2] < temp && money >= d30Count*costs[2] {
				d30Count++
			}
			money = models.MinTriple(temp, (d7Count)*costs[1], (d30Count)*costs[2])
			dayCount++
		}
		cost[i] = money
	}
	return cost[maxsize]
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
	resArray := make([][]int, len(nums))
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
		if len(m[i]) == 0 {
			m[i] = append(m[i], temp)
		}
	}
	for i := range m {
		temp := m[i]
		// if len(temp) == 1 {
		// 	continue
		// }
		// 缺一點組合
		for j := range temp {
			fmt.Printf("i:%v,temp:%v,j:%v\n", i, temp, temp[j])
			tempA := []int{nums[i]}
			tempA = append(tempA, m[j]...)
			// tempA = append(tempA, temp[j])
			fmt.Printf("tempA:%v\n", tempA)
			resArray[i] = append(resArray[i], tempA...)
		}
	}
	fmt.Printf("resArray:%v\n", resArray)
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("m:%v\n", m)
	//最後還要再濾一次
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

func CountIncreaseSubArray_0831(nums []int) int {
	m := make([][]int, len(nums))
	resArray := make([][]int, len(nums))
	resArrayStr := make([][]string, len(nums))
	// resArray := make(map[int]string)
	res := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		isMin := true
		for j := 0; j <= i; j++ {
			if nums[j] < temp {
				// m[i] = append(m[i], nums[j])
				m[i] = append(m[i], j)
				isMin = false
			} else {
				m[i] = append(m[i], -1)
			}
		}
		if isMin {
			m[i] = nil
			m[i] = append(m[i], nums[i])
		}
		// if len(m[i]) == 0 {
		// 	m[i] = append(m[i], temp)
		// }
	}
	for i := range m {
		temp := m[i]
		if len(temp) == 1 {
			// resArray[i] = strconv.Itoa(temp[0])
			resArray[i] = append(resArray[i], temp[0])
			continue
		}
		for j := range temp {
			if temp[j] != -1 {
				index := temp[j]
				aaaa := nums[index]
				// set := resArray[index]
				// fmt.Printf("set:%v\n", set)
				bbbb := nums[i]
				fmt.Printf("i:%v,temp:%v,aaaa:%v,bbbb:%v\n", i, temp, aaaa, bbbb)
				// cccc := resArray[index]
				// str := fmt.Sprintf("%v,%v", bbbb, cccc)
				// resArray[i] = str
				resArray[i] = append(resArray[i], aaaa)
			}

		}

	}
	for i := range m {
		temp := m[i]
		if len(temp) == 1 {
			resArrayStr[i] = append(resArrayStr[i], strconv.Itoa(temp[0]))
			continue
		}
		for j := range temp {
			if temp[j] != -1 {
				index := temp[j]
				aaaaa := strconv.Itoa(index)
				bbbbb := resArrayStr[index]
				// ccccc := strings.Split(bbbbb, ",")
				ccccc := "11"
				fmt.Printf("aaaaa:%v,bbbbb:%v,,ccccc:%v\n", aaaaa, bbbbb, ccccc)
			}
		}

		// else {
		// 	fmt.Printf("resArrayStr:%v\n", resArrayStr[i])
		// }
	}
	// 把resArray再拿去重組答案
	fmt.Printf("resArray:%v\n", resArray)
	fmt.Printf("resArrayStr:%v\n", resArrayStr)
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("m:%v\n", m)
	//最後還要再濾一次
	return res

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

func test() {
	res := 0
	fmt.Print(res)
	fmt.Printf("res:%v\n", res)
}
