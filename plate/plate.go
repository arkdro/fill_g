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

func (plate Plate) Get_left_node(orig node.Node) (node.Node, bool) {
	if orig.X > 0 {
		n := node.Node{
			X: orig.X - 1,
			Y: orig.Y,
		}
		return n, true
	}
	return node.Node{}, false
}

func (plate Plate) Get_right_node(orig node.Node) (node.Node, bool) {
	width := plate.Width
	if orig.X < width - 1 {
		n := node.Node{
			X: orig.X + 1,
			Y: orig.Y,
		}
		return n, true
	}
	return node.Node{}, false
}

func (plate Plate) Get_down_node(orig node.Node) (node.Node, bool) {
	if orig.Y > 0 {
		n := node.Node{
			X: orig.X,
			Y: orig.Y - 1,
		}
		return n, true
	}
	return node.Node{}, false
}

func (plate Plate) Get_up_node(orig node.Node) (node.Node, bool) {
	height := plate.Height
	if orig.Y < height - 1 {
		n := node.Node{
			X: orig.X,
			Y: orig.Y + 1,
		}
		return n, true
	}
	return node.Node{}, false
}

func (plt Plate) Valid_data() bool {
	if plt.Width <= 0 {
		return false
	}
	if plt.Height <= 0 {
		return false
	}
	if len(plt.Data) != plt.Height {
		return false
	}
	for _, row := range plt.Data {
		if len(row) != plt.Width {
			return false
		}
	}
	return true
}

