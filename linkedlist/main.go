package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	size int
}

func NewLinkedList[T any](value T) *LinkedList[T] {

	node := &Node[T]{value: value, next: nil}
	return &LinkedList[T]{
		head: node,
		size: 1,
	}
}

func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size++
}

func (l *LinkedList[T]) Display() {
	fmt.Println("------------")
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

	fmt.Println("------------")
}

func (l *LinkedList[T]) Remove() {
	if l.head.next == nil {
		l.head = nil
	} else {
		current := l.head
		for current.next.next != nil {
			current = current.next
		}
		current.next = nil
	}
	l.size--

}

func (l *LinkedList[T]) AddAt(index int, value T) {
	newNode := &Node[T]{value: value, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		count := 0
		current := l.head
		for count < index-1 {
			if current.next != nil {
				current = current.next
				count++
			}
		}
		newNode.next = current.next
		current.next = newNode
	}
	l.size++
}

func (l *LinkedList[T]) RemoveAt(index int) {
	if index >= l.size {
		return
	}

	if index == l.size-1 {
		l.Remove()
		return
	}

	if index == 0 && index < l.size {
		l.head = l.head.next

	} else {
		count := 0
		current := l.head
		for count < index-1 {
			if current.next != nil {
				current = current.next
				count++
			}
		}

		current.next = current.next.next

	}
	l.size--

}

func main() {
	list := NewLinkedList(10)
	list.Add(12)
	list.Add(13)
	// list.Display()
	// list.Remove()
	// list.Display()
	list.AddAt(1, 11)
	// list.Display()
	list.Display()
	list.RemoveAt(4)
	list.Display()

}
