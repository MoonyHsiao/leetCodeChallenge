package greedy

import (
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/graphNode"
)

func maxDivide(a, b int) (remaining, coin int) {
	c := 0
	for a >= b {
		a = a - b
		c++
	}
	return a, c
}

func ChangeMoney(money int) []int {
	changes := []int{50, 20, 10, 5, 1}
	coins := []int{0, 0, 0, 0, 0}
	coin := 0
	for i := range changes {
		if money >= changes[i] {
			money, coin = maxDivide(money, changes[i])
			coins[i] = coin
		}
	}
	return coins
}

func SimpleGraph() *graphNode.ItemGraph {
	var g graphNode.ItemGraph

	nA := graphNode.Node{"A"}
	nB := graphNode.Node{"B"}
	nC := graphNode.Node{"C"}
	nD := graphNode.Node{"D"}
	nE := graphNode.Node{"E"}
	nF := graphNode.Node{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nD, &nA)
	g.AddEdge(&nE, &nF)
	return &g
}

func WeightGraphinit() *graphNode.WeightGraph {
	wg := graphNode.NewWeightGraph(9)
	wg.AddEdge(0, 1, 7)
	wg.AddEdge(0, 3, 5)
	wg.AddEdge(1, 2, 8)
	wg.AddEdge(1, 3, 9)
	wg.AddEdge(1, 4, 7)
	wg.AddEdge(2, 4, 5)
	wg.AddEdge(3, 4, 15)
	wg.AddEdge(3, 5, 6)
	wg.AddEdge(4, 5, 8)
	wg.AddEdge(4, 6, 9)
	wg.AddEdge(5, 6, 11)
	return wg
}

func WeightGraphinitV2() *graphNode.WeightGraph {
	wg := graphNode.NewWeightGraph(9)
	wg.AddEdge(0, 1, 5)
	wg.AddEdge(0, 5, 3)
	wg.AddEdge(1, 2, 10)
	wg.AddEdge(1, 4, 1)
	wg.AddEdge(1, 6, 4)
	wg.AddEdge(2, 3, 5)
	wg.AddEdge(2, 6, 8)
	wg.AddEdge(3, 4, 7)
	wg.AddEdge(3, 6, 9)
	wg.AddEdge(4, 5, 6)
	wg.AddEdge(4, 6, 2)

	return wg
}

func Kruskal(wg *graphNode.WeightGraph) []graphNode.WEdge {

	subset := [][]int{}
	for i := 0; i < 7; i++ {
		array := []int{i}
		subset = append(subset, array)
	}

	edgesetMST := []graphNode.WEdge{}
	list := wg.GetSortedEdges()
	// fmt.Printf("sort list:%v", list)
	i := 0
	for len(edgesetMST) != 6 {
		e := list[i]
		fromIndex := findIndex(subset, e.From)
		toIndex := findIndex(subset, e.To)
		if fromIndex != toIndex {
			edgesetMST = append(edgesetMST, e)
			for i := range subset[toIndex] {
				subset[fromIndex] = append(subset[fromIndex], subset[toIndex][i])
			}
			subset = RemoveArrayIndex(subset, toIndex)
		}
		i++
	}
	return edgesetMST
}

func findIndex(subset [][]int, target int) int {
	res := -1
	for i := range subset {
		temp := subset[i]
		for j := range temp {
			if target == temp[j] {
				res = i
				break
			}
		}
	}
	return res
}

func RemoveArrayIndex(s [][]int, index int) [][]int {
	return append(s[:index], s[index+1:]...)
}
