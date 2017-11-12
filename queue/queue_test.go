package queue

import (
	"github.com/asdf/fill_g/node"
	"reflect"
	"testing"
)

func Test_Empty1(t *testing.T) {
	q := New()
	actual := q.Empty()
	expected := true
	if actual != expected {
		t.Error("queue empty error, 1")
	}
}

func Test_Empty2(t *testing.T) {
	q := New()
	q.Push(node.Node{})
	actual := q.Empty()
	expected := false
	if actual != expected {
		t.Error("queue empty error, 2")
	}
}

func Test_Push(t *testing.T) {
	q := New()
	q.Push(node.Node{X: 1, Y: 2})
	q.Push(node.Node{X: 2, Y: 22})
	q.Push(node.Node{X: 3, Y: 33})
	data := make([]node.Node, 3)
	data[0] = node.Node{X: 1, Y: 2}
	data[1] = node.Node{X: 2, Y: 22}
	data[2] = node.Node{X: 3, Y: 33}
	if q.Len() != len(data) {
		t.Error("queue push length mismatch")
	}
	for i, _ := range data {
		if !reflect.DeepEqual(q[i], data[i]) {
			t.Error("queue push mismatch at", i, ", items:", q[i], data[i])
		}
	}
}

func Test_Pop1(t *testing.T) {
	q := New()
	q.Push(node.Node{X: 1, Y: 2})
	q.Push(node.Node{X: 2, Y: 22})
	q.Push(node.Node{X: 3, Y: 33})
	actual_node := q.Pop()
	expected_node := node.Node{X: 1, Y: 2}
	if actual_node != expected_node {
		t.Error("queue pop node error")
	}
	data := make([]node.Node, 2)
	data[0] = node.Node{X: 2, Y: 22}
	data[1] = node.Node{X: 3, Y: 33}
	for i, _ := range data {
		if !reflect.DeepEqual(q[i], data[i]) {
			t.Error("queue pop rest mismatch at", i, ", items:", q[i], data[i])
		}
	}
}

