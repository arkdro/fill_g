package plate

import (
	"github.com/asdf/fill_g/node"
	"reflect"
	"testing"
)

func Test_Get_color(t *testing.T) {
	plate := Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,7},
		},
	}
	node := node.Node{
		X: 2,
		Y: 1,
	}
	actual := plate.Get_color(node)
	expected := 7
	if !reflect.DeepEqual(actual, expected) {
		t.Error("color mismatch")
	}
}

func Test_Set_color(t *testing.T) {
	plate := Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,7},
		},
	}
	node := node.Node{
		X: 1,
		Y: 1,
	}
	color := 22
	plate.Set_color(node, color)
	expected := Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,22,7},
		},
	}
	if !reflect.DeepEqual(plate, expected) {
		t.Error("set color error")
	}
}

func Test_Get_left_node1(t *testing.T) {
	n := node.Node{X: 2, Y: 1}
	plate := Plate{}
	actual, ok := plate.Get_left_node(n)
	expected := node.Node{X: 1, Y: 1}
	if !ok {
		t.Error("get left node 1, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get left node 1, mismatch")
	}
}

func Test_Get_left_node2(t *testing.T) {
	n := node.Node{X: 0, Y: 1}
	plate := Plate{}
	actual, ok := plate.Get_left_node(n)
	expected := node.Node{}
	if ok {
		t.Error("get left node 2, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get left node 2, mismatch")
	}
}

func Test_Get_right_node1(t *testing.T) {
	n := node.Node{X: 0, Y: 1}
	plate := Plate{Width: 3, Height: 1}
	actual, ok := plate.Get_right_node(n)
	expected := node.Node{X: 1, Y: 1}
	if !ok {
		t.Error("get right node 1, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get right node 1, mismatch")
	}
}

func Test_Get_right_node2(t *testing.T) {
	n := node.Node{X: 2, Y: 1}
	plate := Plate{Width: 3, Height: 1}
	actual, ok := plate.Get_right_node(n)
	expected := node.Node{}
	if ok {
		t.Error("get right node 2, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get right node 2, mismatch")
	}
}

func Test_Get_up_node1(t *testing.T) {
	n := node.Node{X: 0, Y: 1}
	plate := Plate{Width: 2, Height: 3}
	actual, ok := plate.Get_up_node(n)
	expected := node.Node{X: 0, Y: 2}
	if !ok {
		t.Error("get up node 1, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get up node 1, mismatch")
	}
}

func Test_Get_up_node2(t *testing.T) {
	n := node.Node{X: 0, Y: 2}
	plate := Plate{Width: 2, Height: 3}
	actual, ok := plate.Get_up_node(n)
	expected := node.Node{}
	if ok {
		t.Error("get up node 2, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get up node 2, mismatch")
	}
}

func Test_Get_down_node1(t *testing.T) {
	n := node.Node{X: 0, Y: 1}
	plate := Plate{}
	actual, ok := plate.Get_down_node(n)
	expected := node.Node{X: 0, Y: 0}
	if !ok {
		t.Error("get down node 1, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get down node 1, mismatch")
	}
}

func Test_Get_down_node2(t *testing.T) {
	n := node.Node{X: 1, Y: 0}
	plate := Plate{}
	actual, ok := plate.Get_down_node(n)
	expected := node.Node{}
	if ok {
		t.Error("get down node 2, status error")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("get down node 2, mismatch")
	}
}

func Test_Valid_data1(t *testing.T) {
	plate := Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2,3},
			{5,6,7},
		},
	}
	actual := plate.Valid_data()
	expected := true
	if actual != expected {
		t.Error("valid data mismatch, 1")
	}
}

func Test_Valid_data2(t *testing.T) {
	plate := Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,2},
			{5,6,7},
		},
	}
	actual := plate.Valid_data()
	expected := false
	if actual != expected {
		t.Error("valid data mismatch, 2")
	}
}

