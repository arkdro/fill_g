package queue

// fifo queue for nodes

import (
	"github.com/asdf/fill_g/node"
)

type Queue []node.Node

func New() Queue {
	return make([]node.Node, 0)
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func (q *Queue) Push(x node.Node) {
	*q = append(*q, x)
}

func (q *Queue) Pop() node.Node {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q Queue) Len() int {
	return len(q)
}

