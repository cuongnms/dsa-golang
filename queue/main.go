package main

import (
	"fmt"

)

type QueueBaseSlice[T any] struct {
	store []T
}

func NewQueueBaseSlice[T any]() *QueueBaseSlice[T] {
	return &QueueBaseSlice[T]{}
}

func (q *QueueBaseSlice[T]) dequeue() (T, error){
	if len(q.store) == 0 {
		var zeroVal T
		return zeroVal, fmt.Errorf("queue empty")
	}
	rs := q.store[0]
	q.store = q.store[1:len(q.store)]
	return rs, nil
}

func (q *QueueBaseSlice[T]) enqueue(value T) {
	q.store = append(q.store, value)
}

func (q *QueueBaseSlice[T]) peek() (T, error) {
	if len(q.store) == 0 {
		var zeroVal T
		return zeroVal, fmt.Errorf("queue empty")
	}
	rs := q.store[0]
	return rs, nil
}

func (q *QueueBaseSlice[T]) Print() {
	fmt.Println("---------------------")
	for index := range q.store {
		fmt.Println(q.store[index])
	}

	fmt.Println("---------------------")
}

//---------------------

type Node[T any] struct {
	value T
	next *Node[T]
	prev *Node[T]
}

type QueueBaseList[T any] struct{
	head *Node[T]
	tail *Node[T]
}

func NewQueueBaseList[T any]() *QueueBaseList[T]{
	return &QueueBaseList[T]{}
}

func (q *QueueBaseList[T]) Print() {
	fmt.Println("---------------------")
	current := q.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("---------------------")
}

func (q *QueueBaseList[T]) enqueue(value T) {
	node:= &Node[T]{value: value, next: nil, prev: nil}
	if q.head == nil || q.tail == nil{
		q.head = node
		q.tail = node
		return
	}
	node.next = q.head
	q.head.prev = node
	q.head = node
}

func (q *QueueBaseList[T]) dequeue() (T, error) {
	if q.head == nil || q.tail == nil {
		var zeroVal T
		return zeroVal, fmt.Errorf("queue empty")
	}

	if q.head == q.tail {
		rs := q.head.value
		q.head = nil
		q.tail = nil
		return rs, nil
	}

	rs:=q.tail.value
	q.tail = q.tail.prev
	q.tail.next = nil
	return rs, nil
}

func (q *QueueBaseList[T]) peek() (T, error) {
	if q.head == nil || q.tail == nil {
		var zeroVal T
		return zeroVal, fmt.Errorf("queue empty")
	}

	if q.head.next == q.tail {
		rs := q.head.value
		return rs, nil
	}

	rs:=q.tail.value
	return rs, nil
}

func main() {
	// queueBaseSlice := NewQueueBaseSlice[int]()
	// queueBaseSlice.enqueue(10)
	// queueBaseSlice.enqueue(11)
	// queueBaseSlice.enqueue(12)
	// queueBaseSlice.Print()
	// queueBaseSlice.dequeue()
	// queueBaseSlice.Print()

	queueBaseList := NewQueueBaseList[int]()
	queueBaseList.enqueue(10)
	queueBaseList.enqueue(11)
	queueBaseList.dequeue()
	queueBaseList.dequeue()
	queueBaseList.dequeue()
	queueBaseList.enqueue(10)
	queueBaseList.Print()
}