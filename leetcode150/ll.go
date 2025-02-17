package main

import "fmt"

/*
Given the head of a linked list,
rotate the list to the right by k places.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
Example 2:

Input: head = [0,1,2], k = 4
Output: [2,0,1]

Constraints:

The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 109
  
1 2 3 4 5 
k=2

5.Next = 1
k = 2%5 = 2
k = 5-2 = 3



*/
func rotateRight(head *ListNode, k int) *ListNode {
	fmt.Println()
	if head == nil || k == 0 || head.Next == nil{
		return head
	}
	lenList:= 1
	last:= head
	for last.Next != nil {
		lenList++
		last = last.Next
	}
	times:= k % lenList
	if times == 0 {
		return head
	}
	
	last.Next = head
	for i:=0; i < lenList - times; i++ {
		last =last.Next
	}
	head = last.Next
	last.Next = nil 
	return head
}
