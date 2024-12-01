package main

import "fmt"

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
	rs:=make([]int,0)
	
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

