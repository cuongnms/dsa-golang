package main

import "fmt"

/*
You are given the root of a binary search tree (BST) and an integer val.
Find the node in the BST that the node's value equals val and return the subtree rooted with that node.
If such a node does not exist, return null.

Example 1:

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
Example 2:


Input: root = [4,2,7,1,3], val = 5
Output: []

Constraints:

The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 107
root is a binary search tree.
1 <= val <= 107
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	fmt.Println("--------------")
	// brute force
    // if root == nil {
	// 	return nil
	// }

	// if root.Val == val {
	// 	return root
	// }

	// rs := searchBST(root.Left, val)
	// if rs == nil {
	// 	return searchBST(root.Right, val)
	// }else {
	// 	return rs
	// }

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}else {
		return searchBST(root.Right, val)
	}
}

/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. 
Return the root node reference (possibly updated) of the BST.
Basically, the deletion can be divided into two stages:
Search for a node to remove.
If the node is found, delete the node.
 
Example 1:

Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:

Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.
Example 3:

Input: root = [], key = 0
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 104].
-105 <= Node.val <= 105
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105
 
Follow up: Could you solve it with time complexity O(height of tree)?
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return nil
	}
	// fmt.Println("root.Val: ", root.Val)


	if root.Val > key {
		// fmt.Println("root.Val > key: ", root.Val, key)
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		// fmt.Println("root.Val < key: ", root.Val, key)
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}else if root.Right == nil {
			return root.Left
		}else {
			minNode:=findClosestMin(root.Left)
			root.Val = minNode.Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	
	return root
	
}

func findClosestMin(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
	if root.Right == nil{
		return root
	}
	return findClosestMin(root.Right)
}