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

