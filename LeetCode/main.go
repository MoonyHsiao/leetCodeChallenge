package main

import (
	"fmt"
	"time"

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

// https://golang.org/pkg/fmt/

func main() {
	// listNodeType()
	treeNodeType()
	// graph()
	// intArray()
	// strArray()
	// byteArray()
	// checkSingleInt()
	// compareString()
	// testRobot()
	// sortArray()
	// stackArray()
	// queueArray()
	// searchArray()
	// dp()
	// islands()
	// models.LearnBits()

}

func listNodeType() {
	data := "[3,5]"
	// data2 := "[1,5]"
	// target := 1
	listdata := listNode.CreateTestData(data)
	// listdata2 := listNode.CreateTestData(data2)
	res := unsolveed.ReverseBetween(listdata, 1, 2)
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

func treeNodeType() {
	data := "[0,-1]"
	// data2 := "[1,1,2]"
	listdata := treeNode.CreateTestData(data)
	// listdata2 := treeNode.CreateTestData(data2)
	res := solveed.LargestValuesNotPointer(listdata)
	fmt.Printf("res:%v\n", res)

}

func intArray() {
	// data := []int{1, 3}
	data := 5
	// target := 3
	res := solveed.CountBits(data)
	fmt.Printf("res:%v\n", res)
}

func strArray() {
	// data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// data := []string{"4", "5", "DUP", "2", "+", "8", "-", "POP"}
	data := "172.16.204.1"
	// data := "20EE:Fb8:85a3:0:0:8A2E:0370:7334:12"
	res := solveed.ValidIPAddress(data)
	fmt.Printf("res:%v\n", res)

}

func byteArray() {
	data := [][]byte{}
	obstacle1 := []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	obstacle2 := []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	obstacle3 := []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	obstacle4 := []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	obstacle5 := []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	obstacle6 := []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	obstacle7 := []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	obstacle8 := []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	obstacle9 := []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	data = append(data, obstacle1)
	data = append(data, obstacle2)
	data = append(data, obstacle3)
	data = append(data, obstacle4)
	data = append(data, obstacle5)
	data = append(data, obstacle6)
	data = append(data, obstacle7)
	data = append(data, obstacle8)
	data = append(data, obstacle9)
	res := solveed.IsValidSudoku(data)
	fmt.Printf("res:%v\n", res)

}

func checkSingleInt() {
	data := 4000
	res := solveed.CountPrimes(data)
	fmt.Printf("res:%v\n", res)
}

func compareString() {
	data1 := "AAAAA"
	data2 := "BB"
	res := solveed.MergeString(data1, data2)
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

func dp() {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}
	// days := []int{1, 4, 6, 7, 8, 10}
	// costs := []int{2, 7, 15}
	start := time.Now()
	res := unsolveed.MincostTickets(days, costs)
	elapsed := time.Since(start)
	fmt.Printf("res:%v\n", res)
	fmt.Printf("Took %s\n", elapsed)

}

func islands() {
	data := [][]int{}
	obstacle1 := []int{1, 2}
	obstacle2 := []int{5, 6}
	obstacle3 := []int{1, 1}
	data = append(data, obstacle1)
	data = append(data, obstacle2)
	data = append(data, obstacle3)
	res := solveed.MinPathSum(data)
	fmt.Printf("res:%v\n", res)

}

func intRes() {
	x := 22003
	y := 31237
	z := 123
	res := unsolveed.CanMeasureWater(x, y, z)
	fmt.Printf("res:%v\n", res)
}
