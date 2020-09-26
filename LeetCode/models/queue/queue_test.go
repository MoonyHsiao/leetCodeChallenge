package queue

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNew(t *testing.T) {

	q := &Queue{nodes: make([]*Node, 1)}
	q.Push(&Node{1})
	q.Push(&Node{2})
	q.Push(&Node{3})
	q.Push(&Node{4})
	q.Push(&Node{5})

	// for q.GetLen() != 0 {
	// 	fmt.Printf("q:%v\n", q.Pop().Value)
	// }

	// fmt.Printf("pop q:%v\n", q.Pop().Value)

	res := []int{}
	expected := []int{1, 2, 3, 4, 5}
	for q.GetLen() != 0 {
		// fmt.Printf("q:%v\n", q.Pop().Value)
		res = append(res, q.Pop().Value)
	}

	assert.Equal(t, expected, res)
}
