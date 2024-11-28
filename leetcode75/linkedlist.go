package main

import (
	"fmt"
	"math"
)

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

// 1 -> 2 -> 3 -> 4
func reverseListRecursive(head *ListNode) *ListNode {
	tmp:= reverseByTailRecursive(nil, head)
	return tmp
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


/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4]
Output: [1,4,2,3]
Example 2:

Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
 
Constraints:

The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
*/
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
        return
    }
    mid:=head
	fast:=head
	prevMid:=mid
	for fast!= nil && fast.Next != nil {
		prevMid = mid
		mid = mid.Next
		fast = fast.Next.Next
	}
	prevMid.Next = reverseByTailRecursive(nil, mid)
	

	currentHead := head
	currentMid:= prevMid.Next

	for currentHead != prevMid {
		prevMid.Next = currentMid.Next
		currentMid.Next = currentHead.Next
		currentHead.Next = currentMid
		currentHead = currentMid.Next
		currentMid = prevMid.Next
	}
}

/*
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Example 1:
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Example 2:

Input: head = [], val = 1
Output: []
Example 3:

Input: head = [7,7,7,7], val = 7
Output: []
 
Constraints:
The number of nodes in the list is in the range [0, 104].
1 <= Node.val <= 50
0 <= val <= 50
*/
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
		return nil
	}

	head.Next = removeElements(head.Next, val)
	if head.Val ==val {
		return head.Next
	}else {
		return head
	}

}


/*
Given the head of a singly linked list, return true if it is a 
palindrome or false otherwise.

Example 1:

Input: head = [1,2,2,1]
Output: true
Example 2:

Input: head = [1,2]
Output: false
 
Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
 
Follow up: Could you do it in O(n) time and O(1) space?
*/
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow= slow.Next
		fast = fast.Next.Next
	}

	mid:= reverseByTailRecursive(nil, slow)
	for head != slow {
		if head.Val != mid.Val {
			return false
		}
		head = head.Next
		mid = mid.Next
	}
	return true

}

func reverseByHeadRecursive(prev *ListNode, current *ListNode) *ListNode {
	// this head recursive method will return the old head (first node in the old list)
	if current == nil {
		return prev
	}
	// if current.Next != nil {
	// 	fmt.Println("before recursive: currentVal = ", current.Val, " next: ", current.Next.Val)
	// }else {
	// 	fmt.Println("before recursive: currentVal = ", current.Val, " next: nil")
	// }

	// fmt.Println("----------------------")

	current = reverseByHeadRecursive(current, current.Next) // current = re(4, nil) => 4
	// if current.Next != nil {
	// 	fmt.Println("after recursive: currentVal = ", current.Val, " next: ", current.Next.Val)
	// }else {
	// 	fmt.Println("after recursive: currentVal = ", current.Val, " next: nil")
	// }
	current.Next = prev
	// if current.Next != nil {
	// 	fmt.Println("after redirect: currentVal = ", current.Val, " next: ", current.Next.Val)
	// }else {
	// 	fmt.Println("after redirect: currentVal = ", current.Val, " next: nil")
	// }

	// fmt.Println("----------------------")

	return prev
}

func reverseByTailRecursive(prev *ListNode, current *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	// if current.Next != nil {
	// 	fmt.Println("before redirect: currentVal = ", current.Val, " next: ", current.Next.Val)
	// }else {
	// 	fmt.Println("before redirect: currentVal = ", current.Val, " next: nil")
	// }
	// fmt.Println("======================")

	

	nextNode:=current.Next
	current.Next = prev
	
	// if current.Next != nil {
	// 	fmt.Println("after redirect: currentVal = ", current.Val, " next: ", current.Next.Val)
	// }else {
	// 	fmt.Println("before redirect: currentVal = ", current.Val, " next: nil")
	// }

	current = reverseByTailRecursive(current, nextNode)
	return current
}


/*
You have a list arr of all integers in the range [1, n] sorted in a strictly increasing order. Apply the following algorithm on arr:

Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
Repeat the previous step again, but this time from right to left, remove the rightmost number and every other number from the remaining numbers.
Keep repeating the steps again, alternating left to right and right to left, until a single number remains.
Given the integer n, return the last number that remains in arr.

Example 1:

Input: n = 9
Output: 6
Explanation:
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9] 
arr = [2, 4, 6, 8] = 2[1,2,3,4]
arr = [2, 6] = 2[1,3]
arr = [6] = 6[1]
Example 2:

Input: n = 1
Output: 1
 

Constraints:

1 <= n <= 109
*/

func lastRemaining(n int) int {
	return eliminate(9, true)
}

func eliminate(n int, isLeftToRight bool) int {
	if n == 1 {
		return 1
	}

	if isLeftToRight {
        return 2 * eliminate(n/2, false)
    } else {
		if n%2 == 0 {
			return 2 * eliminate(n/2, true) - 1 
		}else {
			return 2 * eliminate(n/2, true)
		}
    }
}

/*
You are given an integer array nums. Two players are playing a game with this array: player 1 and player 2.
Player 1 and player 2 take turns, with player 1 starting first. Both players start the game with a score of 0. 
At each turn, the player takes one of the numbers from either end of the array (i.e., nums[0] or nums[nums.length - 1]) which reduces the size of the array by 1. 
The player adds the chosen number to their score. The game ends when there are no more elements in the array.
Return true if Player 1 can win the game. If the scores of both players are equal, then player 1 is still the winner, and you should also return true. 
You may assume that both players are playing optimally.

Example 1:

Input: nums = [1,5,2]
Output: false
Explanation: Initially, player 1 can choose between 1 and 2. 
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5, then player 1 will be left with 1 (or 2). 
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5. 
Hence, player 1 will never be the winner and you need to return false.
Example 2:

Input: nums = [1,5,233,7]
Output: true
Explanation: Player 1 first chooses 1. Then player 2 has to choose between 5 and 7. No matter which number player 2 choose, player 1 can choose 233.
Finally, player 1 has more score (234) than player 2 (12), so you need to return True representing player1 can win.
 
Constraints:

1 <= nums.length <= 20
0 <= nums[i] <= 107


*/
func predictTheWinner(nums []int) bool {
	sum:=0
	for _, val:= range nums {
		sum+=val
	}

	p1 := predictRecursive(nums, 0, len(nums)-1)
	if sum - p1 > p1 {
		return false
	}else {
		return true
	}
    
}


func predictRecursive(nums []int, left, right int) int {
	if left > right {
		return 0
	}
	
	sum1:= nums[left] + min(predictRecursive(nums, left+2, right), predictRecursive(nums, left+1, right -1 ))
	sum2:= nums[right] + min(predictRecursive(nums, left, right-2), predictRecursive(nums, left+1, right - 1))
	return max(sum1, sum2)


}

func min(a,b int) int {
	if a>=b {
		return b
	}

	return a
}

func max(a,b int) int {
	if a<b {
		return b
	}else {
		return a
	}
}