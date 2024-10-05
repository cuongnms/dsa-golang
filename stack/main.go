package main

import "fmt"


type StackBaseArr[T any] struct {
	store []T
}

func NewStackBaseArr[T any]() *StackBaseArr[T] {
	return &StackBaseArr[T]{}
} 

func (s *StackBaseArr[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *StackBaseArr[T]) Pop() (T,error){
	if len(s.store)==0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("empty stack")
	}
	rs := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return rs, nil
}

func (s *StackBaseArr[T]) Peek() (T, error){
	if len(s.store)==0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("empty stack")
	}
	rs := s.store[len(s.store)-1]
	return rs, nil
}

func (s *StackBaseArr[T]) Print(){
	fmt.Println("------------------")
	for i:= range s.store {
		fmt.Println(s.store[i])
	}
	fmt.Println("------------------")

}

//--------------------------------------

type Node[T any] struct{
	value T
	next *Node[T]
}
type StackBaseList[T any] struct {
	head *Node[T]
}

func NewStackBaseList[T any]() *StackBaseList[T] {
	return &StackBaseList[T]{}
}

func (s *StackBaseList[T]) Pop() (T, error){
	if s.head == nil {
		var zeroValue T
		return zeroValue, fmt.Errorf("stack empty")
	}
	rs := s.head.value
	s.head = s.head.next
	return rs, nil
}

func (s *StackBaseList[T]) Push(value T) {
	node := &Node[T]{value: value, next: nil}
	if s.head == nil {
		s.head = node
	}else {
		node.next = s.head
		s.head = node
	}
}

func (s *StackBaseList[T]) Peek() (T, error) {
	if s.head == nil {
		var zeroValue T
		return zeroValue, fmt.Errorf("stack empty")
	}
	rs := s.head.value
	return rs, nil
}

func (s *StackBaseList[T]) Print() {
	fmt.Println("-----------------")
	current := s.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("------------------")
}

func main(){
	// s := NewStackBaseArr[int]()
	// s.Push(11)
	// s.Push(12)
	// s.Print()
	// val, err := s.Peek()
	// s.Print()
	// if err == nil{
	// 	fmt.Println(val)
	// }

	s := NewStackBaseList[int]()
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Print()
	val, err := s.Pop()
	if err == nil {
		fmt.Println(val)
	}

	s.Print()
}