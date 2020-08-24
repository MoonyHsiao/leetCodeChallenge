package main

import (
	"fmt"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/greedy"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/listNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/queue"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/search"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/stack"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/treeNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/solveed"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/unsolveed"
)

func main() {
	// twoListNode()
	// twoTreeNode()
	// graph()
	intArray()
	// strArray()
	// checkSingleInt()
	// oneListNode()
	// compareString()
	// testRobot()
	// singleTreeNode()
	// sortArray()
	// stackArray()
	// queueArray()
	// searchArray()
	// data := 1000000
	// res := unsolveed.MinimumGreaterThanGenAnswer(data)
	// fmt.Printf("res:%v\n", res)

}

func twoListNode() {
	data := "[2,6,4]"
	data2 := "[1,5]"
	listdata := listNode.CreateTestData(data)
	listdata2 := listNode.CreateTestData(data2)
	res := solveed.GetIntersectionNode(listdata, listdata2)
	fmt.Printf("res:%v\n", res)
}

func graph() {
	bb := greedy.WeightGraphinit()
	bb.AdjacentEdgesExample()
	bb.GetSortedEdges()
	fmt.Println()
	res := greedy.Kruskal(bb)
	fmt.Printf("res:%v\n", res)
}

func twoTreeNode() {

	data := "[1,1,2]"
	data2 := "[1,1,2]"
	listdata := treeNode.CreateTestData(data)
	listdata2 := treeNode.CreateTestData(data2)
	fmt.Print(listdata)
	fmt.Print(listdata2)
}

func intArray() {
	data := []int{1, 2, 3, 4, 7}
	target := 3
	res := unsolveed.MaxDistance(data, target)
	fmt.Printf("res:%v\n", res)
}

func strArray() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := solveed.GroupAnagrams(data)
	fmt.Printf("res:%v\n", res)
}

func checkSingleInt() {
	data := 4000
	res := solveed.CountPrimes(data)
	fmt.Printf("res:%v\n", res)
}

func oneListNode() {
	data := "[1,2,3,4,5]"
	listdata := listNode.CreateTestData(data)
	res := solveed.ReverseList(listdata)
	fmt.Printf("res:%v\n", res)
}

func compareString() {
	data1 := "abb"
	data2 := "bab"
	res := solveed.IsIsomorphic(data1, data2)
	fmt.Printf("res:%v\n", res)
}

func testRobot() {
	commands := []int{-2, -1, 8, 9, 6}
	obstacles := [][]int{}
	obstacle1 := []int{-1, 3}
	obstacle2 := []int{0, 1}
	obstacle3 := []int{-1, 5}
	obstacle4 := []int{-2, -4}
	obstacle5 := []int{5, 4}
	obstacle6 := []int{-2, -3}
	obstacle7 := []int{5, -1}
	obstacle8 := []int{1, -1}
	obstacle9 := []int{5, 5}
	obstacle10 := []int{5, 2}
	obstacles = append(obstacles, obstacle1)
	obstacles = append(obstacles, obstacle2)
	obstacles = append(obstacles, obstacle3)
	obstacles = append(obstacles, obstacle4)
	obstacles = append(obstacles, obstacle5)
	obstacles = append(obstacles, obstacle6)
	obstacles = append(obstacles, obstacle7)
	obstacles = append(obstacles, obstacle8)
	obstacles = append(obstacles, obstacle9)
	obstacles = append(obstacles, obstacle10)
	res := solveed.RobotSim(commands, obstacles)
	fmt.Printf("res:%v\n", res)
}

func singleTreeNode() {
	data := "[4,2,7,1,3,6,9]"
	listdata := treeNode.CreateTestData(data)
	res := solveed.InvertTree(listdata)
	fmt.Printf("res:%v\n", res)
}

func sortArray() {
	data := []int{1, 3, 2, 3, 4, 5}
	sort.QuickSort(&data, 0, len(data)-1)
	fmt.Printf("res:%v\n", data)
}

func stackArray() {
	var stringStack stack.StringStack

	stringStack.Push("+")
	stringStack.Push("/")
	stringStack.Push("*")
	fmt.Printf("stringStack.GetLen:%v\n", stringStack.GetLen())
	fmt.Printf("stringStack.Peek:%v\n", stringStack.Peek())
	// stringStack.Clear()
	// fmt.Printf("stringStack.GetLen:%v\n", stringStack.GetLen())
	// fmt.Printf("stringStack.Peek:%v\n", stringStack.Peek())
	for len(stringStack) > 0 {
		x, y := stringStack.Pop()
		if y == true {
			fmt.Printf("stringStack.Pop:%v\n", x)
		}
	}
}

func queueArray() {
	var stringQueue queue.StringQueue

	stringQueue.Push("+")
	stringQueue.Push("/")
	stringQueue.Push("*")
	fmt.Printf("stringQueue.GetLen:%v\n", stringQueue.GetLen())
	fmt.Printf("stringQueue.PeekFront:%v\n", stringQueue.PeekFront())
	fmt.Printf("stringQueue.PeekBack:%v\n", stringQueue.PeekBack())
	// stringQueue.Clear()
	// fmt.Printf("stringQueue.GetLen:%v\n", stringQueue.GetLen())
	// fmt.Printf("stringQueue.Peek:%v\n", stringQueue.Peek())
	for len(stringQueue) > 0 {
		x, y := stringQueue.Pop()
		if y == true {
			fmt.Printf("stringStack.Pop:%v\n", x)
		}
	}
}

func searchArray() {
	data := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 2
	// k := 3
	res := search.BinarySreach(data, target)
	fmt.Printf("res:%v\n", res)
}
