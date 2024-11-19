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

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:

Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []

Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000
 
Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/
func reverseList(head *ListNode) *ListNode {
    if head== nil || head.Next == nil {
		return head
	}
	current:= head.Next
	head.Next = nil
	for current.Next != nil {
		tmp:=current.Next
		current.Next = head
		head = current
		current = tmp
	}
	current.Next = head
	return current

}

/*
In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.
For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.
Given the head of a linked list with even length, return the maximum twin sum of the linked list.

Example 1:

Input: head = [5,4,2,1]
Output: 6
Explanation:
Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
There are no other nodes with twins in the linked list.
Thus, the maximum twin sum of the linked list is 6. 
Example 2:

Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7. 
Example 3:

Input: head = [1,100000]
Output: 100001
Explanation:
There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.

Constraints:

The number of nodes in the list is an even integer in the range [2, 105].
1 <= Node.val <= 105
*/
func pairSum(head *ListNode) int {
    
}