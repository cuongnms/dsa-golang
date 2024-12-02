package main

import (
	"fmt"
)

/*
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.



Example 1:

Input: root = [1,2,3,null,5,null,4]

Output: [1,3,4]

Explanation:



Example 2:

Input: root = [1,2,3,4,null,null,null,5]

Output: [1,3,4,5]

Explanation:



Example 3:

Input: root = [1,null,3]

Output: [1,3]

Example 4:

Input: root = []

Output: []

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

func rightSideView(root *TreeNode) []int {
	rs := make([]int, 0)
	if root == nil {
		return rs
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	rs = append(rs, root.Val)
	for len(q) > 0 {
		s := make([]*TreeNode, 0)
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				s = append(s, q[i].Left)
			}
			if q[i].Right != nil {
				s = append(s, q[i].Right)
			}
		}
		q = make([]*TreeNode, 0)
		q = append(q, s...)
		if len(s) > 0 {
			rs = append(rs, s[len(s)-1].Val)
		} else {
			return rs
		}
	}
	return rs

}

func rightSideViewRecursive(root *TreeNode) []int {
	rs := make([]int, 0)

	var dfs func(root *TreeNode, depth int)

	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(rs) {
			rs = append(rs, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)

	}

	dfs(root, 0)
	fmt.Println(rs)
	return rs
}

/*
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

Example 1:

Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.
Example 2:

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2

Constraints:

The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105
*/

func maxLevelSum(root *TreeNode) int {
	rs := 1
	q := make([]*TreeNode, 0)
	q = append(q, root)
	maxSum := root.Val
	depth := 0
	for len(q) > 0 {
		sum := 0
		for i := len(q); i > 0; i-- {
			node := q[0]
			q = q[1:]
			sum+=node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
		// if sum > max
		if sum > maxSum {
			maxSum = sum
			rs = depth
		}
	}
	return rs

	// sumArr := make([]int, 0)
	// var dfs func(root *TreeNode, depth int)

	// dfs = func(root *TreeNode, depth int) {

	// 	if root == nil {
	// 		return
	// 	}

	// 	if len(sumArr) > depth {
	// 		sumArr[depth] +=root.Val
	// 	}else {
	// 		sumArr = append(sumArr, root.Val)
	// 	}

	// 	dfs(root.Left, depth+1)
	// 	dfs(root.Right, depth + 1)

	// }

	// dfs(root, 0)
	// max:= sumArr[0]
	// rs:=0
	// for i, val:=range sumArr {
	// 	if max < val {
	// 		max = val
	// 		rs = i
	// 	}
	// }
	// return rs + 1
}
