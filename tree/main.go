package main

import "golang.org/x/exp/constraints"

// Binary tree
type Node[T constraints.Ordered] struct {
	value T
	left *Node[T]
	right *Node[T]
}


func (t *Node[T]) insert(value T) {
	node := &Node[T]{value: value, left: nil, right: nil}
	if t == nil {
		t = node
		return
	}

	if t.value > value {
		if t.left == nil {
			t.left = node
			return
		}else {
			t.left.insert(value)
		}
	} else {
		if t.right == nil {
			t.right = node
			return
		}else {
			t.right.insert(value)
		}
	}

}

func (t *Node[T]) print() {
	
}
