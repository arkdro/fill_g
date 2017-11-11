package plate

import (
	"github.com/asdf/fill_g/node"
)

type Plate struct {
	Width int
	Height int
	Data [][]int
}

func (plate Plate) Get_color(node node.Node) int {
	x := node.X
	y := node.Y
	color := plate.Data[y][x]
	return color
}

func (plate *Plate) Set_color(node node.Node, color int) {
	x := node.X
	y := node.Y
	plate.Data[y][x] = color
}

