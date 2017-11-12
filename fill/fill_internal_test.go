package fill

import (
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"
	"github.com/asdf/fill_g/queue"
	"reflect"
	"testing"
)

func Test_same_color(t *testing.T) {
	plate := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,7},
		},
	}
	node := node.Node{X: 2, Y: 1}
	color := 7
	actual := same_color(node, color, plate)
	expected := true
	if actual != expected {
		t.Error("same color mismatch")
	}
}

func Test_handle_neighbour_node(t *testing.T) {
	plate1 := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,7},
		},
	}
	node := node.Node{X: 2, Y: 1}
	old_color := 7
	color := 3
	queue1 := queue.New()
	handle_neighbour_node(node, &queue1, old_color, color, plate1)
	expected_plate := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,3},
		},
	}
	actual_plate := plate1
	if !reflect.DeepEqual(actual_plate, expected_plate) {
		t.Error("handle neighbour node 1, plate mismatch")
	}
	queue2 := queue.New()
	queue2.Push(node)
	expected_queue := queue2
	actual_queue := queue1
	if actual_queue.Len() != expected_queue.Len() {
		t.Error("handle neighbour node 1, queue size mismatch, ",
			actual_queue.Len(), "!=", expected_queue.Len())
	}
	for i, _ := range actual_queue {
		if !reflect.DeepEqual(expected_queue[i], actual_queue[i]) {
			t.Error("handle neighbour node 1, queue mismatch at", i,
				", items:", expected_queue[i], actual_queue[i])
		}
	}
}

