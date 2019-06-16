package linkedlist

import (
	"fmt"

	"github.com/syafdia/go-exercise/src/types"
)

type LinkedList interface {
	Head() types.T
	Tail() LinkedList
	Map(mapper func(t types.T) types.U) LinkedList
	Filter(filterer func(t types.T) bool) LinkedList
	Reduce(initial types.U, reducer func(acc types.U, v types.T) types.U) types.U
	String() string
	Size() int
}

type node struct {
	head types.T
	tail LinkedList
}

func New(values ...types.T) LinkedList {
	totVals := len(values)

	if totVals == 0 {
		return node{}
	}

	head := values[0]

	if totVals == 1 {
		return node{head: head, tail: nil}
	}

	return node{head: head, tail: New(values[1:]...)}
}

func (n node) Head() types.T {
	return n.head
}

func (n node) Tail() LinkedList {
	return n.tail
}

func (n node) Map(mapper func(t types.T) types.U) LinkedList {
	if n.tail == nil {
		return &node{
			head: mapper(n.head),
			tail: nil,
		}
	}

	return &node{
		head: mapper(n.head),
		tail: n.tail.Map(mapper),
	}
}

func (n node) Filter(filterer func(t types.T) bool) LinkedList {
	if filterer(n.head) {
		if n.tail == nil {
			return &node{n.head, nil}
		}

		return &node{n.head, n.tail.Filter(filterer)}
	}

	if n.tail == nil {
		return nil
	}

	return n.tail.Filter(filterer)
}

func (n node) Reduce(initial types.U, reducer func(acc types.U, v types.T) types.U) types.U {
	if n.tail == nil {
		return reducer(initial, n.head)
	}

	return n.tail.Reduce(reducer(initial, n.head), reducer)
}

func (n node) String() string {
	return fmt.Sprintf("(%v, %v)", n.head, n.tail)
}

func (n node) Size() int {
	return n.Reduce(0, func(acc types.U, _ types.T) types.U {
		return acc.(int) + 1
	}).(int)
}
