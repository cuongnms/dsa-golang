package main

import (
	"fmt"
)

/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
/*

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Print() {
	fmt.Println(node.Val)
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func maxDepthIter(root *TreeNode) int {
	rs := 0

	nodes := make([]*TreeNode, 0)
	depths := make([]int, 0)

	nodes = append(nodes, root)
	depths = append(depths, 1)

	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[0 : len(nodes)-1]
		depth := depths[len(depths)-1]
		depths = depths[0 : len(depths)-1]

		if rs < depth {
			rs = depth
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			depths = append(depths, depth+1)
		}

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			depths = append(depths, depth+1)
		}

	}

	return rs
}
