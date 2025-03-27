package main

import (
	"fmt"
	"math"
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

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true
Example 2:

Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
 
Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
*/
func isValidBST(root *TreeNode) bool {
    var find func(node *TreeNode, min, max int) bool
	find = func(node *TreeNode, min,max int) bool {
		if node == nil {
			return true
		}

		if node.Val <= min || node.Val >= max {
			return false
		}

		return find(node.Left, min, node.Val) && find(node.Right, node.Val, max)
	} 

	return find(root, math.MinInt64, math.MaxInt64)
}