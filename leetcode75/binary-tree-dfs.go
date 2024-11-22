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

/*
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Example 1:

Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true
Example 2:

Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false

Constraints:

The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].
*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	tree1 := make([]*TreeNode,0)
	tree2 := make([]*TreeNode,0)
	tree1 = append(tree1, root1)
	tree2 = append(tree2, root2)
	for len(tree1) > 0 && len(tree2) > 0 {
		tmpTree1, val1 := getLeaf(tree1)
		tmpTree2, val2 := getLeaf(tree2)
		tree1 = tmpTree1
		tree2 = tmpTree2
		if val1 != val2 {
			return false
		}
	}
	if len(tree1) == 0 && len(tree2) == 0 {
		return true
	}else {
		return false
	}
}

// solution 1: travel all and compare the result
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	return append(left, right...)
}

// solution 2: compare node by node in 2 trees.
func getLeaf(tree []*TreeNode) ([]*TreeNode, int) {
	for {
		node:= tree[len(tree)-1]
		tree = tree[0:len(tree)-1]
		if node.Right != nil {
			tree = append(tree, node.Right)
		}
		if node.Left != nil {
			tree = append(tree, node.Left)
		}
		if node.Left == nil && node.Right == nil {
			return tree, node.Val
		}
	}
}


/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.

Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.
 
Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].
*/
func goodNodes(root *TreeNode) int {
	rs := countRercusive(root, root.Val)
	return rs
}

func countRercusive(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	count :=0
	if max <= root.Val {
		max = root.Val
		count++
	}
	return count + countRercusive(root.Left, max) + countRercusive(root.Right, max)
}


/*
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

Example 1:


Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.
Example 2:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3
 

Constraints:

The number of nodes in the tree is in the range [0, 1000].
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000
*/
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
		return 0
	}
	return pathSumRecursive(root, targetSum - root.Val) + pathSum(root.Right, targetSum) + pathSum(root.Left, targetSum)
	
}

func pathSumRecursive(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	
	if root.Val == targetSum {
		return 1
	}else {
		return pathSumRecursive(root.Left, targetSum - root.Val) + pathSumRecursive(root.Right, targetSum - root.Val)
	}

}