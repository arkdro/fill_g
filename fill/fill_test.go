package fill

import (
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"
	"reflect"
	"testing"
)

func Test_Run1(t *testing.T) {
	plate1 := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,7,7},
			{5,6,7},
		},
	}
	start_point := node.Node{X: 2, Y: 1}
	color := 7
	actual_plate := Run(start_point, color, plate1)
	expected_plate := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{1,7,7},
			{5,6,7},
		},
	}
	if !reflect.DeepEqual(actual_plate, expected_plate) {
		t.Error("fill run 1, plate mismatch")
	}
}

func Test_Run2(t *testing.T) {
	plate1 := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{7,7,7},
			{7,6,7},
		},
	}
	start_point := node.Node{X: 2, Y: 1}
	color := 3
	actual_plate := Run(start_point, color, plate1)
	expected_plate := plate.Plate{
		Width: 3,
		Height: 2,
		Data: [][]int{
			{3,3,3},
			{3,6,3},
		},
	}
	if !reflect.DeepEqual(actual_plate, expected_plate) {
		t.Error("fill run 2, plate mismatch")
	}
}

