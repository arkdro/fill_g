package queue

// fifo queue for int numbers

type Queue []int

func New() Queue {
	return make([]int, 0)
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

