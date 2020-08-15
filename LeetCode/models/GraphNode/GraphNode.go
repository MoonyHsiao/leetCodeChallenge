package graphNode

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/cheekybits/genny/generic"
	"github.com/lexhsiao135/ds-go/LeetCode/models"
)

type Item generic.Type

// Node a single node that composes the tree
type Node struct {
	Value Item
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) String() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

var visited_dfs map[*Node]bool

func (g *ItemGraph) DFSTraverse(f func(*Node)) {
	g.lock.RLock()
	visited_dfs = make(map[*Node]bool)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(g.nodes))
	index = 0
	head := g.nodes[index]
	g.dfs(head, f)

	g.lock.RUnlock()
}

func (g *ItemGraph) dfs(node *Node, f func(*Node)) {
	if f != nil {
		f(node)
	}
	visited_dfs[node] = true
	near := g.edges[*node]
	for i := 0; i < len(near); i++ {
		j := near[i]
		if !visited_dfs[j] {
			g.dfs(j, f)
		}
	}
}

// BFSTraverse
func (g *ItemGraph) BFSTraverse(f func(*Node)) {
	g.lock.RLock()
	q := NodeQueue{}
	q.New()
	n := g.nodes[0]
	q.Enqueue(*n)
	visited := make(map[*Node]bool)
	visited[n] = true
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		// visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
	g.lock.RUnlock()
}

type WeightGraph struct {
	NumNodes int
	Edges    [][]WEdge
}

type WEdge struct {
	From   int
	To     int
	Weight int
}

// NewGraph: Create graph with n nodes.
func NewWeightGraph(n int) *WeightGraph {
	return &WeightGraph{
		NumNodes: n,
		Edges:    make([][]WEdge, n),
	}
}

// AddEdge: Add an edge from u to v.
func (g *WeightGraph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], WEdge{From: u, To: v, Weight: w})

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], wEdge{From: v, To: u, Weight: w})
}

func (g *WeightGraph) AdjacentEdgesExample() {
	u := 0 // Example node label.

	fmt.Printf("Printing all edges adjacent to %d\n", u)
	for _, e := range g.Edges[u] {
		fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
	}

	fmt.Println("Printing all edges in graph.")
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
		}
	}
}

func (g *WeightGraph) GetMinEdge(s map[int]int) *WEdge {
	//這裡應該先把所有的都排序 然後慢慢吐會比較好做
	min := models.MaxInt
	minEdge := WEdge{}
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			if s[e.From] != -1 && s[e.To] != -1 {
				//等於找過了
				continue
			}
			if min >= e.Weight {
				min = e.Weight
				minEdge = e
			}
			// fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
		}
	}
	// fmt.Printf("min Edge: %d -> %d (%d)\n", minEdge.From, minEdge.To, minEdge.Weight)
	return &minEdge
}

type ByWeight []WEdge

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (g *WeightGraph) GetSortedEdges() []WEdge {
	edges := []WEdge{}

	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			edges = append(edges, e)
		}
	}
	sort.Sort(ByWeight(edges))
	return edges
}
