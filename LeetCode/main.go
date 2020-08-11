package main

import (
	"fmt"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/greedy"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/listNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/treeNode"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/solveed"
)

func main() {
	twoListNode()
	twoTreeNode()
	graph()
	intArray()
	strArray()
	checkSingleInt()
	oneListNode()

	a := solveed.ConvertToTitle(52)
	fmt.Printf("res:%v\n", a)

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
	data := []int{1, 2, 3, 4, 5, 6, 7}
	// target := 5
	k := 3
	solveed.Rotate(data, k)
	// fmt.Printf("res:%v\n", res)
}

func strArray() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := solveed.GroupAnagrams(data)
	fmt.Printf("res:%v\n", res)
}

func checkSingleInt() {
	data := 19
	res := solveed.IsHappy(data)
	fmt.Println(res)
}

func oneListNode() {
	data := "[1,2,3,4,5,6]"
	listdata := listNode.CreateTestData(data)
	res := solveed.MiddleNode(listdata)
	fmt.Printf("res:%v\n", res)
}
