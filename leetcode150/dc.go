package main

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

Example 1:


Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:


Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==1 {
		return &TreeNode{Val: nums[0], Right: nil, Left: nil}
	}
    var recursive func(nums []int, left, right int) *TreeNode

	recursive = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid:= left + (right-left)/2

		root:= &TreeNode{
			Val: nums[mid], // 0
			Left: recursive(nums, left, mid-1), // r(0, 1)
			Right: recursive(nums, mid+1, right), // r(3, 5)
		}
		return root
	}

	return recursive(nums, 0, len(nums)-1)
}

/*
Given the head of a linked list, return the list after sorting it in ascending order.

Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]
Example 2:

Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
Example 3:

Input: head = []
Output: []
 
Constraints:

The number of nodes in the list is in the range [0, 5 * 104].
-105 <= Node.val <= 105
 

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/
func sortList(head *ListNode) *ListNode {

	var merge func(l1, l2 *ListNode) *ListNode

	merge = func(l1, l2 *ListNode) *ListNode {
		tmp := &ListNode{Val: 0, Next: nil}
		moving:= tmp
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				moving.Next = l1
				l1 = l1.Next
			}else {
				moving.Next = l2
				l2 = l2.Next
			}
			moving = moving.Next
		}

		if l1 == nil {
			moving.Next = l2
		}

		if l2 == nil {
			moving.Next = l1
		}

		return tmp.Next

	}

	var sort func(head *ListNode) *ListNode


	sort = func(head *ListNode) *ListNode{
		if head == nil || head.Next == nil {
			return head
		}

		current:= head
		mid:=head
		var tmp *ListNode
		for current != nil && current.Next != nil{
			tmp = mid
			mid = mid.Next
			current = current.Next.Next
		}
		tmp.Next = nil

		l1:= sort(head)
		l2:= sort(mid)
		return merge(l1,l2)

	}

	return sort(head)
	

}