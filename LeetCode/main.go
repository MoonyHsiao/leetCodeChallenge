package main

import (
	"fmt"

	"github.com/lexhsiao135/ds-go/LeetCode/greedy"
	"github.com/lexhsiao135/ds-go/LeetCode/models/graphNode"
)

func main() {
	// data := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	// target := 5
	// data := 100

	// data := "[1,1,2]"
	// data2 := "[1,1,2]"
	// listdata := treeNode.CreateTestData(data)
	// listdata2 := treeNode.CreateTestData(data2)

	// for i := 0; i < 10; i++ {
	// 	res := dpProblem.Catalan(i)
	// 	fmt.Printf("%v ", res)
	// }
	// res := dpProblem.Binomial(2, 1)
	// fmt.Printf("res:%v\n", res)
	res := greedy.ChangeMoney(86)
	fmt.Printf("res:%v\n", res)

	aa := greedy.SimpleGraph()
	aa.String()

	aa.DFSTraverse(func(n *graphNode.Node) {
		fmt.Printf("node:%v ->", n)
	})

	fmt.Println()

	aa.BFSTraverse(func(n *graphNode.Node) {
		fmt.Printf("node:%v ->", n)
	})

	bb := greedy.WeightGraphinit()
	bb.AdjacentEdgesExample()
	fmt.Println()

}
