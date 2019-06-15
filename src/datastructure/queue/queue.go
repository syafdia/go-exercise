package queue

import (
	"fmt"

	"github.com/syafdia/go-exercise/src/datastructure/types"
)

type Queue interface {
	Enqueue(v types.T)
	Dequeue() types.T
	Size() int
	String() string
}

type queue struct {
	values []types.T
}

func New() Queue {
	return &queue{
		values: []types.T{},
	}
}

func (q *queue) Enqueue(v types.T) {
	q.values = append(q.values, v)
}

func (q *queue) Dequeue() types.T {
	if q.Size() == 0 {
		return nil
	}

	v := q.values[0]
	q.values = q.values[1:]

	return v
}

func (q *queue) Size() int {
	return len(q.values)
}

func (q *queue) String() string {
	return fmt.Sprintf("%v", q.values)
}
