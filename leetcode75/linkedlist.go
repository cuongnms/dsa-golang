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
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	head.Next = nil
	for current != nil {
		tmp := current.Next
		current.Next = head
		head = current
		current = tmp
	}
	return head
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
	slow, fast := head, head.Next
	rs := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	startRevert := slow.Next
	current := startRevert.Next
	startRevert.Next = nil
	for current != nil {
		tmp := current.Next
		current.Next = startRevert
		startRevert = current
		current = tmp
	}
	for startRevert != nil {
		if rs < startRevert.Val+head.Val {
			rs = startRevert.Val + head.Val
		}
		startRevert = startRevert.Next
		head = head.Next
	}
	return rs
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := addTwoNumbersRecur(l1, l2, 0)

	return head
}

func addTwoNumbersRecur(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0{
		return nil
	}

	val:= carry
	if l1 != nil {
		val += l1.Val
	}
	if l2!= nil {
		val += l2.Val
	}
	node:=&ListNode{Val: val%10, Next: nil}
	var nextL1 *ListNode
	var nextL2 *ListNode
	if l1 == nil {
		nextL1 = nil
	}else {
		nextL1 = l1.Next
	}
	if l2 == nil {
		nextL2 = nil
	}else {
		nextL2 = l2.Next
	}
	node.Next = addTwoNumbersRecur(nextL1, nextL2, val/10)
	return node
}


/*
Given a linked list, swap every two adjacent nodes and return its head. 
You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
 
Example 1:

Input: head = [1,2,3,4]

Output: [2,1,4,3]

Explanation:

Example 2:

Input: head = []

Output: []

Example 3:

Input: head = [1]

Output: [1]

Example 4:

Input: head = [1,2,3]

Output: [2,1,3]

Constraints:

The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	// 1 -> 2 -> 3

	// 1 <- 2 -> 3
	// |______|

	firstNode := head
	thirdNode := head.Next.Next
	head = head.Next
	head.Next = firstNode
	head.Next.Next = swapPairs(thirdNode)	

	return head

}

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100
Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
 

Constraints:

-100.0 < x < 100.0
-231 <= n <= 231-1
n is an integer.
Either x is not zero or n > 0.
-104 <= xn <= 104
*/
func myPow(x float64, n int) float64 {
	if n==1 {
		return x
	}
    if n == -1 {
        return 1/x
    }

    if n < -1 {
		if n%2 == 0 {
			return myPow(x*x, n/2)

		}else {
			return myPow(x*x, n/2)* (1/x)
		}
    } else if n > 1 {
		if n%2 == 0 {
			return myPow(x*x, n/2)

		}else {
			return myPow(x*x, n/2)* x
		}
	}else {
		return 1
	}
}

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
 

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 != nil {
		return list2
	}
	var firstNode *ListNode
	var firstVal int
	if list1.Val < list2.Val {
		firstVal = list1.Val
		firstNode = &ListNode{Val: firstVal, Next: nil}
		firstNode.Next = mergeTwoLists(list1.Next, list2)

	}else {
		firstVal = list2.Val
		firstNode = &ListNode{Val: firstVal, Next: nil}
		firstNode.Next = mergeTwoLists(list1, list2.Next)
	}

	return firstNode

}