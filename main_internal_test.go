package main

import (
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"
	"reflect"
	"testing"
)

func Test_run_request(t *testing.T) {
	plate1 := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,3,3},
			{5,6,7},
		},
	}
	expected_plate := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{5,5,5},
			{5,6,5},
		},
	}
	request := Request{
		Steps: []Step{
			{Point: node.Node{X: 2, Y: 1}, Color: 3},
			{Point: node.Node{X: 1, Y: 0}, Color: 1},
			{Point: node.Node{X: 0, Y: 0}, Color: 5},
		},
		Input_data: plate1,
		Expected_data: expected_plate,
	}
	actual_plate := run_request(request)
	if !reflect.DeepEqual(actual_plate, expected_plate) {
		t.Error("man run request, plate mismatch,\nactual:", actual_plate,
			"\nexpected:", expected_plate)
	}
}

