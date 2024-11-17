package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ll *ListNode) Print() {
	current:= ll
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func deleteMiddle(head *ListNode) *ListNode {
	// count:=0
	// current:= head
	// for current != nil {
	// 	current = current.Next
	// 	count++
	// }
	// fmt.Println("count: ", count)
	// fmt.Println("mid: ", count/2 - 1)
	// check:=0
	// current=head
	// for check< count/2 - 1 {
	// 	current=current.Next
	// 	check++
	// }
	// current.Next= current.Next.Next
	// return head
	if head.Next == nil {
		return nil
	}

	first:=head
	second:=head.Next.Next
	for second != nil && second.Next != nil {
		second = second.Next.Next
		first = first.Next
	}
	first.Next = first.Next.Next
	return head
}
