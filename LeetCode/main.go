package main

import (
	"fmt"

	"github.com/lexhsiao135/ds-go/LeetCode/greedy"
	"github.com/lexhsiao135/ds-go/LeetCode/models/listNode"
	"github.com/lexhsiao135/ds-go/LeetCode/models/treeNode"
	"github.com/lexhsiao135/ds-go/LeetCode/solveed"
)

func main() {
	// twoListNode()
	// twoTreeNode()
	// graph()
	intArray()
	// strArray()
	// checkSingleInt()
	// oneListNode()

	// a := solveed.ConvertToTitle(52)
	// fmt.Printf("res:%v\n", a)

	// data := []int{2, 7, 11, 14}
	// t := 16
	// i1, i2 := solveed.FindTwoSumIndex(data, t)
	// fmt.Print(i1)
	// fmt.Print(i2)

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
	data := []int{2, 7, 4, 13, 6}
	// target := 8
	// k := 3
	res := solveed.SortArrayByParityV2(data)
	fmt.Printf("res:%v\n", res)
}

func strArray() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := solveed.GroupAnagrams(data)
	fmt.Printf("res:%v\n", res)
}

func checkSingleInt() {
	data := 19
	res := solveed.TrailingZeroes(data)
	fmt.Println(res)
}

func oneListNode() {
	data := "[1,2,3,4,5,6]"
	listdata := listNode.CreateTestData(data)
	res := solveed.MiddleNode(listdata)
	fmt.Printf("res:%v\n", res)
}
