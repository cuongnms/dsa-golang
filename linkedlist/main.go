package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any](value T) *LinkedList[T] {

	node := &Node[T]{value: value, next: nil}
	return &LinkedList[T]{
		head: node,
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
	if(l.head.next == nil){
		l.head = nil
	}else{
		current := l.head
		for current.next.next != nil {
			current = current.next
		}
		current.next = nil
	}
	
}

func (l *LinkedList[T]) AddAt(index int, value T) {
	newNode := &Node[T]{value: value, next: nil}
	if l.head == nil {
		l.head = newNode
	}else{
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
}

func (l *LinkedList[T]) RemoveAt(index int) {
	if l.head.next == nil && index != 0 {
		l.head = nil
	}else {
		count:=0
		current := l.head
		remove := l.head.next
		for count < index - 1 {
			if current.next != nil {
				current = current.next
				remove = remove.next
				count++
			}
		}
		if count < index - 1 {
			return
		}else {
			current.next = remove.next
		}
	}
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
	list.Remove()
	list.Remove()
	list.Remove()
	list.Remove()
	list.Display()
	
}
