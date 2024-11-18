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
	current := ll
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

	first := head
	second := head.Next.Next
	for second != nil && second.Next != nil {
		second = second.Next.Next
		first = first.Next
	}
	first.Next = first.Next.Next
	return head
}

/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.
You must solve the problem in O(1) extra space complexity and O(n) time complexity.

Example 1:

Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
Example 2:

Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
2--> 3 --> 6 --> 7 -->

0 1 2 3 4 5 6
e o e o e o e

Constraints:

The number of nodes in the linked list is in the range [0, 104].
-106 <= Node.val <= 106
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	evenNode := head
	oddNode := head.Next
	
	second := head.Next
	for evenNode.Next != nil && oddNode.Next != nil {
		evenNode.Next = oddNode.Next
		evenNode = evenNode.Next
		oddNode.Next = evenNode.Next
		oddNode = oddNode.Next
	}
	evenNode.Next = second
	return head
}
