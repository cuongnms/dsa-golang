package main

import (
	"fmt"
)

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104

Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently,
how would you optimize?
*/
func kthSmallest(root *TreeNode, k int) int {
	fmt.Println()
	var travel func(root *TreeNode) 
	index:=0
	rs:= -1
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		index++
		if index == k {
			rs = root.Val
			return
		}
		travel(root.Right)
	}
	travel(root)
	return rs
}