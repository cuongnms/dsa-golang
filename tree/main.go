package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Binary tree
type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	root *Node[T]
}

func (b *BinaryTree[T]) insert(value T) {
	b.insertRec(b.root, value)
}

func (b *BinaryTree[T]) insertRec(node *Node[T], value T) *Node[T] {
	if b.root == nil {
		b.root = &Node[T]{value: value, right: nil, left: nil}
		return b.root
	}

	if node == nil {
		node = &Node[T]{value: value, right: nil, left: nil}
		return node
	}

	if node.value > value {
		node.left = b.insertRec(node.left, value)
	} else {
		node.right = b.insertRec(node.right, value)
	}
	return node
}

func (b *BinaryTree[T]) travelLeftToRootToRight(node *Node[T]) {
	if b == nil {
		return
	}
	if node.right == nil && node.left == nil {
		fmt.Println(node.value)
		return
	} else {
		b.travelLeftToRootToRight(node.left)
		fmt.Println(node.value)
		b.travelLeftToRootToRight(node.right)
	}
}

func (b *BinaryTree[T]) travelLeftToRightToRoot(node *Node[T]) {
	if b == nil {
		return
	}
	if node.right == nil && node.left == nil {
		fmt.Println(node.value)
		return
	} else {
		b.travelLeftToRightToRoot(node.left)
		b.travelLeftToRightToRoot(node.right)
		fmt.Println(node.value)
	}
}

func (b *BinaryTree[T]) travelRootToLeftToRight(node *Node[T]) {
	if b == nil {
		return
	}
	if node.right == nil && node.left == nil {
		fmt.Println(node.value)
		return
	} else {
		fmt.Println(node.value)
		b.travelRootToLeftToRight(node.left)
		b.travelRootToLeftToRight(node.right)
	}
}

func (b *BinaryTree[T]) travelByLevel() {
	if b == nil {
		return
	}
	queue := []*Node[T]{b.root}

	for len(queue) > 0 {
		last := queue[len(queue)-1]
		fmt.Println(last.value)
		queue = queue[:len(queue)-1]
		if last.left != nil {
			queue = append([]*Node[T]{last.left}, queue...)
		}
		if last.right != nil {
			queue = append([]*Node[T]{last.right}, queue...)
		}
	}
}

func (b *BinaryTree[T]) find(node *Node[T], value T) (*Node[T], error) {
	if b == nil {
		var zeroVal *Node[T]
		return zeroVal, fmt.Errorf("empty tree")
	}

	if node == nil {
		var zeroVal *Node[T]
		fmt.Println("Not found!!!")
		return zeroVal, fmt.Errorf("not found")
	}

	if node.value == value {
		fmt.Println("Found!!!")
		return node, nil
	}

	if node.value > value {
		return b.find(node.left, value)
	} else {
		return b.find(node.right, value)
	}

}

func (b *BinaryTree[T]) remove(node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if node.value > value {
		node.left = b.remove(node.left, value)
		return node
	} else if node.value < value {
		node.right = b.remove(node.right, value)
		return node
	} else {
		if node.left == nil && node.right == nil {
			node = nil
			return node
		} else if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			minNode, _ := b.minNode(node)
			node.value = minNode.value
			node.right = b.remove(minNode, minNode.value)
		}
		return node
	}

}

func (b *BinaryTree[T]) minNode(node *Node[T]) (*Node[T], error) {
	if node.right == nil {
		return node.left, nil
	}

	node = node.right
	for node.left != nil {
		node = node.left
	}
	return node, nil
}

func main() {
	tree := &BinaryTree[int]{}
	tree.insert(10)
	tree.insert(8)
	tree.insert(12)
	tree.insert(6)
	tree.insert(9)

	// tree.travelLeftToRootToRight(tree.root)
	// tree.travelLeftToRightToRoot(tree.root)
	// tree.travelRootToLeftToRight(tree.root)
	// tree.travelByLevel()
	// tree.find(tree.root, 9)
	tree.remove(tree.root, 10)
	fmt.Println(tree.root.value)           // 10
	fmt.Println(tree.root.left.value)      // 8
	fmt.Println(tree.root.right)     // 12
	fmt.Println(tree.root.left.right.value)      // 9
	fmt.Println(tree.root.left.left.value) // 9
}
