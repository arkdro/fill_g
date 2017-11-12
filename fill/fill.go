package fill

import (
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"
	"github.com/asdf/fill_g/queue"
)

func Run(start_point node.Node, color int, plate plate.Plate) plate.Plate {
	if already_filled(start_point, color, plate) {
		return plate
	}
	queue := queue.New()
	old_color := plate.Get_color(start_point)
	plate.Set_color(start_point, color)
	queue.Push(start_point)
	loop(queue, old_color, color, plate)
	return plate
}

func loop(queue queue.Queue, old_color int, color int, plate plate.Plate) {
	for {
		if queue.Empty() {
			return
		}
		n := queue.Pop()
		left, ok := plate.Get_left_node(n)
		if ok {
			handle_neighbour_node(left, queue, old_color, color, plate)
		}
		right, ok := plate.Get_right_node(n)
		if ok {
			handle_neighbour_node(right, queue, old_color, color, plate)
		}
		up, ok := plate.Get_up_node(n)
		if ok {
			handle_neighbour_node(up, queue, old_color, color, plate)
		}
		down, ok := plate.Get_down_node(n)
		if ok {
			handle_neighbour_node(down, queue, old_color, color, plate)
		}
	}
}

func handle_neighbour_node(node node.Node, queue queue.Queue, old_color int, color int, plate plate.Plate) {
	if node_fits(node, old_color, plate) {
		plate.Set_color(node, color)
		queue.Push(node)
	}
}

func node_fits(node node.Node, old_color int, plate plate.Plate) bool {
	return same_color(node, old_color, plate)
}

func already_filled(start_point node.Node, color int, plate plate.Plate) bool {
	return same_color(start_point, color, plate)
}

func same_color(node node.Node, color int, plate plate.Plate) bool {
	existing_color := plate.Get_color(node)
	return color == existing_color
}

