package greedy

import "github.com/lexhsiao135/ds-go/LeetCode/models/graphNode"

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
	wg.AddEdge(0, 1, 10)
	wg.AddEdge(0, 5, 11)
	wg.AddEdge(1, 2, 18)
	wg.AddEdge(1, 6, 16)
	wg.AddEdge(1, 8, 12)
	wg.AddEdge(2, 3, 22)
	wg.AddEdge(2, 8, 8)
	wg.AddEdge(3, 4, 20)
	wg.AddEdge(3, 6, 24)
	wg.AddEdge(3, 7, 16)
	wg.AddEdge(3, 8, 21)
	wg.AddEdge(4, 5, 26)
	wg.AddEdge(4, 7, 7)
	wg.AddEdge(5, 6, 17)
	wg.AddEdge(6, 7, 19)
	return wg
}
